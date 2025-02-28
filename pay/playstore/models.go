package playstore

import (
	"context"
	"google.golang.org/api/androidpublisher/v3"
)

// IABProduct 接口定义了商品服务的基本操作。
type IABProduct interface {
	// VerifyProduct 验证指定的内购产品购买信息。
	// ctx: 上下文，用于控制请求的取消、超时等。
	// packageName: 应用包名。
	// productId: 内购商品ID。
	// purchaseToken: 购买凭证。
	// 返回经过验证的购买信息和可能的错误。
	VerifyProduct(context.Context, string, string, string) (*androidpublisher.ProductPurchase, error)

	// AcknowledgeProduct 确认指定的内购产品的购买。
	// ctx: 上下文。
	// packageName: 应用包名。
	// productId: 内购商品ID。
	// purchaseToken: 购买凭证。
	// orderId: 订单ID。
	// 返回可能发生的错误。
	AcknowledgeProduct(context.Context, string, string, string, string) error
}

// IABSubscription 接口定义了订阅服务的基本操作。
type IABSubscription interface {
	// AcknowledgeSubscription 确认指定订阅的购买。
	// ctx: 上下文。
	// packageName: 应用包名。
	// subscriptionId: 订阅ID。
	// purchaseToken: 购买凭证。
	// acknowledgeRequest: 确认请求参数。
	// 返回可能发生的错误。
	AcknowledgeSubscription(context.Context, string, string, string, *androidpublisher.SubscriptionPurchasesAcknowledgeRequest) error

	// VerifySubscription 验证指定订阅的购买信息。
	// ctx: 上下文。
	// packageName: 应用包名。
	// subscriptionId: 订阅ID。
	// purchaseToken: 购买凭证。
	// 返回经过验证的订阅购买信息和可能的错误。
	VerifySubscription(context.Context, string, string, string) (*androidpublisher.SubscriptionPurchase, error)

	// CancelSubscription 取消指定的订阅。
	// ctx: 上下文。
	// packageName: 应用包名。
	// subscriptionId: 订阅ID。
	// purchaseToken: 购买凭证。
	// 返回可能发生的错误。
	CancelSubscription(context.Context, string, string, string) error

	// RefundSubscription 为指定的订阅办理退款。
	// ctx: 上下文。
	// packageName: 应用包名。
	// subscriptionId: 订阅ID。
	// purchaseToken: 购买凭证。
	// 返回可能发生的错误。
	RefundSubscription(context.Context, string, string, string) error

	// RevokeSubscription 撤销指定的订阅。
	// ctx: 上下文。
	// packageName: 应用包名。
	// subscriptionId: 订阅ID。
	// purchaseToken: 购买凭证。
	// 返回可能发生的错误。
	RevokeSubscription(context.Context, string, string, string) error
}

// Client 结构体实现了 IABSubscription 接口，提供了具体的操作实现。
type Client struct {
	service *androidpublisher.Service
}

// InAppProduct 结构体封装了 androidpublisher.InAppProduct，并提供了一些辅助方法。
type InAppProduct struct {
	AndroidPublisherInAppProduct *androidpublisher.InAppProduct
}
