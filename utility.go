package utility_go

import (
	"context"
	"time"

	"github.com/ayflying/utility_go/config"
	"github.com/ayflying/utility_go/internal/boot"
	_ "github.com/ayflying/utility_go/internal/logic"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
)

var (
	Config = config.Cfg{}
	ctx    = gctx.GetInitCtx()
)

func init() {
	var err error
	g.Log().Debug(ctx, "utility_go init启动完成")
	// 初始化配置
	gtimer.SetTimeout(ctx, time.Second*5, func(ctx context.Context) {
		err = boot.Boot()
	})

	if err != nil {
		panic(err)
	}

}
