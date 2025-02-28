package utility_go

import (
	"context"
	v1 "github.com/ayflying/utility_go/api/system/v1"
	_ "github.com/ayflying/utility_go/internal/logic"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"

	"github.com/ayflying/utility_go/config"
)

var (
	Config = config.Cfg{}
	ctx    = gctx.GetInitCtx()
)

func init() {
	service.SystemCron().StartCron()

	//用户活动持久化
	gtimer.SetTimeout(ctx, time.Minute, func(ctx context.Context) {
		service.SystemCron().AddCron(v1.CronType_DAILY, func() error {
			service.GameAct().Saves()
			return nil
		})
	})

}
