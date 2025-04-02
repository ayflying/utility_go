package utility_go

import (
	"github.com/ayflying/utility_go/internal/boot"
	_ "github.com/ayflying/utility_go/internal/logic"

	"github.com/ayflying/utility_go/config"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Config = config.Cfg{}
	ctx    = gctx.GetInitCtx()
)

func init() {
	var err = boot.Boot()
	if err != nil {
		panic(err)
	}
}
