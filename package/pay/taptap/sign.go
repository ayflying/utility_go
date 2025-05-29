package taptap

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type pTapTap struct {
	Secret   string `json:"secret" dc:"秘钥"`
	OrderId  string `json:"order_id" dc:"订单唯一 ID"`
	ClientId string `json:"client_id" dc:"应用的 Client ID"`
}

func New(orderId string) *pTapTap {
	return &pTapTap{
		Secret:   "5AFEWnadBA0NgJK2mxeBLQEde0qyIefxLSc4XKHsx9AwkQRhxzkQ9DixsOkK6gcV",
		ClientId: "mox88lbz43edfukdgk",
		OrderId:  orderId,
	}
}

func (p *pTapTap) Sign(url string, body []byte) (sign string, ts int64, nonce string, err error) {
	//nolint:gosec
	secret := p.Secret
	//body := gjson.MustEncode(g.Map{})
	//body := []byte(`{"event_type":"charge.succeeded","order":{"order_id":"1790288650833465345","purchase_token":"rT2Et9p0cfzq4fwjrTsGSacq0jQExFDqf5gTy1alp+Y=","client_id":"o6nD4iNavjQj75zPQk","open_id":"4+Axcl2RFgXbt6MZwdh++w==","user_region":"US","goods_open_id":"com.goods.open_id","goods_name":"TestGoodsName","status":"charge.succeeded","amount":"19000000000","currency":"USD","create_time":"1716168000","pay_time":"1716168000","extra":"1111111111111111111"}}`)
	//url := "https://example.com/my-service/v1/my-method"
	ts = gtime.Now().Unix()
	nonce = grand.S(5)
	method := "POST"
	header := http.Header{
		"Content-Type": {"Content-Type: application/json; charset=utf-8"},
		"X-Tap-Ts":     {strconv.FormatInt(ts, 10)},
		"X-Tap-Nonce":  {nonce},
	}
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	req.Header = header
	sign, err = Sign(req, secret)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Tap-Sign", sign)
	return
}

// Sign signs the request.
func Sign(req *http.Request, secret string) (string, error) {
	methodPart := req.Method
	urlPathAndQueryPart := req.URL.RequestURI()
	headersPart, err := getHeadersPart(req.Header)
	if err != nil {
		return "", err
	}
	bodyPart, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	signParts := methodPart + "\n" + urlPathAndQueryPart + "\n" + headersPart + "\n" + string(bodyPart) + "\n"
	fmt.Println(signParts)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signParts))
	rawSign := h.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(rawSign)
	return sign, nil
}

// getHeadersPart returns the headers part of the request.
func getHeadersPart(header http.Header) (string, error) {
	var headerKeys []string
	for k, v := range header {
		k = strings.ToLower(k)
		if !strings.HasPrefix(k, "x-tap-") {
			continue
		}
		if k == "x-tap-sign" {
			continue
		}
		if len(v) > 1 {
			return "", fmt.Errorf("invalid header, %q has multiple values", k)
		}
		headerKeys = append(headerKeys, k)
	}
	sort.Strings(headerKeys)
	headers := make([]string, 0, len(headerKeys))
	for _, k := range headerKeys {
		headers = append(headers, fmt.Sprintf("%s:%s", k, header.Get(k)))
	}
	return strings.Join(headers, "\n"), nil
}
