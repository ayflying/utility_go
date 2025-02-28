package playstore

// GetStatus 获取产品的状态，例如产品是否处于活跃状态。
//
// 返回值 EProductStatus 代表产品状态。
// 可能的状态包括：
//
//	ProductStatus_Unspecified // 未指定状态。
//	ProductStatus_active // 产品已发布且在商店中处于活跃状态。
//	ProductStatus_inactive // 产品未发布，因此在商店中处于非活跃状态。
func (iap InAppProduct) GetStatus() EProductStatus {
	return EProductStatus(iap.AndroidPublisherInAppProduct.Status)
}

// GetSubscriptionPeriod 获取订阅的周期。
//
// 返回值 ESubscriptionPeriod 代表订阅周期。
// 可能的周期包括：
//
//	SubscriptionPeriod_Invalid : 无效的订阅（可能是消耗品）。
//	SubscriptionPeriod_OneWeek (一周)。
//	SubscriptionPeriod_OneMonth (一个月)。
//	SubscriptionPeriod_ThreeMonths (三个月)。
//	SubscriptionPeriod_SixMonths (六个月)。
//	SubscriptionPeriod_OneYear (一年)。
func (iap InAppProduct) GetSubscriptionPeriod() ESubscriptionPeriod {
	return ESubscriptionPeriod(iap.AndroidPublisherInAppProduct.SubscriptionPeriod)
}

// GetPurchaseType 获取产品的购买类型。
//
// 返回值 EPurchaseType 代表产品的购买类型。
// 可能的类型包括：
//
//	EPurchaseType_Unspecified (未指定购买类型)。
//	EPurchaseType_ManagedUser 可以被单次或多次购买（消耗品、非消耗品）。
//	EPurchaseType_Subscription （应用内产品，具有周期性消费）。
func (iap InAppProduct) GetPurchaseType() EPurchaseType {
	return EPurchaseType(iap.AndroidPublisherInAppProduct.PurchaseType)
}
