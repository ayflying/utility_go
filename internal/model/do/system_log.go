// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLog is the golang structure of table shiningu_system_log for DAO operations like Where/Data.
type SystemLog struct {
	g.Meta    `orm:"table:shiningu_system_log, do:true"`
	Id        interface{} // 主键
	Uid       interface{} // 操作的用户
	Url       interface{} // 当前访问的url
	Data      interface{} // 操作数据
	CreatedAt *gtime.Time // 创建时间
	Ip        interface{} // 当前ip地址
}
