// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberFriend is the golang structure for table member_friend.
type MemberFriend struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"当前用户"` // 当前用户
	Uid2      int64       `json:"uid_2"      orm:"uid2"       description:"对方编号"` // 对方编号
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
