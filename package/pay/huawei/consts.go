package huawei

import (
	"net/http"
	"time"
)

const (
	TokenUrl = "https://oauth-login.cloud.huawei.com/oauth2/v3/token"
)

func getOrderUrl(accountFlag int) string {
	if accountFlag == 1 {
		// site for telecom carrier
		//return "https://orders-at-dre.iap.dbankcloud.com"
		return "https://orders-drcn.iap.cloud.huawei.com.cn"
	} else {
		// TODO: replace the (ip:port) to the real one
		return "http://exampleserver/_mockserver_"
	}

}

// default http client with 5 seconds timeout
var RequestHttpClient = http.Client{Timeout: time.Second * 5}
