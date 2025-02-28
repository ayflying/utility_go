// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsHotDel is the golang structure of table shiningu_community_posts_hot_del for DAO operations like Where/Data.
type CommunityPostsHotDel struct {
	g.Meta    `orm:"table:shiningu_community_posts_hot_del, do:true"`
	Id        interface{} // 帖子编号
	CreatedAt *gtime.Time // 创建时间
}
