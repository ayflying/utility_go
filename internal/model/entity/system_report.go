// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemReport is the golang structure for table system_report.
type SystemReport struct {
	Id        int         `json:"id"         orm:"id"         description:""`      //
	Rid       int         `json:"rid"        orm:"rid"        description:"举报id"`  // 举报id
	Uid       int         `json:"uid"        orm:"uid"        description:"举报人编号"` // 举报人编号
	Type      int         `json:"type"       orm:"type"       description:"举报类型"`  // 举报类型
	Desc      string      `json:"desc"       orm:"desc"       description:"举报正文"`  // 举报正文
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"举报时间"`  // 举报时间
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间"`  // 删除时间
	Status    int         `json:"status"     orm:"status"     description:"处理状态"`  // 处理状态
}
