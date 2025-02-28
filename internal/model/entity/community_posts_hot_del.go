// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsHotDel is the golang structure for table community_posts_hot_del.
type CommunityPostsHotDel struct {
	Id        int         `json:"id"         orm:"id"         description:"帖子编号"` // 帖子编号
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
