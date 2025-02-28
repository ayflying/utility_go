package pay

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/v2/frame/g"
)

type AliPay struct {
	Client *alipay.Client
}

func Alipay() *AliPay {
	var pay = &AliPay{}
	var err error

	cfg, err := g.Cfg().Get(ctx, "pay.alipay")
	cfgMap := cfg.MapStrStr()
	appId := cfgMap["appid"]
	privateKey := cfgMap["privateKey"]
	isProd, _ := g.Cfg().Get(ctx, "pay.alipay.isProd")
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	pay.Client, err = alipay.NewClient(appId, privateKey, isProd.Bool())
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//pay.Client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	pay.Client.DebugSwitch = gopay.DebugOn

	pay.Client.SetCharset(alipay.UTF8). // 设置字符编码，不设置默认 utf-8
						SetSignType(alipay.RSA2) // 设置签名类型，不设置默认 RSA2

	return pay
}
