// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsRecommend is the golang structure for table community_posts_recommend.
type CommunityPostsRecommend struct {
	Pid       int         `json:"pid"        orm:"pid"        description:"帖子编号"` // 帖子编号
	Type      int         `json:"type"       orm:"type"       description:"推荐类型"` // 推荐类型
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"结束时间"` // 结束时间
}
