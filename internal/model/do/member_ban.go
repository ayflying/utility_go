// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberBan is the golang structure of table shiningu_member_ban for DAO operations like Where/Data.
type MemberBan struct {
	g.Meta    `orm:"table:shiningu_member_ban, do:true"`
	Uid       interface{} // 用户编号
	CreatedAt *gtime.Time // 禁用时间
	Type      interface{} // 禁用类型
}
