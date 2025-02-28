// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameStage is the golang structure of table shiningu_game_stage for DAO operations like Where/Data.
type GameStage struct {
	g.Meta    `orm:"table:shiningu_game_stage, do:true"`
	Uid       interface{} // 用户标识
	Chapter   interface{} // 章节
	WinData   interface{} // 通关过的数据
	StageData interface{} // 关卡数据
	Star      interface{} // 章节获得的总星
	RwdData   interface{} // 通关奖励领取
}
