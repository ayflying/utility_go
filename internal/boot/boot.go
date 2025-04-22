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
	err = service.SystemCron().StartCron()

	//用户活动持久化
	service.SystemCron().AddCronV2(v1.CronType_DAILY, func(ctx context.Context) error {
		return service.GameAct().Saves(ctx)
	})

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
