// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityFeeds is the golang structure for table community_feeds.
type CommunityFeeds struct {
	Id        int         `json:"id"         orm:"id"         description:"流水"`   // 流水
	Uid       int64       `json:"uid"        orm:"uid"        description:""`     //
	PostId    int         `json:"post_id"    orm:"post_id"    description:"帖子编号"` // 帖子编号
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
