package vivo

import (
	"context"
	"errors"
	"github.com/ayflying/utility_go/package/pay/common"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"sort"
	"strings"
)

func (p *Pay) VerifySign(ctx context.Context, key string) bool {
	bm, _ := common.ParseNotifyToBodyMap(g.RequestFromCtx(ctx).Request)
	signature := bm["signature"]
	delete(bm, "signature")
	delete(bm, "signMethod")
	sign := p.sign(bm, key)
	return signature == sign
}

func (p *Pay) sign(bm g.Map, key string) string {
	s, _ := p.buildSignStr(bm)
	s += "&" + gmd5.MustEncrypt(key)
	return gmd5.MustEncrypt(s)
}

func (p *Pay) buildSignStr(bm g.Map) (string, error) {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm[k]; v != "" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(gconv.String(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return "", errors.New("length is error")
	}
	return buf.String()[:buf.Len()-1], nil
}
