package honor

type PayCallbackHeader struct {
	Charset   string `json:"charset" dc:"字符集，当前只支持utf-8。"`
	SignType  string `json:"signType" dc:"签名算法类型, 当前只支持RSA"`
	Sign      string `json:"sign" dc:"notificationMessage的签名，已废弃，请用signature。"`
	Signature string `json:"signature" dc:"对data的签名。"`
	AppId     string `json:"appId" dc:"应用ID"`
}

type PayCallback struct {
	Env       string          `json:"env" dc:"发送通知的环境，sandbox为沙盒测试环境，非sandbox为正式环境"`
	EventType string          `json:"eventType" dc:"事件类型，如付款成功、退款失败等"`
	EventCode int             `json:"eventCode" dc:"事件类型对应的code值"`
	Version   string          `json:"version" dc:"iap版本"`
	EventTime string          `json:"eventTime" dc:"通知时间"`
	Data      PayCallbackData `json:"data" dc:"通知内容notificationMessage的json字符串"`
}

type PayCallbackData struct {
	AppId            string `json:"appId"        dc:"应用ID"`
	OrderId          string `json:"orderId"      dc:"订单ID"`
	BizOrderNo       string `json:"bizOrderNo,omitempty" dc:"max-length:64#业务订单号"`
	ProductType      int    `json:"productType"  dc:"商品类型0：消耗型，1：非消耗型，2：订阅型"`
	ProductId        string `json:"productId"    dc:"商品ID"`
	ProductName      string `json:"productName"  dc:"商品名称"`
	PurchaseTime     int64  `json:"purchaseTime" dc:"购买时间UTC时间戳(毫秒)"`
	PurchaseState    int    `json:"purchaseState" dc:"订单状态  0:已购买 1:已退款 2:付款失败 3:退款失败 4:未支付 5:退款中"`
	ConsumptionState int    `json:"consumptionState" dc:"消耗状态  0:未消耗 1:已消耗"`
	PurchaseToken    string `json:"purchaseToken" dc:"购买令牌"`
	Currency         string `json:"currency"     dc:"币种"`
	Price            string `json:"price"        dc:"商品价格"`
	PayMoney         string `json:"payMoney"     dc:"实际支付金额"`
	DeveloperPayload string `json:"developerPayload,omitempty" dc:"max-length:1024#商户信息"`
	OriOrder         string `json:"oriOrder"     dc:"原订单信息"`
	SandboxFlag      int    `json:"sandboxFlag"  dc:"沙盒标识"`
	AgreementNo      string `json:"agreementNo,omitempty" dc:"订阅合约号"`
	ExecuteTime      string `json:"executeTime,omitempty" dc:"下次扣费时间(订阅)"`
	SecondChargeTime int64  `json:"secondChargeTime,omitempty" dc:"第二次扣费时间(订阅升级)"`
	OldProductId     string `json:"oldProductId,omitempty" dc:"老商品ID(订阅升级)"`
	SubStartTime     string `json:"subStartTime,omitempty" dc:"订阅开始时间"`
	SubEndTime       string `json:"subEndTime,omitempty" dc:"订阅结束时间"`
	OriginalPrice    string `json:"originalPrice"  dc:"原始价格"`
	CancelTime       string `json:"cancelTime,omitempty" dc:"订阅取消时间"`
}
