// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberBan is the golang structure for table member_ban.
type MemberBan struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户编号"` // 用户编号
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"禁用时间"` // 禁用时间
	Type      int         `json:"type"       orm:"type"       description:"禁用类型"` // 禁用类型
}
