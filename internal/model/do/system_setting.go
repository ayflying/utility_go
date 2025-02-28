// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SystemSetting is the golang structure of table shiningu_system_setting for DAO operations like Where/Data.
type SystemSetting struct {
	g.Meta `orm:"table:shiningu_system_setting, do:true"`
	Name   interface{} // 配置名称
	Value  interface{} // 配置详情
	Type   interface{} // 类型
}
