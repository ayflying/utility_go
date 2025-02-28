// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameAct is the golang structure for table game_act.
type GameAct struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"玩家编号"` // 玩家编号
	ActId     int         `json:"act_id"     orm:"act_id"     description:"活动编号"` // 活动编号
	Action    string      `json:"action"     orm:"action"     description:"活动配置"` // 活动配置
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
}
