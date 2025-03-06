package pay

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/api/androidpublisher/v3"
	"new-gitlab.adesk.com/public_project/utility_go/package/pay/playstore"
)

var (
	ctx = gctx.New()
)

// GooglePay 是一个处理Google支付的结构体。
type GooglePay struct {
	c *playstore.Client
}

// Init 初始化GooglePay客户端。
// data: 初始化客户端所需的配置数据。
func (p *GooglePay) Init(data []byte) {
	var err error
	p.c, err = playstore.New(data)
	if err != nil {
		panic(err) // 如果初始化失败，则panic。
	}
}

// VerifyPay 验证用户的支付。
// userId: 用户ID。
// OrderId: 订单ID。
// package1: 应用包名。
// subscriptionID: 订阅ID。
// purchaseToken: 购买凭证。
// cb: 验证结果的回调函数，如果验证成功，会调用此函数。
// 返回值: 执行错误。
func (p *GooglePay) VerifyPay(userId int64, OrderId, package1, subscriptionID, purchaseToken string, cb func(string, string) error) error {
	info, err := p.c.VerifyProduct(context.Background(), package1, subscriptionID, purchaseToken)
	if err != nil {
		return gerror.Cause(err) // 验证产品失败，返回错误。
	}
	if info.PurchaseState == 0 {
		if err := cb(subscriptionID, info.OrderId); err != nil {
			return gerror.Cause(err) // 调用回调函数失败，返回错误。
		}
	} else {
		return nil // 验证结果不为购买状态，直接返回nil。
	}
	return nil
}

// VerifyPayV1 是VerifyPay的另一个版本，用于验证订阅支付。
// package1: 应用包名。
// subscriptionID: 订阅ID。
// purchaseToken: 购买凭证。
// cb: 验证结果的回调函数。
// 返回值: 执行错误。
func (p *GooglePay) VerifyPayV1(package1, subscriptionID, purchaseToken string, cb func(string, string) error) error {
	//g.Log().Infof(ctx, "VerifyPayV1: package = %v subscriptionID = %v, purchaseToken = %v", package1, subscriptionID, purchaseToken)
	info, err := p.c.VerifyProduct(context.Background(), package1, subscriptionID, purchaseToken)
	if err != nil {
		return gerror.Cause(err) // 验证产品失败，返回错误。
	}
	if info.PurchaseState == 0 {
		if err := cb(subscriptionID, info.OrderId); err != nil {
			return gerror.Cause(err) // 调用回调函数失败，返回错误。
		}
	} else {
		return nil // 验证结果不为购买状态，直接返回nil。
	}
	return nil
}

// VerifyPayV2 是VerifyPay的另一个版本，支持不同类型产品的验证。
// types: 验证的产品类型。
// package1: 应用包名。
// subscriptionID: 订阅ID。
// purchaseToken: 购买凭证。
// cb: 验证结果的回调函数。
// 返回值: 执行错误。
func (p *GooglePay) VerifyPayV2(types int32, package1, subscriptionID, purchaseToken string, cb func(string, string) error) error {
	//g.Log().Infof(ctx, "VerifyPayV1: package = %v subscriptionID = %v, purchaseToken = %v", package1, subscriptionID, purchaseToken)
	switch types {
	case 0:
		info, err := p.c.VerifyProduct(context.Background(), package1, subscriptionID, purchaseToken)
		if err != nil {
			return gerror.Cause(err) // 验证产品失败，返回错误。
		}
		if info.PurchaseState == 0 {
			if err := cb(subscriptionID, info.OrderId); err != nil {
				return gerror.Cause(err) // 调用回调函数失败，返回错误。
			}
		}
	case 1:
		info, err := p.c.VerifySubscription(context.Background(), package1, subscriptionID, purchaseToken)
		if err != nil {
			return gerror.Cause(err) // 验证订阅失败，返回错误。
		}
		if len(info.OrderId) != 0 {
			if err := cb(subscriptionID, info.OrderId); err != nil {
				return gerror.Cause(err) // 调用回调函数失败，返回错误。
			}
		}
	}

	return nil
}

//func (p *GooglePay) VerifyPayTest(package1, subscriptionID, purchaseToken string) (*androidpublisher.ProductPurchase, error) {
//	return p.c.VerifyProduct(context.Background(), package1, subscriptionID, purchaseToken)
//}

func (p *GooglePay) VerifySubscriptionTest(package1, subscriptionID, purchaseToken string) (interface{}, error) {
	return p.c.VerifySubscription(context.Background(), package1, subscriptionID, purchaseToken)
}

// VerifySubSciption google 检查订阅是否有效
func (p *GooglePay) VerifySubSciption(package1, subscriptionID, purchaseToken string) (string, error) {
	info, err := p.c.VerifySubscription(context.Background(), package1, subscriptionID, purchaseToken)
	if err != nil {
		return "", gerror.Cause(err)
	}
	if len(info.OrderId) != 0 {
		return info.OrderId, nil
	}
	return "", nil
}

// 获取已撤销的购买列表
func (p *GooglePay) GetRevokedPurchaseList(package1 string) (res *androidpublisher.VoidedPurchasesListResponse, err error) {
	res, err = p.c.Voidedpurchases(package1)
	//return p.c.GetRevokedPurchaseList(context.Background(), package1)
	return
}

// Acknowledge 确认购买应用内商品。
// Method: purchases.products.acknowledge y
func (p *GooglePay) Acknowledge(ctx context.Context, packageName, productID, token, developerPayload string) (err error) {
	err = p.c.AcknowledgeProduct(ctx, packageName, productID, token, developerPayload)
	return
}

// Consume 消费购买应用内商品。
func (p *GooglePay) Consume(ctx context.Context, packageName, productID, token string) (err error) {
	err = p.c.ConsumeProduct(ctx, packageName, productID, token)
	return
}

// 谷歌支付支付凭证校验V1
func (s *GooglePay) GooglePayTokenV1(token string) (err error) {

	type PayOrderType struct {
		Payload       string `json:"Payload"`
		Store         string `json:"Store"`
		TransactionID string `json:"TransactionID"`
	}
	type PayloadType struct {
		Json       string   `json:"json"`
		Signature  string   `json:"signature"`
		SkuDetails []string `json:"skuDetails"`
	}
	type PayJson struct {
		PackageName   string `json:"packageName"`
		ProductId     string `json:"productId"`
		PurchaseTime  int64  `json:"purchaseTime"`
		PurchaseState int    `json:"purchaseState"`
		PurchaseToken string `json:"purchaseToken"`
		Quantity      int    `json:"quantity"`
		Acknowledged  bool   `json:"acknowledged"`
		OrderId       string `json:"orderId"`
	}

	var data PayOrderType
	gconv.Struct(token, &data)
	var payload PayloadType
	gconv.Struct(data.Payload, &payload)
	var payJson PayJson
	gconv.Struct(payload.Json, &payJson)
	if gstr.Pos(payJson.OrderId, "GPA.") < 0 {
		err = gerror.New("GPA验证失败")
		return
	}
	if payJson.Quantity != 1 {
		err = gerror.New("Quantity验证失败")
		return
	}
	return
}
