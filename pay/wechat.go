package pay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/go-pay/crypto/xpem"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/util/convert"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

var (
// ctx = gctx.New()

)

// GooglePay 是一个处理Google支付的结构体。
type WechatPay struct {
	Client     *wechat.ClientV3
	PrivateKey string
}

func Wechat() *WechatPay {
	var pay = &WechatPay{}
	var err error

	cfg, _ := g.Cfg().Get(ctx, "pay.wechat")
	cfgMap := cfg.MapStrStr()
	MchId := cfgMap["mchid"]
	SerialNo := cfgMap["serialNo"]
	APIv3Key := cfgMap["apiV3Key"]
	PrivateKey := gfile.GetContents("manifest/pay/apiclient_key.pem")
	//PrivateKey := cfgMap["privateKey"]

	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	pay.Client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, PrivateKey)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}

	err = pay.Client.AutoVerifySign()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}

	return pay
}

// v3 鉴权请求Header
func (c *WechatPay) Authorization(appid string, timestamp int64, nonceStr string, prepay_id string) (string, error) {
	//var (
	//	jb        = ""
	//	timestamp = time.Now().Unix()
	//	nonceStr  = util.RandomString(32)
	//)
	//if bm != nil {
	//	jb = bm.JsonBody()
	//}
	//path = strings.TrimSuffix(path, "?")
	ts := convert.Int642String(timestamp)
	_str := strings.Join([]string{appid, ts, nonceStr, prepay_id}, "\n") + "\n"
	//_str := appid + "\n" + timestamp + "\n"  + nonceStr + "\n" + jb + "\n"

	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (c *WechatPay) rsaSign(str string) (string, error) {
	//if c.privateKey == nil {
	//	return "", errors.New("privateKey can't be nil")
	//}

	privateKey := gfile.GetContents("manifest/pay/apiclient_key.pem")
	priKey, err := xpem.DecodePrivateKey([]byte(privateKey))

	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return gopay.NULL, fmt.Errorf("[%w]: %+v", gopay.SignatureErr, err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}
