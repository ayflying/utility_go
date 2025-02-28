// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameKv is the golang structure for table game_kv.
type GameKv struct {
	Uid       int         `json:"uid"        orm:"uid"        description:"用户"`   // 用户
	Kv        string      `json:"kv"         orm:"kv"         description:"变量"`   // 变量
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
}
