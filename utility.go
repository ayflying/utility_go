package utility_go

import (
	v1 "github.com/ayflying/utility_go/api/system/v1"
	_ "github.com/ayflying/utility_go/internal/logic"
	"github.com/ayflying/utility_go/service"

	"github.com/ayflying/utility_go/config"
)

var (
	Config = config.Cfg{}
)

func init() {
	service.SystemCron().StartCron()

	//用户活动持久化
	service.SystemCron().AddCron(v1.CronType_DAILY, func() error {
		service.GameAct().Saves()
		return nil
	})
}
