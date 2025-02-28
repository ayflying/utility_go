package playstore

import (
	"context"
	"google.golang.org/api/androidpublisher/v3"
)

// VerifyProduct 验证产品状态
//
// 参数:
// - ctx: 上下文，用于控制请求的生命周期。
// - packageName: 应用的包名（例如，'com.some.thing'）。
// - productID: 内购产品的SKU（例如，'com.some.thing.inapp1'）。
// - token: 用户购买内购产品时设备上提供的令牌。
//
// 返回值:
// - *androidpublisher.ProductPurchase: 验证购买后的详细信息。
// - error: 执行过程中出现的错误。
func (c *Client) VerifyProduct(ctx context.Context, packageName string, productID string, token string) (*androidpublisher.ProductPurchase, error) {
	ps := androidpublisher.NewPurchasesProductsService(c.service)
	result, err := ps.Get(packageName, productID, token).Context(ctx).Do()
	return result, err
}

// AcknowledgeProduct 确认内购商品购买
//
// 注意！此函数必须在购买后的约24小时内对所有购买调用，否则购买将被自动撤销。
//
// 参数:
// - ctx: 上下文，用于控制请求的生命周期。
// - packageName: 应用的包名（例如，'com.some.thing'）。
// - productId: 内购产品的SKU（例如，'com.some.thing.inapp1'）。
// - token: 用户购买内购产品时设备上提供的令牌。
// - developerPayload: 开发者自定义信息。
//
// 返回值:
// - error: 执行过程中出现的错误。
func (c *Client) AcknowledgeProduct(ctx context.Context, packageName, productID, token, developerPayload string) error {
	ps := androidpublisher.NewPurchasesProductsService(c.service)
	acknowledgeRequest := &androidpublisher.ProductPurchasesAcknowledgeRequest{DeveloperPayload: developerPayload}
	err := ps.Acknowledge(packageName, productID, token, acknowledgeRequest).Context(ctx).Do()
	return err
}

// ConsumeProduct 消费购买应用内商品。
func (c *Client) ConsumeProduct(ctx context.Context, packageName, productID, token string) error {
	ps := androidpublisher.NewPurchasesProductsService(c.service)
	//acknowledgeRequest := &androidpublisher.PurchasesProductsConsumeCall{DeveloperPayload: developerPayload}
	//err := ps.Consume(packageName, productID, token).Context(ctx).Do()
	_, err := ps.Get(packageName, productID, token).Context(ctx).Do()
	return err
}

// Voidedpurchases 获取已撤销的购买列表
//
// 参数:
// - packageName: 应用的包名（例如，'com.some.thing'）。
//
// 返回值:
// - *androidpublisher.VoidedPurchasesListResponse: 已撤销购买的列表响应。
// - error: 执行过程中出现的错误。
func (c *Client) Voidedpurchases(packageName string) (*androidpublisher.VoidedPurchasesListResponse, error) {
	return androidpublisher.NewPurchasesVoidedpurchasesService(c.service).List(packageName).Do()
}
