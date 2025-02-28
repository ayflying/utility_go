// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberAuthorize is the golang structure for table member_authorize.
type MemberAuthorize struct {
	Code      string      `json:"code"       orm:"code"       description:"授权码"`  // 授权码
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户标识"` // 用户标识
	Type      string      `json:"type"       orm:"type"       description:"认证方式"` // 认证方式
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
	CreateIp  string      `json:"create_ip"  orm:"create_ip"  description:"创建ip"` // 创建ip
}
