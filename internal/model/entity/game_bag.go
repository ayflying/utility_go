// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GameBag is the golang structure for table game_bag.
type GameBag struct {
	Uid  int64  `json:"uid"  orm:"uid"  description:"用户标识"` // 用户标识
	List string `json:"list" orm:"list" description:"道具数据"` // 道具数据
	Book string `json:"book" orm:"book" description:"图鉴"`   // 图鉴
	Hand string `json:"hand" orm:"hand" description:"手势"`   // 手势
}
