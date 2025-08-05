package huawei

import (
	"net/http"
	"time"
)

const (
	AuthTokenUrl     = "https://oauth-api.cloud.huawei.com/rest.php?nsp_fmt=JSON&nsp_svc=huawei.oauth2.user.getTokenInfo"
	OrderUrl         = "https://orders-drcn.iap.hicloud.com/applications/purchases/tokens/verify"
	LocationShanghai = "Asia/Shanghai"

	RSA  = "RSA"
	RSA2 = "RSA2"

	OrderResponseOk = "0"
	PurchaseStateOk = 0
)

func getOrderUrl(accountFlag int) string {
	if accountFlag == 1 {
		// site for telecom carrier
		return "https://orders-at-dre.iap.dbankcloud.com"
	} else {
		// TODO: replace the (ip:port) to the real one
		return "http://exampleserver/_mockserver_"
	}

}

// default http client with 5 seconds timeout
var RequestHttpClient = http.Client{Timeout: time.Second * 5}
