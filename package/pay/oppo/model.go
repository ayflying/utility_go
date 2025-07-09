package oppo

// OPPO支付回调参数结构体
type PayCallback struct {
	NotifyId     string `json:"notifyId" dc:"回调通知单号，以GC开头，必填，示例:GC20230314657000"`
	PartnerOrder string `json:"partnerOrder" dc:"开发者订单号，必填，示例:123456"`
	ProductName  string `json:"productName" dc:"商品名称，必填，示例:10元宝"`
	ProductDesc  string `json:"productDesc" dc:"商品描述，必填，示例:10元宝等于1元"`
	Price        int64  `json:"price" dc:"商品价格，单位为分，需要游戏服务端做验证，必填，示例:100"`
	Count        int    `json:"count" dc:"商品数量（一般为1），必填，示例:1"`
	Attach       string `json:"attach" dc:"请求支付时上传的附加参数，可能为空，选填"`
	Sign         string `json:"sign" dc:"OPPO服务端签名，需要游戏服务端做验证，必填"`
}
