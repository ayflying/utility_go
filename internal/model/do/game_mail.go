// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameMail is the golang structure of table shiningu_game_mail for DAO operations like Where/Data.
type GameMail struct {
	g.Meta    `orm:"table:shiningu_game_mail, do:true"`
	Id        interface{} // 流水
	Uid       interface{} // 用户标识
	Type      interface{} // 类型
	Title     interface{} // 标题
	Content   interface{} // 正文
	Items     interface{} // 奖励道具
	HaveItems interface{} // 是否有附件
	CreatedAt *gtime.Time // 创建时间
	Sign      interface{} // 署名
	EndTime   *gtime.Time // 结束时间
	Extend    interface{} // 附加参数
	Status    interface{} // 状态
}
