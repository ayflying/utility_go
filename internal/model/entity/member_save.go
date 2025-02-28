// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberSave is the golang structure for table member_save.
type MemberSave struct {
	Uid       int64       `json:"uid"        orm:"uid"        description:"用户编号"`    // 用户编号
	Type      int         `json:"type"       orm:"type"       description:"存档类型"`    // 存档类型
	Slot      int         `json:"slot"       orm:"slot"       description:"存档槽位"`    // 存档槽位
	Data      string      `json:"data"       orm:"data"       description:"存档内容"`    // 存档内容
	S3        string      `json:"s_3"        orm:"s3"         description:"s3地址"`    // s3地址
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`    // 更新时间
	Name      string      `json:"name"       orm:"name"       description:"自定义名字"`   // 自定义名字
	Image     string      `json:"image"      orm:"image"      description:"上传图片"`    // 上传图片
	UseIds    string      `json:"use_ids"    orm:"use_ids"    description:"使用的道具id"` // 使用的道具id
}
