// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameBag is the golang structure of table shiningu_game_bag for DAO operations like Where/Data.
type GameBag struct {
	g.Meta `orm:"table:shiningu_game_bag, do:true"`
	Uid    interface{} // 用户标识
	List   interface{} // 道具数据
	Book   interface{} // 图鉴
	Hand   interface{} // 手势
}
