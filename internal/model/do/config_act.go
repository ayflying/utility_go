// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigAct is the golang structure of table shiningu_config_act for DAO operations like Where/Data.
type ConfigAct struct {
	g.Meta    `orm:"table:shiningu_config_act, do:true"`
	Id        interface{} // 流水编号
	Type      interface{} // 活动类型
	Actid     interface{} // 活动编号
	Name      interface{} // 活动名称
	Hid       interface{} // 活动标识
	Data      interface{} // 活动数据
	StartTime *gtime.Time // 开始时间
	EndTime   *gtime.Time // 结束时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
