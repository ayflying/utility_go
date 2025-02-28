// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityGift is the golang structure of table shiningu_community_gift for DAO operations like Where/Data.
type CommunityGift struct {
	g.Meta    `orm:"table:shiningu_community_gift, do:true"`
	Id        interface{} //
	Uid       interface{} // 收礼玩家编号
	FromUid   interface{} // 送礼玩家编号
	Type      interface{} // 送礼类型
	Pid       interface{} // 帖子编号
	ItemId    interface{} // 礼物编号
	Count     interface{} // 礼物数量
	CreatedAt *gtime.Time // 创建时间
}
