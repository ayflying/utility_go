package xiaomi

import "github.com/gogf/gf/v2/os/gtime"

type PayCallback struct {
	AppID              string      `json:"appId" dc:"游戏ID" required:"true"`
	CPOrderID          string      `json:"cpOrderId" dc:"开发商订单ID" required:"true"`
	CPUserInfo         string      `json:"cpUserInfo" dc:"开发商透传信息" required:"false"`
	OrderConsumeType   int         `json:"orderConsumeType" dc:"订单类型：10：普通订单 11：直充直消订单" required:"false"`
	OrderID            string      `json:"orderId" dc:"游戏平台订单ID" required:"true"`
	OrderStatus        string      `json:"orderStatus" dc:"订单状态，TRADE_SUCCESS代表成功" required:"true"`
	PayFee             int         `json:"payFee" dc:"支付金额，单位为分，即0.01米币。（请务必使用payFee字段值与游戏发起订单金额做校验，确保订单金额一致性）" required:"true"`
	PayTime            *gtime.Time `json:"payTime" dc:"支付时间，格式yyyy-MM-dd HH:mm:ss" required:"true"`
	ProductCode        string      `json:"productCode" dc:"商品代码" required:"true"`
	ProductCount       int         `json:"productCount" dc:"商品数量" required:"true"`
	ProductName        string      `json:"productName" dc:"商品名称" required:"true"`
	UID                string      `json:"uid" dc:"用户ID" required:"true"`
	PartnerGiftConsume int64       `json:"partnerGiftConsume" dc:"使用游戏券金额（如果订单使用游戏券则有，long型），如果有则参与签名" required:"false"`
	Signature          string      `json:"signature" dc:"签名，签名方法见后面说明" required:"true"`
}
