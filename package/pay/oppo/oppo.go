package oppo

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// 跟充值平台通信的加密key
//const PUBLIC_KEY = `dfsdfs`

type OppoType struct {
	AppId     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	PublicKey string `json:"public_key"`
}

func New(cfg *OppoType) *OppoType {

	return &OppoType{
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		PublicKey: cfg.PublicKey,
	}
}

func (p *OppoType) Verify(ctx context.Context) (err error) {
	// OPPO公钥. 在官方给的 demo 中. 无需修改,改了就验证不过
	oppoPublicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmreYIkPwVovKR8rLHWlFVw7YDfm9uQOJKL89Smt6ypXGVdrAKKl0wNYc3/jecAoPi2ylChfa2iRu5gunJyNmpWZzlCNRIau55fxGW0XEu553IiprOZcaw5OuYGlf60ga8QT6qToP0/dpiL/ZbmNUO9kUhosIjEu22uFgR+5cYyQIDAQAB"
	//oppoPublicKey := p.PublicKey
	// 解析请求参数
	bodyMap, err := p.ParseNotifyToBodyMap(g.RequestFromCtx(ctx).Request)
	if err != nil {
		// 解析失败, 处理错误逻辑
		return
	}

	err = p.VerifySign(oppoPublicKey, bodyMap)
	return
}
