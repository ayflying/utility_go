// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberUser is the golang structure for table member_user.
type MemberUser struct {
	Uid            int64       `json:"uid"             orm:"uid"             description:"用户标识"`     // 用户标识
	Guid           string      `json:"guid"            orm:"guid"            description:"用户本次登录标识"` // 用户本次登录标识
	Gid            int         `json:"gid"             orm:"gid"             description:"用户组编号"`    // 用户组编号
	AccountLogin   string      `json:"account_login"   orm:"account_login"   description:"社交账号登录"`   // 社交账号登录
	CreatedAt      *gtime.Time `json:"created_at"      orm:"created_at"      description:"创建时间"`     // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"      orm:"updated_at"      description:"更新时间"`     // 更新时间
	DeletedAt      *gtime.Time `json:"deleted_at"      orm:"deleted_at"      description:"删除时间"`     // 删除时间
	Nickname       string      `json:"nickname"        orm:"nickname"        description:"昵称"`       // 昵称
	Phone          string      `json:"phone"           orm:"phone"           description:"绑定手机"`     // 绑定手机
	Email          string      `json:"email"           orm:"email"           description:"绑定邮箱"`     // 绑定邮箱
	Money          int64       `json:"money"           orm:"money"           description:"充值货币"`     // 充值货币
	Save           string      `json:"save"            orm:"save"            description:"储存路径"`     // 储存路径
	Slots          string      `json:"slots"           orm:"slots"           description:"槽位数量"`     // 槽位数量
	OnlineDuration int64       `json:"online_duration" orm:"online_duration" description:"在线时长"`     // 在线时长
	OnlineStatus   int         `json:"online_status"   orm:"online_status"   description:"在线状态"`     // 在线状态
	OnlineTime     *gtime.Time `json:"online_time"     orm:"online_time"     description:"上线时间"`     // 上线时间
	OfflineTime    *gtime.Time `json:"offline_time"    orm:"offline_time"    description:"离线时间"`     // 离线时间
	CreateIp       string      `json:"create_ip"       orm:"create_ip"       description:"创号ip地址"`   // 创号ip地址
	UpdateIp       string      `json:"update_ip"       orm:"update_ip"       description:"更新IP地址"`   // 更新IP地址
	Level          int         `json:"level"           orm:"level"           description:"等级"`       // 等级
	Exp            int64       `json:"exp"             orm:"exp"             description:"经验"`       // 经验
	Title          string      `json:"title"           orm:"title"           description:"称号"`       // 称号
	Avatar         string      `json:"avatar"          orm:"avatar"          description:"头像"`       // 头像
	AvatarFrame    int         `json:"avatar_frame"    orm:"avatar_frame"    description:"头像框"`      // 头像框
	Popularity     int         `json:"popularity"      orm:"popularity"      description:"人气度"`      // 人气度
	Charm          int         `json:"charm"           orm:"charm"           description:"魅力值"`      // 魅力值
	Gift           int         `json:"gift"            orm:"gift"            description:"礼物值"`      // 礼物值
}
