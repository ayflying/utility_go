package huawei

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Pay struct {
	ClientSecret string `json:"client_secret"`
	ClientId     string `json:"client_id"`
	//TokenUrl             string `json:"token_url"`
	ApplicationPublicKey string `json:"application_public_key"`
}

func New(cfg *Pay) *Pay {
	return cfg
}

// ConfirmPurchase 发货后确认购买接口（华为支付）消耗商品
// 功能：通知华为支付平台当前订单已完成发货，触发支付完成流程（需在商品实际发货后调用）
// 参数说明：
//   purchaseToken: 华为支付返回的购买令牌（唯一标识一笔具体的购买交易，由客户端支付成功后返回）
//   productId: 应用内商品的唯一标识（需与客户端发起支付时使用的productId一致）
//   accountFlag: 账户标识（用于区分不同账户体系/环境，如0-普通用户、1-企业用户，具体值由业务定义）
func (p *Pay) ConfirmPurchase(purchaseToken, productId string, accountFlag int) {
	// 构造请求体参数（包含购买令牌和产品ID）
	bodyMap := map[string]string{
		"purchaseToken": purchaseToken, // 华为支付返回的购买凭证
		"productId":     productId,     // 对应应用内商品的唯一标识
	}
	url := getOrderUrl(accountFlag) + "/applications/v2/purchases/confirm"
	bodyBytes, err := p.SendRequest(url, bodyMap)
	if err != nil {
		// 请求失败时记录错误日志（实际业务中建议增加重试或异常处理逻辑）
		log.Printf("err is %s", err)
	}
	// 打印响应结果（实际业务中需替换为具体处理逻辑，如更新订单状态、校验响应数据等）
	// TODO:  建议根据华为支付文档解析响应数据（如检查responseCode是否为0表示成功）
	log.Printf("%s", bodyBytes)
}

// VerifyToken 验证回调订单
//您可以调用本接口向华为应用内支付服务器校验支付结果中的购买令牌，确认支付结果的准确性。
func (p *Pay) VerifyToken(purchaseToken, productId string, accountFlag int) (res *PurchaseTokenData, err error) {
	bodyMap := map[string]string{"purchaseToken": purchaseToken, "productId": productId}
	url := getOrderUrl(accountFlag) + "/applications/purchases/tokens/verify"
	bodyBytes, err := p.SendRequest(url, bodyMap)
	if err != nil {
		g.Log().Error(gctx.New(), "err is %s", err)
	}
	var data *VerifyTokenRes
	err = gjson.DecodeTo(bodyBytes, &data)
	err = gjson.DecodeTo(data.PurchaseTokenData, &res)

	return
}

func (p *Pay) SendRequest(url string, bodyMap map[string]string) (string, error) {
	authHeaderString, err := p.BuildAuthorization()
	if err != nil {
		return "", err
	}
	bodyString, err := json.Marshal(bodyMap)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyString))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", authHeaderString)
	response, err := RequestHttpClient.Do(req)
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)

	//req, err := g.Client().Header(g.MapStrStr{
	//	"Content-Type":  "application/json; charset=UTF-8",
	//	"Authorization": authHeaderString,
	//}).Post(gctx.New(), url, bodyString)
	//defer req.Close()
	//var bodyBytes = req.ReadAll()

	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func (p *Pay) VerifyRsaSign(content string, sign string, publicKey string) error {
	//publicKey = common.FormatPublicKey(publicKey)
	publicKeyByte, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return err
	}
	pub, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(content))
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)
}

func (p *Pay) GetAppAt() (string, error) {
	//demoConfig := GetDefaultConfig()
	urlValue := url.Values{
		"grant_type":    {"client_credentials"},
		"client_secret": {p.ClientSecret},
		"client_id":     {p.ClientId},
	}
	resp, err := RequestHttpClient.PostForm(TokenUrl, urlValue)
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	//post := g.MapStrStr{
	//	"grant_type":    "client_credentials",
	//	"client_secret": p.ClientSecret,
	//	"client_id":     p.ClientId,
	//}
	//resp, err := g.Client().PostForm(gctx.New(), p.TokenUrl, post)
	//if err != nil {
	//	return "", err
	//}
	//resp.Close()
	//bodyBytes := resp.ReadAll()
	if err != nil {
		return "", err
	}
	var atResponse AtResponse
	json.Unmarshal(bodyBytes, &atResponse)
	if atResponse.AccessToken != "" {
		return atResponse.AccessToken, nil
	} else {
		return "", errors.New("Get token fail, " + string(bodyBytes))
	}
}

func (p *Pay) BuildAuthorization() (string, error) {
	appAt, err := p.GetAppAt()
	if err != nil {
		return "", err
	}
	oriString := fmt.Sprintf("APPAT:%s", appAt)
	var authString = base64.StdEncoding.EncodeToString([]byte(oriString))
	var authHeaderString = fmt.Sprintf("Basic %s", authString)
	return authHeaderString, nil
}
