// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameRank is the golang structure for table game_rank.
type GameRank struct {
	Key       int         `json:"key"        orm:"key"        description:""`         //
	RankId    int         `json:"rank_id"    orm:"rank_id"    description:"排行榜编号"`    // 排行榜编号
	Type      int         `json:"type"       orm:"type"       description:"排行榜类型"`    // 排行榜类型
	Data      string      `json:"data"       orm:"data"       description:"数据"`       // 数据
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"排行榜创建时间"`  // 排行榜创建时间
	StartTime *gtime.Time `json:"start_time" orm:"start_time" description:"榜单开始时间"`   // 榜单开始时间
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"榜单结束时间"`   // 榜单结束时间
	Status    int         `json:"status"     orm:"status"     description:""`         //
	Image     string      `json:"image"      orm:"image"      description:"结算封面"`     // 结算封面
	FirstUid  int64       `json:"first_uid"  orm:"first_uid"  description:"第一名的用户id"` // 第一名的用户id
}
