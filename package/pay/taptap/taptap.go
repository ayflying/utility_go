package taptap

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"strconv"
)

type GetPayData struct {
	Data struct {
		Order struct {
		} `json:"order"`
	} `json:"data"`
	Success bool `json:"success"`
}

//查询订单信息
func (p *pTapTap) Info(orderId string, clientId string, token []byte) (res string, err error) {
	url := fmt.Sprintf("https://cloud-payment.tapapis.cn/order/v1/info?client_id=%v&order_id=%v", orderId, clientId)
	res, err = p.get(url, token)
	return
}

//验证并核销订单
func (p *pTapTap) Verify(orderId string, clientId string, token []byte) (res string, err error) {
	url := fmt.Sprintf("https://cloud-payment.tapapis.cn/order/v1/verify?client_id=%v", clientId)
	res, err = p.get(url, token)
	return
}

func (p *pTapTap) get(url string, token []byte, _data ...any) (res string, err error) {
	sign, ts, nonce, err := p.Sign(url, token)
	if err != nil {
		return
	}
	var _get *gclient.Response
	if len(_data) == 0 {
		_get, err = g.Client().Header(map[string]string{
			"X-Tap-Sign":  sign,
			"X-Tap-Nonce": nonce,
			"X-Tap-Ts":    strconv.FormatInt(ts, 10),
		}).Get(gctx.New(), url)

	} else {
		_get, err = g.Client().Header(map[string]string{
			"X-Tap-Sign":  sign,
			"X-Tap-Nonce": nonce,
			"X-Tap-Ts":    strconv.FormatInt(ts, 10),
		}).Post(gctx.New(), url, _data[0])
	}

	if err != nil {
		return
	}
	res = _get.ReadAllString()
	return
}
