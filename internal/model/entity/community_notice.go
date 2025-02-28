// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommunityNotice is the golang structure for table community_notice.
type CommunityNotice struct {
	Id        int         `json:"id"         orm:"id"         description:""`       //
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户编号"`   // 用户编号
	FromUid   int64       `json:"from_uid"   orm:"from_uid"   description:"对方用户编号"` // 对方用户编号
	Type      int         `json:"type"       orm:"type"       description:"类型"`     // 类型
	Sid       int         `json:"sid"        orm:"sid"        description:"来源编号"`   // 来源编号
	Content   string      `json:"content"    orm:"content"    description:"正文"`     // 正文
	Extend    string      `json:"extend"     orm:"extend"     description:"附加属性"`   // 附加属性
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`   // 更新时间
	Status    int         `json:"status"     orm:"status"     description:"通知状态"`   // 通知状态
}
