// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityFans is the golang structure of table shiningu_community_fans for DAO operations like Where/Data.
type CommunityFans struct {
	g.Meta    `orm:"table:shiningu_community_fans, do:true"`
	Uid       interface{} // 用户编号
	Fans      interface{} // 粉丝编号
	CreatedAt *gtime.Time // 关注时间
}
