// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

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
		// Deprecated: Use message.New(message.DingTalk, DingTalkWebHook).Send(value)
		DingTalk(DingTalkWebHook string, value string) (res *gclient.Response)
		ReadLog()
		// AddCron 添加一个定时任务到相应的调度列表中。
		//
		// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
		// 确保自定义任务正确处理上下文取消信号，即可充分发挥超时打断功能。
		// @receiver s: sSystemCron的实例，代表一个调度系统。
		// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
		// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
		// deprecated: 弃用，请使用 AddCronV2
		AddCron(typ v1.CronType, _func func() error)
		// AddCronV2  添加一个定时任务到相应的调度列表中。
		//
		// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
		// @receiver s: sSystemCron的实例，代表一个调度系统。
		// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
		// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
		// @param unique: 是否只在唯一服务器上执行
		AddCronV2(typ v1.CronType, _func func(context.Context) error, unique ...bool)
		// StartCron 开始计划任务执行
		//
		//	@Description:
		//	@receiver s
		//	@return err
		StartCron() (err error)
		// AddFuncChan 添加方法到通道
		AddFuncChan(list []func(context.Context) error)
		// RunFuncChan 统一执行方法
		RunFuncChan()
		// RunFunc 统一执行方法
		// deprecated: 弃用，会造成周期任务并发执行，to service.SystemCron().AddFuncChan
		RunFunc(list []func() error)
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
