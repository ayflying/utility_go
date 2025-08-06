package huawei

type CallbackType struct {
	Version           string             `json:"version"`
	NotifyTime        int64              `json:"notifyTime"`
	EventType         string             `json:"eventType"`
	ApplicationId     string             `json:"applicationId"`
	OrderNotification *OrderNotification `json:"orderNotification"`
	SubNotification   *SubNotification   `json:"subNotification"`
}

type OrderNotification struct {
	Version          string `json:"version" dc:"通知版本：v2"`
	NotificationType int    `json:"notificationType" dc:"通知事件的类型，取值如下：1：支付成功 2：退款成功"`
	PurchaseToken    string `json:"purchaseToken" dc:"待下发商品的购买Token"`
	ProductId        string `json:"productId" dc:"商品ID"`
}

type SubNotification struct {
	StatusUpdateNotification *StatusUpdateNotification `json:"statusUpdateNotification" dc:"通知消息"`
	NotificationSignature    string                    `json:"notificationSignature" dc:"对statusUpdateNotification字段的签名字符串，签名算法为signatureAlgorithm表示的签名算法。"`
	Version                  string                    `json:"version" dc:"通知版本：v2"`
	SignatureAlgorithm       string                    `json:"signatureAlgorithm" dc:"签名算法。"`
}

// StatusUpdateNotification 订阅状态更新通知
type StatusUpdateNotification struct {
	Environment                       string `json:"environment" dc:"发送通知的环境。PROD：正式环境；Sandbox：沙盒测试"`
	NotificationType                  int    `json:"notificationType" dc:"通知事件的类型，具体定义需参考相关文档说明"`
	SubscriptionID                    string `json:"subscriptionId" dc:"订阅ID"`
	CancellationDate                  int64  `json:"cancellationDate" dc:"撤销订阅时间或退款时间，UTC时间戳，以毫秒为单位，仅在notificationType取值为CANCEL的场景下会传入"`
	OrderID                           string `json:"orderId" dc:"订单ID，唯一标识一笔需要收费的收据，由华为应用内支付服务器在创建订单以及订阅型商品续费时生成。每一笔新的收据都会使用不同的orderId。通知类型为NEW_RENEWAL_PREF时不存在"`
	LatestReceipt                     string `json:"latestReceipt" dc:"最近的一笔收据的token，仅在notificationType取值为INITIAL_BUY 、RENEWAL或INTERACTIVE_RENEWAL并且续期成功情况下传入"`
	LatestReceiptInfo                 string `json:"latestReceiptInfo" dc:"最近的一笔收据，JSON字符串格式，包含的参数请参见InappPurchaseDetails，在notificationType取值为CANCEL时无值"`
	LatestReceiptInfoSignature        string `json:"latestReceiptInfoSignature" dc:"对latestReceiptInfo的签名字符串，签名算法为statusUpdateNotification中的signatureAlgorithm。服务器在收到签名字符串后，需要参见对返回结果验签使用IAP公钥对latestReceiptInfo的JSON字符串进行验签。公钥获取请参见查询支付服务信息"`
	LatestExpiredReceipt              string `json:"latestExpiredReceipt" dc:"最近的一笔过期收据的token"`
	LatestExpiredReceiptInfo          string `json:"latestExpiredReceiptInfo" dc:"最近的一笔过期收据，JSON字符串格式，在notificationType取值为RENEWAL或INTERACTIVE_RENEWAL时有值"`
	LatestExpiredReceiptInfoSignature string `json:"latestExpiredReceiptInfoSignature" dc:"对latestExpiredReceiptInfo的签名字符串，签名算法为statusUpdateNotification中的signatureAlgorithm。服务器在收到签名字符串后，需要参见对返回结果验签使用IAP公钥对latestExpiredReceiptInfo的JSON字符串进行验签。公钥获取请参见查询支付服务信息"`
	AutoRenewStatus                   int    `json:"autoRenewStatus" dc:"续期状态。取值说明：1：当前周期到期后正常续期；0：用户已终止续期"`
	RefundPayOrderId                  string `json:"refundPayOrderId" dc:"退款交易号，在notificationType取值为CANCEL时有值"`
	ProductID                         string `json:"productId" dc:"订阅型商品ID"`
	ApplicationID                     string `json:"applicationId" dc:"应用ID"`
	ExpirationIntent                  int    `json:"expirationIntent" dc:"超期原因，仅在notificationType为RENEWAL或INTERACTIVE_RENEWAL时并且续期失败情况下有值"`
	PurchaseToken                     string `json:"purchaseToken" dc:"订阅token，与上述订阅ID字段subscriptionId对应。"`
}

type AtResponse struct {
	AccessToken string `json:"access_token"`
}

type VerifyTokenRes struct {
	ResponseCode       string `json:"responseCode"`
	PurchaseTokenData  string `json:"purchaseTokenData"`
	DataSignature      string `json:"dataSignature"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
}

type PurchaseTokenData struct {
	AutoRenewing        bool   `json:"autoRenewing" dc:"表示订阅是否自动续费"`
	OrderId             string `json:"orderId" dc:"订单ID，唯一标识一笔订单"`
	PackageName         string `json:"packageName" dc:"应用的包名"`
	ApplicationId       int    `json:"applicationId" dc:"应用ID，以整数形式表示"`
	ApplicationIdString string `json:"applicationIdString" dc:"应用ID的字符串形式"`
	Kind                int    `json:"kind" dc:"购买类型的某种标识，具体含义可能取决于业务逻辑"`
	ProductId           string `json:"productId" dc:"商品ID，用于标识购买的商品"`
	ProductName         string `json:"productName" dc:"商品名称"`
	PurchaseTime        int64  `json:"purchaseTime" dc:"购买时间，可能是某种特定格式的时间表示"`
	PurchaseTimeMillis  int64  `json:"purchaseTimeMillis" dc:"购买时间，以毫秒为单位的时间戳"`
	PurchaseState       int    `json:"purchaseState" dc:"购买状态，不同的整数值代表不同的状态，具体含义取决于业务逻辑"`
	DeveloperPayload    string `json:"developerPayload" dc:"开发者自定义负载数据"`
	PurchaseToken       string `json:"purchaseToken" dc:"购买令牌"`
	ResponseCode        string `json:"responseCode" dc:"响应代码，用于表示购买操作的响应结果"`
	ConsumptionState    int    `json:"consumptionState" dc:"消费状态，不同的整数值代表不同的消费状态，具体含义取决于业务逻辑"`
	Confirmed           int    `json:"confirmed" dc:"确认状态，不同的整数值代表不同的确认情况，具体含义取决于业务逻辑"`
	PurchaseType        int    `json:"purchaseType" dc:"购买类型，不同的整数值代表不同的购买类型，具体含义取决于业务逻辑"`
	Currency            string `json:"currency" dc:"货币类型"`
	Price               int    `json:"price" dc:"商品价格"`
	Country             string `json:"country" dc:"购买所在国家"`
	PayOrderId          string `json:"payOrderId" dc:"支付订单ID"`
	PayType             string `json:"payType" dc:"支付类型"`
	SdkChannel          string `json:"sdkChannel" dc:"SDK渠道"`
}
