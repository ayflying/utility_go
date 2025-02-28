// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberLimit is the golang structure of table shiningu_member_limit for DAO operations like Where/Data.
type MemberLimit struct {
	g.Meta    `orm:"table:shiningu_member_limit, do:true"`
	Uid       interface{} // 用户uid
	CreatedAt *gtime.Time // 创建时间
	Data      interface{} // 玩家权限
}
