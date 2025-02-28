// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameMailMass is the golang structure of table shiningu_game_mail_mass for DAO operations like Where/Data.
type GameMailMass struct {
	g.Meta    `orm:"table:shiningu_game_mail_mass, do:true"`
	Id        interface{} // 主键
	Title     interface{} // 标题
	Type      interface{} // 类型
	Content   interface{} // 正文
	CreatedAt *gtime.Time // 创建时间
	Items     interface{} // 奖励
	Sign      interface{} // 署名
	EndTime   *gtime.Time // 结束时间
}
