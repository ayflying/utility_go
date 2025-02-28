// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigAct is the golang structure for table config_act.
type ConfigAct struct {
	Id        int         `json:"id"         orm:"id"         description:"流水编号"` // 流水编号
	Type      int         `json:"type"       orm:"type"       description:"活动类型"` // 活动类型
	Actid     int         `json:"actid"      orm:"actid"      description:"活动编号"` // 活动编号
	Name      string      `json:"name"       orm:"name"       description:"活动名称"` // 活动名称
	Hid       string      `json:"hid"        orm:"hid"        description:"活动标识"` // 活动标识
	Data      string      `json:"data"       orm:"data"       description:"活动数据"` // 活动数据
	StartTime *gtime.Time `json:"start_time" orm:"start_time" description:"开始时间"` // 开始时间
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"结束时间"` // 结束时间
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
}
