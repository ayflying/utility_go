// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPosts is the golang structure of table shiningu_community_posts for DAO operations like Where/Data.
type CommunityPosts struct {
	g.Meta       `orm:"table:shiningu_community_posts, do:true"`
	Id           interface{} // 帖子编号
	Tid          interface{} // 帖子栏目
	Uid          interface{} // 发布者
	Click        interface{} // 点击数
	LikeCount    interface{} // 点赞数量
	CollectCount interface{} // 收藏数量
	Popularity   interface{} // 人气度热度
	Charm        interface{} // 魅力值
	Language     interface{} // 语言
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
	Topic1       interface{} // 话题1
	Topic2       interface{} // 话题2
	Status       interface{} // 帖子状态 -1限流 0 正常
	Recommend    interface{} // 小编推荐
	Extend       interface{} // 附加信息
	At           interface{} // 用户的at功能
	Period       interface{} // 最新期数
	Price        interface{} // 当前价格
	Index        interface{} // 索引编号
	Gesture      interface{} // 手势
	CharacterNum interface{} // 角色数量
}
