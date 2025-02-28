// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameAct is the golang structure of table shiningu_game_act for DAO operations like Where/Data.
type GameAct struct {
	g.Meta    `orm:"table:shiningu_game_act, do:true"`
	Uid       interface{} // 玩家编号
	ActId     interface{} // 活动编号
	Action    interface{} // 活动配置
	UpdatedAt *gtime.Time // 更新时间
}
