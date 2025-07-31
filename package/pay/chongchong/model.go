package chongchong

type Pay struct {
	ApiKey string `json:"api_key"`
}

func New(pay *Pay) *Pay {
	return &Pay{
		ApiKey: pay.ApiKey,
	}
}

// CallbackData 用于处理回调数据的结构体
type CallbackData struct {
	TransactionNo        string  `json:"transactionNo" dc:"平台交易单号，唯一标识一笔交易"`
	PartnerTransactionNo string  `json:"partnerTransactionNo" dc:"合作方交易单号，由合作方生成"`
	StatusCode           string  `json:"statusCode" dc:"交易状态码，SUCCESS表示成功，FAIL表示失败"`
	ProductId            int     `json:"productId" dc:"产品ID，对应后台配置的商品"`
	OrderPrice           float64 `json:"orderPrice" dc:"订单金额，单位为元"`
	PackageId            int     `json:"packageId" dc:"套餐ID，可选字段，部分商品有套餐区分"`
	ProductName          string  `json:"productName" dc:"产品名称，展示用"`
	ExtParam             string  `json:"extParam" dc:"扩展参数，回调时原样返回"`
	UserId               int     `json:"userId" dc:"用户ID，标识购买者"`
	Sign                 string  `json:"sign" dc:"签名，用于验证请求合法性"`
}
