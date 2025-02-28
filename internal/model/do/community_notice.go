// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityNotice is the golang structure of table shiningu_community_notice for DAO operations like Where/Data.
type CommunityNotice struct {
	g.Meta    `orm:"table:shiningu_community_notice, do:true"`
	Id        interface{} //
	Uid       interface{} // 用户编号
	FromUid   interface{} // 对方用户编号
	Type      interface{} // 类型
	Sid       interface{} // 来源编号
	Content   interface{} // 正文
	Extend    interface{} // 附加属性
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Status    interface{} // 通知状态
}
