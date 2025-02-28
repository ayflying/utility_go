// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GameStage is the golang structure for table game_stage.
type GameStage struct {
	Uid       int64  `json:"uid"        orm:"uid"        description:"用户标识"`    // 用户标识
	Chapter   int    `json:"chapter"    orm:"chapter"    description:"章节"`      // 章节
	WinData   string `json:"win_data"   orm:"win_data"   description:"通关过的数据"`  // 通关过的数据
	StageData string `json:"stage_data" orm:"stage_data" description:"关卡数据"`    // 关卡数据
	Star      int    `json:"star"       orm:"star"       description:"章节获得的总星"` // 章节获得的总星
	RwdData   string `json:"rwd_data"   orm:"rwd_data"   description:"通关奖励领取"`  // 通关奖励领取
}
