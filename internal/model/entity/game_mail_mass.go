// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameMailMass is the golang structure for table game_mail_mass.
type GameMailMass struct {
	Id        int         `json:"id"         orm:"id"         description:"主键"`   // 主键
	Title     string      `json:"title"      orm:"title"      description:"标题"`   // 标题
	Type      int         `json:"type"       orm:"type"       description:"类型"`   // 类型
	Content   string      `json:"content"    orm:"content"    description:"正文"`   // 正文
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	Items     string      `json:"items"      orm:"items"      description:"奖励"`   // 奖励
	Sign      string      `json:"sign"       orm:"sign"       description:"署名"`   // 署名
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"结束时间"` // 结束时间
}
