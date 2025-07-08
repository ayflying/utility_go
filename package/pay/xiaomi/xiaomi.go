package xiaomi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// Config 小米支付配置信息
type Config struct {
	AppID     string `json:"app_id"`     // 应用ID
	AppSecret string `json:"app_secret"` // 应用密钥
	//PrivateKey string // 商户私钥（如需证书）
	//MIAPIURL   string // 小米支付API基础地址
	//IsSandbox  bool   // 是否沙箱环境
}

// Miipay 小米支付客户端
type MiPay struct {
	config *Config
}

func New() *MiPay {
	_cfg, _ := g.Cfg().Get(gctx.New(), "pay.xiaomi")
	var cfg *Config
	_cfg.Scan(&cfg)
	return &MiPay{
		config: cfg,
	}
}
