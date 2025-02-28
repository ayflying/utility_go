// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityPostsDetails is the golang structure of table shiningu_community_posts_details for DAO operations like Where/Data.
type CommunityPostsDetails struct {
	g.Meta     `orm:"table:shiningu_community_posts_details, do:true"`
	Id         interface{} //
	Attachment interface{} // 帖子附件
	SlotsImg   interface{} // 槽位图片
}
