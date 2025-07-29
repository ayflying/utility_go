package oppo

import (
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/ayflying/utility_go/package/pay/common"
	"github.com/gogf/gf/v2/util/gconv"
	"hash"
	"math/rand"
	"net/url"
	"time"
)

func (p *OppoType) GenLoginBaseStr(bm map[string]interface{}, appKey, appSecret string) (string, string) {
	baseStr := fmt.Sprintf("oauthConsumerKey=%s&oauthToken=%s&oauthSignatureMethod=HMAC-SHA1&oauthTimestamp=%d&oauthNonce=%d&oauthVersion=1.0&",
		appKey, url.QueryEscape(gconv.String(bm["token"])), time.Now().Unix(), rand.Int31n(100000000))

	var h hash.Hash
	h = hmac.New(sha1.New, []byte(appSecret+"&"))
	h.Write([]byte(baseStr))

	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return baseStr, sign
}

func (p *OppoType) VerifySign(oppoPayPublicKey string, bm map[string]interface{}) (err error) {
	if oppoPayPublicKey == "" || bm == nil {
		return errors.New("oppoPayPublicKey or bm is nil")
	}

	bodySign := bm["sign"].(string)
	bodySignType := RSA
	signData := fmt.Sprintf("notifyId=%s&partnerOrder=%s&productName=%s&productDesc=%s&price=%s&count=%s&attach=%s",
		bm["notifyId"], bm["partnerOrder"], bm["productName"],
		bm["productDesc"], bm["price"], bm["count"], bm["attach"])
	pKey := common.FormatPublicKey(oppoPayPublicKey)
	if err = p.verifySign(signData, bodySign, bodySignType, pKey); err != nil {
		return err
	}
	return nil
}

func (p *OppoType) verifySign(signData, sign, signType, oppoPayPublicKey string) (err error) {
	var (
		h         hash.Hash
		hashs     crypto.Hash
		block     *pem.Block
		pubKey    interface{}
		publicKey *rsa.PublicKey
		ok        bool
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	if block, _ = pem.Decode([]byte(oppoPayPublicKey)); block == nil {
		return errors.New("OPPO公钥Decode错误")
	}
	if pubKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
	}
	if publicKey, ok = pubKey.(*rsa.PublicKey); !ok {
		return errors.New("OPPO公钥转换错误")
	}
	switch signType {
	case RSA:
		hashs = crypto.SHA1
	case RSA2:
		hashs = crypto.SHA256
	default:
		hashs = crypto.SHA256
	}
	h = hashs.New()
	h.Write([]byte(signData))
	return rsa.VerifyPKCS1v15(publicKey, hashs, h.Sum(nil), signBytes)
}
