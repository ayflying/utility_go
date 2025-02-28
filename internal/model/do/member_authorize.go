// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberAuthorize is the golang structure of table shiningu_member_authorize for DAO operations like Where/Data.
type MemberAuthorize struct {
	g.Meta    `orm:"table:shiningu_member_authorize, do:true"`
	Code      interface{} // 授权码
	Uid       interface{} // 用户标识
	Type      interface{} // 认证方式
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	CreateIp  interface{} // 创建ip
}
