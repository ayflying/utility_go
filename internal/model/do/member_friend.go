// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberFriend is the golang structure of table shiningu_member_friend for DAO operations like Where/Data.
type MemberFriend struct {
	g.Meta    `orm:"table:shiningu_member_friend, do:true"`
	Uid       interface{} // 当前用户
	Uid2      interface{} // 对方编号
	CreatedAt *gtime.Time // 创建时间
}
