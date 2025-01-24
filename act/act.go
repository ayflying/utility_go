package act

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type act struct {
}

type Sql struct {
	UID       uint64      `gorm:"primaryKey;comment:玩家编号"`
	ActID     int         `gorm:"primaryKey;comment:活动编号"`
	Action    string      `gorm:"type:text;comment:活动配置"`
	UpdatedAt *gtime.Time `gorm:"index;comment:更新时间"`
}

func New() *act {

	return &act{}
}

func (s *act) CreateTable(name string) {

	//prefix := g.DB().GetPrefix()
	//
	//g.DB()
	//g.DB().Exec(gctx.New())

}
