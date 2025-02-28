// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityFeeds is the golang structure of table shiningu_community_feeds for DAO operations like Where/Data.
type CommunityFeeds struct {
	g.Meta    `orm:"table:shiningu_community_feeds, do:true"`
	Id        interface{} // 流水
	Uid       interface{} //
	PostId    interface{} // 帖子编号
	CreatedAt *gtime.Time // 创建时间
}
