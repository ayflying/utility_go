package playstore

import "context"

// GetProduct 获取应用内商品信息，该商品可以是管理型商品或订阅。
//
// - packageName: 应用的包名。
// - productID: 应用内商品的唯一标识符（SKU）。
//
// 返回值为InAppProduct类型的商品信息和可能出现的错误。
func (c *Client) GetProduct(ctx context.Context, packageName string, productID string) (*InAppProduct, error) {
	// 通过Google Play 商店API获取指定商品的信息
	var iap, err = c.service.Inappproducts.Get(packageName, productID).Context(ctx).Do()
	return &InAppProduct{iap}, err
}

// ConvertRegionPrices 将商品的价格区域配置转换为指定货币单位。
//
// - ctx: 上下文，用于控制请求的取消、超时等。
// - packageName: 应用的包名。
// - productID: 应用内商品的唯一标识符。
// - inAppProduct: 需要转换价格区域的InAppProduct对象。
//
// 返回转换后的InAppProduct对象和可能出现的错误。
//
// 注：此函数暂未实现。
//func (c *Client) ConvertRegionPrices(ctx context.Context, packageName string, productID string, inAppProduct InAppProduct) (*InAppProduct, error) {
//    // TODO: 实现商品价格区域转换逻辑
//    // c.service.
//
//    // 返回未实现的错误
//    return &InAppProduct{iap}, err
//}
