package utility_go

import (
	"github.com/ayflying/utility_go/config"
	"github.com/ayflying/utility_go/internal/boot"
	_ "github.com/ayflying/utility_go/internal/logic"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Config = config.Cfg{}
	ctx    = gctx.GetInitCtx()
)

func init() {
	g.Log().Debug(ctx, "utility_go init启动完成")
	// 初始化配置
	var err = boot.Boot()

	if err != nil {
		panic(err)
	}

}
