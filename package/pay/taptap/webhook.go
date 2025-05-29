package taptap

import "github.com/gogf/gf/v2/encoding/gjson"

type WebhookData struct {
	Order     *Order `json:"order"`
	EventType string `json:"event_type"`
}

// Order 订单信息结构体
type Order struct {
	OrderID       string `json:"order_id" dc:"订单唯一ID"`
	PurchaseToken string `json:"purchase_token" dc:"用于订单核销的token"`
	ClientID      string `json:"client_id" dc:"应用的Client ID"`
	OpenID        string `json:"open_id" dc:"用户的开放平台ID"`
	UserRegion    string `json:"user_region" dc:"用户地区"`
	GoodsOpenID   string `json:"goods_open_id" dc:"商品唯一ID"`
	GoodsName     string `json:"goods_name" dc:"商品名称"`
	Status        string `json:"status" dc:"订单状态"`
	Amount        string `json:"amount" dc:"金额（本币金额x1,000,000）"`
	Currency      string `json:"currency" dc:"币种"`
	CreateTime    string `json:"create_time" dc:"创建时间"`
	PayTime       string `json:"pay_time" dc:"支付时间"`
	Extra         string `json:"extra" dc:"商户自定义数据，如角色信息等，长度不超过255 UTF-8字符"`
}

func (p *pTapTap) Webhook(body []byte) (res string, err error) {
	var data *WebhookData
	gjson.DecodeTo(body, &data)

	switch data.EventType {
	case "charge.succeeded": //充值成功
		//todo 处理订单信息

	case "refund.succeeded": //退款成功
	case "refund.failed": //退款失败
	}

	return
}
