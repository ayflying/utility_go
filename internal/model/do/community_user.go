// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityUser is the golang structure of table shiningu_community_user for DAO operations like Where/Data.
type CommunityUser struct {
	g.Meta    `orm:"table:shiningu_community_user, do:true"`
	Uid       interface{} //
	Like      interface{} // 点赞
	Like2     interface{} // 帖子回复点赞
	Collect   interface{} // 收藏
	FollowNum interface{} // 关注数量
	FansNum   interface{} // 粉丝数量
	Gift      interface{} // 礼物值
	Blacklist interface{} // 黑名单
}
