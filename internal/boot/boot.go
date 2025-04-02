package boot

import (
	"context"
	v1 "github.com/ayflying/utility_go/api/system/v1"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	ctx = gctx.GetInitCtx()
)

func Boot() (err error) {
	err = service.SystemCron().StartCron()

	//用户活动持久化
	gtimer.SetTimeout(ctx, time.Minute, func(ctx context.Context) {
		service.SystemCron().AddCron(v1.CronType_DAILY, func() error {
			service.GameAct().Saves()
			return nil
		})
	})

	return nil
}
