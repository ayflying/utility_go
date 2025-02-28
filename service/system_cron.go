// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "github.com/ayflying/utility_go/api/system/v1"
	"github.com/gogf/gf/v2/net/gclient"
)

type (
	ISystemCron interface {
		Guardian(DingTalkWebHook string)
		// DingTalk 发送钉钉消息
		//
		// @Description: 向指定的钉钉机器人发送消息。
		// @receiver s: 系统定时任务结构体指针。
		// @param value: 要发送的消息内容。
		DingTalk(DingTalkWebHook string, value string) (res *gclient.Response)
		ReadLog()
		// AddCron 添加一个定时任务到相应的调度列表中。
		//
		// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
		// @receiver s: sSystemCron的实例，代表一个调度系统。
		// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
		// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
		AddCron(typ v1.CronType, _func func() error)
		// StartCron 开始计划任务执行
		//
		//	@Description:
		//	@receiver s
		//	@return err
		StartCron() (err error)
	}
)

var (
	localSystemCron ISystemCron
)

func SystemCron() ISystemCron {
	if localSystemCron == nil {
		panic("implement not found for interface ISystemCron, forgot register?")
	}
	return localSystemCron
}

func RegisterSystemCron(i ISystemCron) {
	localSystemCron = i
}
