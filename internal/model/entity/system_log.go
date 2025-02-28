// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLog is the golang structure for table system_log.
type SystemLog struct {
	Id        int         `json:"id"         orm:"id"         description:"主键"`       // 主键
	Uid       int         `json:"uid"        orm:"uid"        description:"操作的用户"`    // 操作的用户
	Url       string      `json:"url"        orm:"url"        description:"当前访问的url"` // 当前访问的url
	Data      string      `json:"data"       orm:"data"       description:"操作数据"`     // 操作数据
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`     // 创建时间
	Ip        string      `json:"ip"         orm:"ip"         description:"当前ip地址"`   // 当前ip地址
}
