// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsRecommend is the golang structure of table shiningu_community_posts_recommend for DAO operations like Where/Data.
type CommunityPostsRecommend struct {
	g.Meta    `orm:"table:shiningu_community_posts_recommend, do:true"`
	Pid       interface{} // 帖子编号
	Type      interface{} // 推荐类型
	CreatedAt *gtime.Time // 创建时间
	EndTime   *gtime.Time // 结束时间
}
