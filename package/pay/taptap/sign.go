package taptap

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

type pTapTap struct {
	Secret string `json:"secret" dc:"秘钥"`
	//OrderId  string `json:"order_id" dc:"订单唯一 ID"`
	ClientId string `json:"client_id" dc:"应用的 Client ID"`
}

func New(clientId string, secret string) *pTapTap {
	return &pTapTap{
		Secret:   secret,
		ClientId: clientId,
		//OrderId:  orderId,
	}
}

// Sign signs the request.
func (p *pTapTap) Sign(req *http.Request, secret string) (string, error) {
	//获取请求参数
	//req := g.RequestFromCtx(ctx).Request
	return Sign(req, secret)
}

//func (p *pTapTap) SignOld(ctx context.Context, method, url string, token string, data any) (sign string, ts int64, nonce string, err error) {
//	//secret := p.Secret
//	//ts = gtime.Now().Unix()
//	//nonce = grand.S(5)
//	//header := http.Header{
//	//	"Content-Type": {"Content-Type: application/json; charset=utf-8"},
//	//	"X-Tap-Ts":     {strconv.FormatInt(ts, 10)},
//	//	"X-Tap-Nonce":  {nonce},
//	//}
//	//if method == "POST" {
//	//	header.Set("Content-Type", "application/json; charset=utf-8")
//	//}
//	////ctx := context.Background()
//	//request := g.RequestFromCtx(ctx).Request
//	//body, _ := json.Marshal(data)
//	////req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(string(body)))
//	//req.Header = header
//	//sign, err = Sign(req, secret)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//req.Header.Set("X-Tap-Sign", sign)
//	//return
//}

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
