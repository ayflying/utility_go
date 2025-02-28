package playstore

// EProductStatus 定义了产品的状态，例如产品是否处于活跃状态。
type EProductStatus string

// 定义了产品可能的状态常量。
const (
	ProductStatus_Unspecified EProductStatus = "statusUnspecified" // 未指定状态。
	ProductStatus_active      EProductStatus = "active"            // 产品已发布且在商店中处于活跃状态。
	ProductStatus_inactive    EProductStatus = "inactive"          // 产品未发布，因此在商店中处于非活跃状态。
)

// ESubscriptionPeriod 定义了订阅的周期。
type ESubscriptionPeriod string

// 定义了订阅可能的周期常量。
const (
	SubscriptionPeriod_Invalid     ESubscriptionPeriod = ""    // 无效的订阅（可能是消耗品）。
	SubscriptionPeriod_OneWeek     ESubscriptionPeriod = "P1W" // 一周。
	SubscriptionPeriod_OneMonth    ESubscriptionPeriod = "P1M" // 一个月。
	SubscriptionPeriod_ThreeMonths ESubscriptionPeriod = "P3M" // 三个月。
	SubscriptionPeriod_SixMonths   ESubscriptionPeriod = "P6M" // 六个月。
	SubscriptionPeriod_OneYear     ESubscriptionPeriod = "P1Y" // 一年。
)

// EPurchaseType 定义了产品的购买类型，例如周期性订阅。
type EPurchaseType string

// 定义了产品可能的购买类型常量。
const (
	EPurchaseType_Unspecified  EPurchaseType = "purchaseTypeUnspecified" // 未指定购买类型。
	EPurchaseType_ManagedUser  EPurchaseType = "managedUser"             // 默认的产品类型 - 可以单次或多次购买（消耗品、非消耗品）。
	EPurchaseType_Subscription EPurchaseType = "subscription"            // 应用内具有周期性的产品。
)
