// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberUser is the golang structure of table shiningu_member_user for DAO operations like Where/Data.
type MemberUser struct {
	g.Meta         `orm:"table:shiningu_member_user, do:true"`
	Uid            interface{} // 用户标识
	Guid           interface{} // 用户本次登录标识
	Gid            interface{} // 用户组编号
	AccountLogin   interface{} // 社交账号登录
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	Nickname       interface{} // 昵称
	Phone          interface{} // 绑定手机
	Email          interface{} // 绑定邮箱
	Money          interface{} // 充值货币
	Save           interface{} // 储存路径
	Slots          interface{} // 槽位数量
	OnlineDuration interface{} // 在线时长
	OnlineStatus   interface{} // 在线状态
	OnlineTime     *gtime.Time // 上线时间
	OfflineTime    *gtime.Time // 离线时间
	CreateIp       interface{} // 创号ip地址
	UpdateIp       interface{} // 更新IP地址
	Level          interface{} // 等级
	Exp            interface{} // 经验
	Title          interface{} // 称号
	Avatar         interface{} // 头像
	AvatarFrame    interface{} // 头像框
	Popularity     interface{} // 人气度
	Charm          interface{} // 魅力值
	Gift           interface{} // 礼物值
}
