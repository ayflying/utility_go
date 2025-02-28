// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPosts is the golang structure for table community_posts.
type CommunityPosts struct {
	Id           int         `json:"id"            orm:"id"            description:"帖子编号"`           // 帖子编号
	Tid          int         `json:"tid"           orm:"tid"           description:"帖子栏目"`           // 帖子栏目
	Uid          int64       `json:"uid"           orm:"uid"           description:"发布者"`            // 发布者
	Click        int64       `json:"click"         orm:"click"         description:"点击数"`            // 点击数
	LikeCount    int         `json:"like_count"    orm:"like_count"    description:"点赞数量"`           // 点赞数量
	CollectCount int         `json:"collect_count" orm:"collect_count" description:"收藏数量"`           // 收藏数量
	Popularity   int         `json:"popularity"    orm:"popularity"    description:"人气度热度"`          // 人气度热度
	Charm        int         `json:"charm"         orm:"charm"         description:"魅力值"`            // 魅力值
	Language     string      `json:"language"      orm:"language"      description:"语言"`             // 语言
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`           // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"更新时间"`           // 更新时间
	DeletedAt    *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:"删除时间"`           // 删除时间
	Topic1       int         `json:"topic_1"       orm:"topic1"        description:"话题1"`            // 话题1
	Topic2       int         `json:"topic_2"       orm:"topic2"        description:"话题2"`            // 话题2
	Status       int         `json:"status"        orm:"status"        description:"帖子状态 -1限流 0 正常"` // 帖子状态 -1限流 0 正常
	Recommend    int         `json:"recommend"     orm:"recommend"     description:"小编推荐"`           // 小编推荐
	Extend       string      `json:"extend"        orm:"extend"        description:"附加信息"`           // 附加信息
	At           string      `json:"at"            orm:"at"            description:"用户的at功能"`        // 用户的at功能
	Period       int         `json:"period"        orm:"period"        description:"最新期数"`           // 最新期数
	Price        int         `json:"price"         orm:"price"         description:"当前价格"`           // 当前价格
	Index        int         `json:"index"         orm:"index"         description:"索引编号"`           // 索引编号
	Gesture      int         `json:"gesture"       orm:"gesture"       description:"手势"`             // 手势
	CharacterNum int         `json:"character_num" orm:"character_num" description:"角色数量"`           // 角色数量
}
