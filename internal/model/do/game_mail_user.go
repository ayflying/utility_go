// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameMailUser is the golang structure of table shiningu_game_mail_user for DAO operations like Where/Data.
type GameMailUser struct {
	g.Meta `orm:"table:shiningu_game_mail_user, do:true"`
	Uid    interface{} //
	Mass   interface{} // 群发邮件领取列表
}
