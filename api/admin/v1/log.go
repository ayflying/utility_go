package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"new-gitlab.adesk.com/public_project/utility_go/internal/model/entity"
)

type SystemLog struct {
	entity.SystemLog
	Data g.Map `json:"data" dc:"操作数据"`
	//Post g.Map `json:"post" dc:"提交数据"`
}
