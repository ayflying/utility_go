// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberLimit is the golang structure for table member_limit.
type MemberLimit struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户uid"` // 用户uid
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`  // 创建时间
	Data      string      `json:"data"       orm:"data"       description:"玩家权限"`  // 玩家权限
}
