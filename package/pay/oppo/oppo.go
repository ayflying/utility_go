package oppo

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
	"strings"
)

// 跟充值平台通信的加密key
//const PUBLIC_KEY = `dfsdfs`

type OppoType struct {
	PublicKey string `json:"public_key"`
}

func New(PublicKey string) *OppoType {

	return &OppoType{
		PublicKey: PublicKey,
	}
}

func (p *OppoType) Verify(ctx context.Context, data map[string]string) error {
	// 解析请求参数
	for k, v := range data {
		if v == "" || v == "0" {
			delete(data, k)
		}
	}

	//data["notifyId"] = getParam(r, "notifyId")
	//data["partnerOrder"] = getParam(r, "partnerOrder")
	//data["productName"] = getParam(r, "productName")
	//data["productDesc"] = getParam(r, "productDesc")
	//data["price"] = getParam(r, "price")
	//data["count"] = getParam(r, "count")
	//data["attach"] = getParam(r, "attach")
	//data["sign"] = getParam(r, "sign")

	// 验证签名
	result, err := p.rsaVerify(data)
	if err != nil {
		//http.Error(w, "Verification error: "+err.Error(), http.StatusInternalServerError)
		g.Log().Errorf(ctx, "Verification error: %v", err.Error())
		return err
	}

	if result {
		// TODO::验证成功，处理后续逻辑
		//fmt.Fprint(w, "Verification successful")
		//g.Log().Errorf(ctx, "Verification error: %v", err.Error())
	} else {
		// TODO::验证失败，处理后续逻辑
		//http.Error(w, "Verification failed", http.StatusBadRequest)
		g.Log().Error(ctx, "Verification failed")
		err = gerror.New("Verification failed")
	}
	return nil
}

func (p *OppoType) getParam(r *http.Request, paramName string) string {
	r.ParseForm()
	if value := r.FormValue(paramName); value != "" {
		return strings.TrimSpace(value)
	}
	return ""
}

func (p *OppoType) rsaVerify(contents map[string]string) (bool, error) {
	// 构建待签名字符串
	strContents := fmt.Sprintf("notifyId=%s&partnerOrder=%s&productName=%s&productDesc=%s&price=%s&count=%s&attach=%s",
		contents["notifyId"], contents["partnerOrder"], contents["productName"],
		contents["productDesc"], contents["price"], contents["count"], contents["attach"])

	// 解析公钥
	publicKey := p.PublicKey
	pemData := []byte("-----BEGIN PUBLIC KEY-----\n" +
		strings.ReplaceAll(publicKey, " ", "\n") +
		"\n-----END PUBLIC KEY-----")

	block, _ := pem.Decode(pemData)
	if block == nil {
		return false, fmt.Errorf("failed to decode PEM block")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("public key is not an RSA public key")
	}

	// 解码签名
	signature, err := base64.StdEncoding.DecodeString(contents["sign"])
	if err != nil {
		return false, err
	}

	// 计算内容的哈希值
	hash := sha1.New()
	hash.Write([]byte(strContents))
	hashed := hash.Sum(nil)

	// 验证签名
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, hashed, signature)
	return err == nil, err
}
