package oppo

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"net/url"
	"strings"
)

const host = "https://iopen.game.heytapmobi.com"

// oppo参数类型
type LoginType struct {
	Token   string `json:"token"`
	Ssoid   string `json:"ssoid"`
	Channel int    `json:"channel"`
	AdId    string `json:"adId"`
}

//登录回复
type LoginResType struct {
	ResultCode string `json:"resultCode" dc:"响应码，成功为 200"`
	ResultMsg  string `json:"resultMsg" dc:"响应信息"`
	LoginToken string `json:"loginToken" dc:"透传的token"`
	Ssoid      string `json:"ssoid" dc:"透传的ssoid"`
	//AppKey       string `json:"appKey" dc:"秘钥key,因隐私安全规范，该字段目前已不返回信息"`
	UserName string `json:"userName" dc:"用户ssoid绑定的账户昵称"`
	//Email        string `json:"email" dc:"因隐私安全规范，该字段目前已不返回信息"`
	//MobileNumber string `json:"mobileNumber" dc:"因隐私安全规范，该字段目前已不返回信息"`
	//CreateTime   string `json:"createTime" dc:"因隐私安全规范，该字段目前已不返回信息"`
	UserStatus string `json:"userStatus" dc:"用户状态：NORMAL 表示正常"`
}

func (p *OppoType) FileIdInfo(ctx context.Context, oauthToken string, ssoid string) (res *LoginResType, err error) {
	url := host + "/sdkopen/user/fileIdInfo"
	header := p.GetHeader(oauthToken)
	getHtml, err := g.Client().Header(header).Get(ctx, url, g.Map{
		"token":  oauthToken,
		"fileId": ssoid,
	})
	getRes := getHtml.ReadAllString()
	gjson.DecodeTo(getRes, &res)
	g.Log().Debugf(ctx, "当前登陆请求的：%v", res)
	return

}

func (p *OppoType) GenParam(oauthToken, oauthTimestamp, oauthNonce string) string {
	// 注意：拼接的顺序不能有改变，不然会导致联运方验签失败
	params := []string{
		"oauthConsumerKey=" + url.QueryEscape(p.AppKey),
		"oauthToken=" + url.QueryEscape(oauthToken),
		"oauthSignatureMethod=" + url.QueryEscape("HMAC-SHA1"),
		"oauthTimestamp=" + url.QueryEscape(oauthTimestamp),
		"oauthNonce=" + url.QueryEscape(oauthNonce),
		"oauthVersion=" + url.QueryEscape("1.0"),
	}
	return strings.Join(params, "&") + "&"
}

// 生成签名
func (p *OppoType) GenOauthSignature(param string) string {
	oauthSignatureKey := p.AppSecret + "&"
	mac := hmac.New(sha1.New, []byte(oauthSignatureKey))
	io.WriteString(mac, param)
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return url.QueryEscape(signature)
}

func (p *OppoType) GetHeader(oauthToken string) (headers map[string]string) {

	// 没有做过 urlEncode 的 token，由游戏客户端调用 OPPO SDK 直接获取
	//oauthToken := "TICKET_Ajnxxxxx"
	oauthTimestamp := gtime.Now().TimestampStr()
	oauthNonce := grand.S(5)

	// 生成请求头参数和签名
	param := p.GenParam(oauthToken, oauthTimestamp, oauthNonce)
	oauthSignature := p.GenOauthSignature(param)

	// 封装请求头
	headers = map[string]string{
		"param":          param,
		"oauthSignature": oauthSignature,
	}

	//fmt.Println("游戏服务端登录鉴权请求头为：", headers)

	return
}
