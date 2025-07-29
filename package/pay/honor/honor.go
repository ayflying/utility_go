package honor

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/ayflying/utility_go/package/pay/common"
)

type Pay struct {
	PubKey string `json:"pubKey"`
	AppId  string `json:"appId"`
}

func New(pay *Pay) *Pay {
	return &Pay{
		AppId:  pay.AppId,
		PubKey: pay.PubKey,
	}
}

// VerifyRSASignature 验证RSA数字签名
// data: 原始数据字节
// sign: 签名的Base64编码字符串
// pubKey: PEM格式的公钥字符串
// 返回验证结果和可能的错误
func (p *Pay) VerifyRSASignature(data []byte, sign string) (bool, error) {
	// 解码Base64格式的签名
	signBytes, err := base64.StdEncoding.DecodeString(sign)
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
