// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameRank is the golang structure of table shiningu_game_rank for DAO operations like Where/Data.
type GameRank struct {
	g.Meta    `orm:"table:shiningu_game_rank, do:true"`
	Key       interface{} //
	RankId    interface{} // 排行榜编号
	Type      interface{} // 排行榜类型
	Data      interface{} // 数据
	CreatedAt *gtime.Time // 排行榜创建时间
	StartTime *gtime.Time // 榜单开始时间
	EndTime   *gtime.Time // 榜单结束时间
	Status    interface{} //
	Image     interface{} // 结算封面
	FirstUid  interface{} // 第一名的用户id
}
