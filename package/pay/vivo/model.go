package vivo

type TokenAuthResponse struct {
	ReturnCode int                    `json:"retcode"`
	Data       *TokenAuthResponseData `json:"data,omitempty"`
}

type TokenAuthResponseData struct {
	Success bool   `json:"success,omitempty"`
	OpenId  string `json:"openid,omitempty"`
}

type LoginType struct {
	Token   string `json:"token"`
	Ssoid   string `json:"ssoid"`
	Channel int    `json:"channel"`
	AdId    string `json:"adId"`
}

type PayCallback struct {
	AppId         string `json:"appId"`
	CpId          string `json:"cpId"`
	CpOrderNumber string `json:"cpOrderNumber"`
	ExtInfo       string `json:"extInfo"`
	OrderAmount   string `json:"orderAmount"`
	OrderNumber   string `json:"orderNumber"`
	PayTime       string `json:"payTime"`
	RespCode      string `json:"respCode"`
	RespMsg       string `json:"respMsg"`
	SignMethod    string `json:"signMethod"`
	Signature     string `json:"signature"`
	TradeStatus   string `json:"tradeStatus"`
	TradeType     string `json:"tradeType"`
}
