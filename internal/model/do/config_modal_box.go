// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigModalBox is the golang structure of table shiningu_config_modal_box for DAO operations like Where/Data.
type ConfigModalBox struct {
	g.Meta      `orm:"table:shiningu_config_modal_box, do:true"`
	Id          interface{} // 主键
	ModalBoxId  interface{} // 弹框id
	UserType    interface{} // 特定用户
	Tips        interface{} // 弹框tips选项
	Name        interface{} // 名称
	Title       interface{} // 标题
	Content     interface{} // 正文
	Type        interface{} // 类型
	Style       interface{} // 样式
	Weight      interface{} // 权重
	Attachments interface{} // 附件
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Status      interface{} // 状态 1开始 0关闭
	Notes       interface{} // 备注
	StartTime   *gtime.Time // 开始时间
	EndTime     *gtime.Time // 结束时间
}
