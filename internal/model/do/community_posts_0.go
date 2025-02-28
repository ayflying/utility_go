// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityPosts0 is the golang structure of table shiningu_community_posts_0 for DAO operations like Where/Data.
type CommunityPosts0 struct {
	g.Meta      `orm:"table:shiningu_community_posts_0, do:true"`
	Id          interface{} // 帖子编号
	Title       interface{} // 标题
	Content     interface{} // 帖子正文
	Images      interface{} // 帖子图片批量
	Image       interface{} // 图片
	ImagesRatio interface{} // 图片长宽比
	Like        interface{} // 点赞
	Collect     interface{} // 收藏
	Extend      interface{} // 附加信息
	At          interface{} // at
	Data        interface{} // 内容属性
	UseIds      interface{} // 最新期数
	Share       interface{} // 分享帖子
	AiScore     interface{} //
}
