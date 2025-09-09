package honor

import (
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/grand"

	"net/http"
)

// 响应结果结构体
type Response struct {
	Code    int          `json:"code"`    // 结果码 0: 成功，其他: 失败
	Message string       `json:"message"` // 错误信息
	Data    *DataContent `json:"data"`    // 包含购买信息的结构体
}

// 数据内容结构体，对应data字段
type DataContent struct {
	PurchaseProductInfo string `json:"purchaseProductInfo"` // 消耗结果数据的JSON字符串
	DataSig             string `json:"dataSig"`             // purchaseProductInfo的签名
	SigAlgorithm        string `json:"sigAlgorithm"`        // 签名算法，云侧加密算法为"RSA"
}

func (p *Pay) Notification(r *http.Request) {

}

// ConsumeProduct 商品消耗
func (p *Pay) ConsumeProduct(purchaseToken string) (err error) {
	url := Host + "/iap/server/consumeProduct"
	get, err := g.Client().ContentJson().Post(gctx.New(), url, g.Map{
		"purchaseToken":      purchaseToken,
		"developerChallenge": grand.S(16),
	})
	if err != nil {

		return
	}

	var res *Response
	gjson.DecodeTo(get.ReadAllString(), &res)
	if res.Code != 0 {
		g.Log().Error(gctx.New(), "商品消耗失败: "+res.Message)
		return errors.New(res.Message)
	}

	return
}
