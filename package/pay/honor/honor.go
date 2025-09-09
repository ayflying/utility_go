package honor

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"time"

	"github.com/ayflying/utility_go/package/pay/common"
	"github.com/ayflying/utility_go/pkg"
	"github.com/gogf/gf/v2/frame/g"
)

type Pay struct {
	PubKey       string `json:"pubKey"`
	AppId        string `json:"appId"`
	ClientSecret string `json:"client_secret"`
}

func New(pay *Pay) *Pay {
	return pay
}

func (p *Pay) GetToken(ctx context.Context) (accessToken string, err error) {
	type TokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	get, err := pkg.Cache("redis", "cache").GetOrSetFunc(ctx, "pay:honor:Sign:token", func(ctx context.Context) (value interface{}, err error) {

		url := TokenHost + "/oauth2/v3/token"
		get, err := g.Client().Post(ctx, url, g.Map{
			"client_id":     p.AppId,
			"client_secret": p.ClientSecret,
			"grant_type":    "client_credentials",
		})

		//var res *TokenResp
		//gjson.DecodeTo(get, &res)
		value = get.ReadAllString()
		return
	}, time.Hour)

	var res *TokenResp
	err = get.Scan(&res)
	accessToken = res.AccessToken

	return
}

// VerifyRSASignature 验证RSA数字签名
// data: 原始数据字节
// sign: 签名的Base64编码字符串
// pubKey: PEM格式的公钥字符串
// 返回验证结果和可能的错误
func (p *Pay) VerifyRSASignature(ctx context.Context, data []byte, signature string) (bool, error) {
	//req := g.RequestFromCtx(ctx).Request
	//post, err := common.ParseNotifyToBodyMap(req)
	//var data = gjson.MustEncode(post)

	// 解码Base64格式的签名
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, errors.New("签名解码失败: " + err.Error())
	}

	pubkey := common.FormatPublicKey(p.PubKey)
	// 解析PEM格式的公钥
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		return false, errors.New("无效的PEM格式公钥")
	}

	// 解析公钥
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, errors.New("公钥解析失败: " + err.Error())
	}

	// 类型断言为公钥
	rsaPubKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("不是有效的RSA公钥")
	}

	// 计算数据的SHA-256哈希
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)

	// 验证签名
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hash, signBytes)
	return err == nil, err
}
