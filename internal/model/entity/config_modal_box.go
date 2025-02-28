// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigModalBox is the golang structure for table config_modal_box.
type ConfigModalBox struct {
	Id          int         `json:"id"           orm:"id"           description:"主键"`         // 主键
	ModalBoxId  int         `json:"modal_box_id" orm:"modal_box_id" description:"弹框id"`       // 弹框id
	UserType    string      `json:"user_type"    orm:"user_type"    description:"特定用户"`       // 特定用户
	Tips        string      `json:"tips"         orm:"tips"         description:"弹框tips选项"`   // 弹框tips选项
	Name        string      `json:"name"         orm:"name"         description:"名称"`         // 名称
	Title       string      `json:"title"        orm:"title"        description:"标题"`         // 标题
	Content     string      `json:"content"      orm:"content"      description:"正文"`         // 正文
	Type        int         `json:"type"         orm:"type"         description:"类型"`         // 类型
	Style       string      `json:"style"        orm:"style"        description:"样式"`         // 样式
	Weight      int         `json:"weight"       orm:"weight"       description:"权重"`         // 权重
	Attachments string      `json:"attachments"  orm:"attachments"  description:"附件"`         // 附件
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:"创建时间"`       // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:"更新时间"`       // 更新时间
	Status      int         `json:"status"       orm:"status"       description:"状态 1开始 0关闭"` // 状态 1开始 0关闭
	Notes       string      `json:"notes"        orm:"notes"        description:"备注"`         // 备注
	StartTime   *gtime.Time `json:"start_time"   orm:"start_time"   description:"开始时间"`       // 开始时间
	EndTime     *gtime.Time `json:"end_time"     orm:"end_time"     description:"结束时间"`       // 结束时间
}
