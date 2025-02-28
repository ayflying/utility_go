// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameKv is the golang structure of table shiningu_game_kv for DAO operations like Where/Data.
type GameKv struct {
	g.Meta    `orm:"table:shiningu_game_kv, do:true"`
	Uid       interface{} // 用户
	Kv        interface{} // 变量
	UpdatedAt *gtime.Time // 更新时间
}
