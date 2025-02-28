package playstore

// SubscriptionNotificationType 定义了订阅通知的类型。
type SubscriptionNotificationType int

// 预定义的订阅通知类型。
const (
	SubscriptionNotificationTypeRecovered            SubscriptionNotificationType = iota + 1 // 订阅已恢复
	SubscriptionNotificationTypeRenewed                                                      // 订阅已续订
	SubscriptionNotificationTypeCanceled                                                     // 订阅已取消
	SubscriptionNotificationTypePurchased                                                    // 订阅已购买
	SubscriptionNotificationTypeAccountHold                                                  // 订阅账户暂停
	SubscriptionNotificationTypeGracePeriod                                                  // 宽限期通知
	SubscriptionNotificationTypeRestarted                                                    // 订阅已重新开始
	SubscriptionNotificationTypePriceChangeConfirmed                                         // 订阅价格变更已确认
	SubscriptionNotificationTypeDeferred                                                     // 订阅延迟
	SubscriptionNotificationTypePaused                                                       // 订阅已暂停
	SubscriptionNotificationTypePauseScheduleChanged                                         // 暂停计划已更改
	SubscriptionNotificationTypeRevoked                                                      // 订阅已撤销
	SubscriptionNotificationTypeExpired                                                      // 订阅已过期
)

// OneTimeProductNotificationType 定义了一次性产品通知的类型。
type OneTimeProductNotificationType int

// 预定义的一次性产品通知类型。
const (
	OneTimeProductNotificationTypePurchased OneTimeProductNotificationType = iota + 1 // 一次性产品已购买
	OneTimeProductNotificationTypeCanceled                                            // 一次性产品已取消
)

// DeveloperNotification 是通过 Pub/Sub 主题发送给开发者的通知。
// 详细描述请参见：https://developer.android.com/google/play/billing/rtdn-reference#json_specification
type DeveloperNotification struct {
	Version                    string                     `json:"version"`                              // 版本号
	PackageName                string                     `json:"packageName"`                          // 应用包名
	EventTimeMillis            string                     `json:"eventTimeMillis"`                      // 事件发生时间（毫秒）
	SubscriptionNotification   SubscriptionNotification   `json:"subscriptionNotification,omitempty"`   // 订阅通知
	OneTimeProductNotification OneTimeProductNotification `json:"oneTimeProductNotification,omitempty"` // 一次性产品通知
	TestNotification           TestNotification           `json:"testNotification,omitempty"`           // 测试通知
}

// SubscriptionNotification 包含订阅状态通知类型、token 和订阅ID，用于通过Google Android Publisher API确认状态。
type SubscriptionNotification struct {
	Version          string                       `json:"version"`                    // 版本号
	NotificationType SubscriptionNotificationType `json:"notificationType,omitempty"` // 通知类型
	PurchaseToken    string                       `json:"purchaseToken,omitempty"`    // 购买token
	SubscriptionID   string                       `json:"subscriptionId,omitempty"`   // 订阅ID
}

// OneTimeProductNotification 包含一次性产品状态通知类型、token 和产品ID（SKU），用于通过Google Android Publisher API确认状态。
type OneTimeProductNotification struct {
	Version          string                         `json:"version"`                    // 版本号
	NotificationType OneTimeProductNotificationType `json:"notificationType,omitempty"` // 通知类型
	PurchaseToken    string                         `json:"purchaseToken,omitempty"`    // 购买token
	SKU              string                         `json:"sku,omitempty"`              // 产品ID（SKU）
}

// TestNotification 是仅通过Google Play开发者控制台发送的测试发布通知。
type TestNotification struct {
	Version string `json:"version"` // 版本号
}
