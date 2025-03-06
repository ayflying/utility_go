// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "github.com/ayflying/utility_go/api/admin/v1"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	ISystemLog interface {
		List(page int) (list []*v1.SystemLog, max int, err error)
		// 写入操作日志
		AddLog(uid int, url string, ip string, data g.Map) (id int64, err error)
	}
)

var (
	localSystemLog ISystemLog
)

func SystemLog() ISystemLog {
	if localSystemLog == nil {
		panic("implement not found for interface ISystemLog, forgot register?")
	}
	return localSystemLog
}

func RegisterSystemLog(i ISystemLog) {
	localSystemLog = i
}
