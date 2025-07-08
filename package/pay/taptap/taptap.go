package taptap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"net/http"
)

type GetPayData struct {
	Data struct {
		Order struct {
			Order
		} `json:"order"`
	} `json:"data"`
	Success bool `json:"success"`
}

type Order struct {
	OrderID       string        `json:"order_id"`       // 订单唯一 ID
	PurchaseToken string        `json:"purchase_token"` // 用于订单核销的 token
	ClientID      string        `json:"client_id"`      // 应用的 Client ID
	OpenID        string        `json:"open_id"`        // 用户的开放平台 ID
	UserRegion    string        `json:"user_region"`    // 用户地区
	GoodsOpenID   string        `json:"goods_open_id"`  // 商品唯一 ID
	GoodsName     string        `json:"goods_name"`     // 商品名称
	Status        PaymentStatus `json:"status"`         // 订单状态
	Amount        string        `json:"amount"`         // 金额（本币金额 x 1,000,000）
	Currency      string        `json:"currency"`       // 币种
	CreateTime    string        `json:"create_time"`    // 创建时间
	PayTime       string        `json:"pay_time"`       // 支付时间
	Extra         string        `json:"extra"`          // 商户自定义数据，如角色信息等，长度不超过 255 UTF-8 字符
}
type PaymentStatus string

const (
	ChargePending   PaymentStatus = "charge.pending"   // 待支付
	ChargeSucceeded PaymentStatus = "charge.succeeded" //支付成功
	ChargeConfirmed PaymentStatus = "charge.confirmed" //已核销
	ChargeOverdue   PaymentStatus = "charge.overdue"   //支付超时关闭
	RefundPending   PaymentStatus = "refund.pending"   //退款中
	RefundSucceeded PaymentStatus = "refund.succeeded" //退款成功
	RefundFailed    PaymentStatus = "refund.failed"    //退款失败
	RefundRejected  PaymentStatus = "refund.rejected"  //退款被拒绝
)

// 查询订单信息
func (p *pTapTap) Info(ctx context.Context, order string) (getPayData *GetPayData, err error) {
	url := fmt.Sprintf("https://cloud-payment.tapapis.cn/order/v1/info?client_id=%v&order_id=%v", p.ClientId, order)
	getPayData, err = p.get(ctx, url)
	if err != nil {
		return
	}
	return
}

// 验证并核销订单
func (p *pTapTap) Verify(ctx context.Context, req any) (getPayData *GetPayData, err error) {
	url := fmt.Sprintf("https://cloud-payment.tapapis.cn/order/v1/verify?client_id=%v", p.ClientId)
	getPayData, err = p.get(ctx, url, req)
	if err != nil {
		return
	}
	return
}

func (p *pTapTap) get(ctx context.Context, url string, _data ...any) (getPayData *GetPayData, err error) {

	var _get *gclient.Response

	var header = map[string]string{
		"Content-Type": "Content-Type: application/json; charset=utf-8",
		"X-Tap-Nonce":  grand.S(6),
		"X-Tap-Ts":     gtime.Now().TimestampStr(),
	}
	ctx2 := context.Background()
	var method = "GET"
	if len(_data) > 0 {
		method = "POST"

	}

	//temp := []byte(`{"event_type":"charge.succeeded","order":{"order_id":"1790288650833465345","purchase_token":"rT2Et9p0cfzq4fwjrTsGSacq0jQExFDqf5gTy1alp+Y=","client_id":"o6nD4iNavjQj75zPQk","open_id":"4+Axcl2RFgXbt6MZwdh++w==","user_region":"US","goods_open_id":"com.goods.open_id","goods_name":"TestGoodsName","status":"charge.succeeded","amount":"19000000000","currency":"USD","create_time":"1716168000","pay_time":"1716168000","extra":"1111111111111111111"}}`)
	var body io.Reader
	if len(_data) > 0 {
		body = bytes.NewBuffer(gjson.MustEncode(_data[0]))
	} else {
		body = bytes.NewBuffer([]byte{})
	}
	req, _ := http.NewRequestWithContext(ctx2, method, url, body)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	sign, err2 := p.Sign(req, p.Secret)
	if err2 != nil {
		err = err2
		return
	}
	req.Header.Set("X-Tap-Sign", sign)
	header["X-Tap-Sign"] = sign
	if len(_data) == 0 {
		_get, err = g.Client().Header(header).ContentJson().Get(gctx.New(), url)
	} else {
		_get, err = g.Client().Header(header).ContentJson().Post(gctx.New(), url, _data[0])
	}

	if err != nil {
		return
	}
	getPayData = &GetPayData{}
	resData := _get.ReadAll()
	g.Dump(resData)
	if err = json.Unmarshal(resData, &getPayData); err != nil {
		return
	}
	if !getPayData.Success {
		err = errors.New(string(resData))
	}
	return
}
