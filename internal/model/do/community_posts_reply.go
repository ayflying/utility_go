// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsReply is the golang structure of table shiningu_community_posts_reply for DAO operations like Where/Data.
type CommunityPostsReply struct {
	g.Meta    `orm:"table:shiningu_community_posts_reply, do:true"`
	Id        interface{} // 唯一id
	Pid       interface{} // 主贴id
	Uid       interface{} //
	Uid2      interface{} // 被回复者的uid
	Content   interface{} // 正文
	Extend    interface{} // 附加数据
	TopId     interface{} // 上级id
	CreatedAt *gtime.Time // 创建时间
	Sort      interface{} // 跟帖顺序
	AiScore   interface{} // 机器评分
	Status    interface{} // 状态
	At        interface{} // 用户的at功能
	Like      interface{} // 回复点赞
}
