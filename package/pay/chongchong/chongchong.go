package chongchong

import (
	"fmt"
	"github.com/ayflying/utility_go/package/pay/common"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

//验单
func (p *Pay) Verify(req *CallbackData, sign string) (isOk bool, err error) {
	//req := g.RequestFromCtx(ctx).Request
	//data, err := common.ParseNotifyToBodyMap(req)

	var data = map[string]interface{}{
		"orderPrice":           req.OrderPrice,
		"packageId":            req.PackageId,
		"partnerTransactionNo": req.PartnerTransactionNo,
		"productId":            req.ProductId,
		"statusCode":           req.StatusCode,
		"transactionNo":        req.TransactionNo,
	}

	dataStr, err := common.BuildSignStr(data)

	var SingStr = fmt.Sprintf("%v&%v", dataStr, p.ApiKey)
	sign2, err := gmd5.EncryptString(SingStr)

	if sign == sign2 {
		isOk = true
	}
	return
}
