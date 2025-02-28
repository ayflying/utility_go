// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityFans is the golang structure for table community_fans.
type CommunityFans struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户编号"` // 用户编号
	Fans      int64       `json:"fans"       orm:"fans"       description:"粉丝编号"` // 粉丝编号
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"关注时间"` // 关注时间
}
