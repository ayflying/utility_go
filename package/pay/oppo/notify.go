package oppo

import (
	"net/http"
	"net/url"
)

func (p *OppoType) ParseNotifyToBodyMap(req *http.Request) (bm map[string]interface{}, err error) {
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	var form map[string][]string = req.Form
	bm = make(map[string]interface{}, len(form)+1)
	for k, v := range form {
		if len(v) == 1 {
			bm[k] = v[0]
			//bm.Set(k, v[0])
		}
	}
	return
}

func (p *OppoType) ParseNotifyByURLValues(value url.Values) (bm map[string]interface{}, err error) {
	bm = make(map[string]interface{}, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			bm[k] = v[0]
			//bm.Set(k, v[0])
		}
	}
	return
}
