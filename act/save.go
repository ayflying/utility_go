package act

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// 查询表格时候存在
func (s *act) TableExists(name string) bool {
	Prefix := g.DB().GetPrefix()
	get, err := g.DB().TableFields(gctx.New(), Prefix+name)
	if err != nil {
		g.Log().Error(gctx.New(), err)
	}
	if get != nil {
		return true
	}
	return false
}
