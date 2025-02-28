// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameMail is the golang structure for table game_mail.
type GameMail struct {
	Id        int64       `json:"id"         orm:"id"         description:"流水"`    // 流水
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户标识"`  // 用户标识
	Type      int         `json:"type"       orm:"type"       description:"类型"`    // 类型
	Title     string      `json:"title"      orm:"title"      description:"标题"`    // 标题
	Content   string      `json:"content"    orm:"content"    description:"正文"`    // 正文
	Items     string      `json:"items"      orm:"items"      description:"奖励道具"`  // 奖励道具
	HaveItems bool        `json:"have_items" orm:"have_items" description:"是否有附件"` // 是否有附件
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`  // 创建时间
	Sign      string      `json:"sign"       orm:"sign"       description:"署名"`    // 署名
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"结束时间"`  // 结束时间
	Extend    string      `json:"extend"     orm:"extend"     description:"附加参数"`  // 附加参数
	Status    int         `json:"status"     orm:"status"     description:"状态"`    // 状态
}
