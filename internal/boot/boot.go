package boot

import (
	"context"

	v1 "github.com/ayflying/utility_go/api/system/v1"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx   = gctx.GetInitCtx()
	_func = []func(){}
)

func Boot() (err error) {
	// 启动计划任务定时器，预防debug工具激活计划任务造成重复执行，此处不执行计划任务
	//err = service.SystemCron().StartCron()

	//用户活动持久化每小时执行一次
	service.SystemCron().AddCronV2(v1.CronType_HOUR, func(ctx context.Context) error {
		err = service.GameKv().SavesV1()
		err = service.GameAct().Saves(ctx)
		return err
	}, true)

	//初始化自启动方法
	for _, v := range _func {
		v()
	}
	return nil
}

// AddFunc 注册方法，在启动时执行
func AddFunc(f func()) {
	_func = append(_func, f)
}
