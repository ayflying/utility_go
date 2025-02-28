// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityGift is the golang structure for table community_gift.
type CommunityGift struct {
	Id        int         `json:"id"         orm:"id"         description:""`       //
	Uid       int64       `json:"uid"        orm:"uid"        description:"收礼玩家编号"` // 收礼玩家编号
	FromUid   int64       `json:"from_uid"   orm:"fromUid"    description:"送礼玩家编号"` // 送礼玩家编号
	Type      int         `json:"type"       orm:"type"       description:"送礼类型"`   // 送礼类型
	Pid       int         `json:"pid"        orm:"pid"        description:"帖子编号"`   // 帖子编号
	ItemId    int64       `json:"item_id"    orm:"itemId"     description:"礼物编号"`   // 礼物编号
	Count     int         `json:"count"      orm:"count"      description:"礼物数量"`   // 礼物数量
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
}
