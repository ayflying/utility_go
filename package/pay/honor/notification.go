package honor

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
)

func (p *Pay) Notification(r *http.Request) {

}

// ConsumeProduct 商品消耗
func (p *Pay) ConsumeProduct(purchaseToken string) (err error) {
	url := Host + "/iap/server/consumeProduct"
	_, err = g.Client().ContentJson().Post(gctx.New(), url, g.Map{
		"purchaseToken":      purchaseToken,
		"developerChallenge": "",
	})
	if err != nil {

		return
	}

	return
}
