// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberSave is the golang structure of table shiningu_member_save for DAO operations like Where/Data.
type MemberSave struct {
	g.Meta    `orm:"table:shiningu_member_save, do:true"`
	Uid       interface{} // 用户编号
	Type      interface{} // 存档类型
	Slot      interface{} // 存档槽位
	Data      interface{} // 存档内容
	S3        interface{} // s3地址
	UpdatedAt *gtime.Time // 更新时间
	Name      interface{} // 自定义名字
	Image     interface{} // 上传图片
	UseIds    interface{} // 使用的道具id
}
