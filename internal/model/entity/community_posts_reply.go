// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityPostsReply is the golang structure for table community_posts_reply.
type CommunityPostsReply struct {
	Id        int         `json:"id"         orm:"id"         description:"唯一id"`     // 唯一id
	Pid       int         `json:"pid"        orm:"pid"        description:"主贴id"`     // 主贴id
	Uid       int64       `json:"uid"        orm:"uid"        description:""`         //
	Uid2      int64       `json:"uid_2"      orm:"uid2"       description:"被回复者的uid"` // 被回复者的uid
	Content   string      `json:"content"    orm:"content"    description:"正文"`       // 正文
	Extend    string      `json:"extend"     orm:"extend"     description:"附加数据"`     // 附加数据
	TopId     int         `json:"top_id"     orm:"top_id"     description:"上级id"`     // 上级id
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`     // 创建时间
	Sort      int         `json:"sort"       orm:"sort"       description:"跟帖顺序"`     // 跟帖顺序
	AiScore   float64     `json:"ai_score"   orm:"ai_score"   description:"机器评分"`     // 机器评分
	Status    int         `json:"status"     orm:"status"     description:"状态"`       // 状态
	At        string      `json:"at"         orm:"at"         description:"用户的at功能"`  // 用户的at功能
	Like      string      `json:"like"       orm:"like"       description:"回复点赞"`     // 回复点赞
}
