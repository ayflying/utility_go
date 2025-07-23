package vivo

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"net/url"
)

func (p *Pay) AuthToken(bm g.Map) (rsp *TokenAuthResponse, err error) {
	if _, ok := bm["opentoken"]; !ok {
		return
	}
	//err = bm.CheckEmptyError("opentoken")
	if err != nil {
		return nil, err
	}
	bs, err := p.doAuthToken(bm)
	if err != nil {
		return nil, err
	}
	rsp = new(TokenAuthResponse)
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	return rsp, nil
}

func (p *Pay) doAuthToken(bm g.Map) (bs []byte, err error) {
	param := p.FormatURLParam(bm)
	//httpClient := xhttp.NewClient()
	//res, bs, errs := httpClient.Type(xhttp.TypeFormData).Post(AuthTokenUrl).SendString(param).EndBytes()
	res, err := g.Client().Post(gctx.New(), AuthTokenUrl, param)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return res.ReadAll(), nil
}

// 格式化请求URL参数
func (p *Pay) FormatURLParam(body g.Map) (urlParam string) {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
