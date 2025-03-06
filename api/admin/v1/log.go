package v1

import (
	"github.com/ayflying/utility_go/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type SystemLog struct {
	entity.SystemLog
	Data g.Map `json:"data" dc:"操作数据"`
	//Post g.Map `json:"post" dc:"提交数据"`
}
