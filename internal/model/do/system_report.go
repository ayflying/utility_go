// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemReport is the golang structure of table shiningu_system_report for DAO operations like Where/Data.
type SystemReport struct {
	g.Meta    `orm:"table:shiningu_system_report, do:true"`
	Id        interface{} //
	Rid       interface{} // 举报id
	Uid       interface{} // 举报人编号
	Type      interface{} // 举报类型
	Desc      interface{} // 举报正文
	CreatedAt *gtime.Time // 举报时间
	DeletedAt *gtime.Time // 删除时间
	Status    interface{} // 处理状态
}
