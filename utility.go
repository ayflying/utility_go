package utility_go

import (
	v1 "github.com/ayflying/utility_go/api/system/v1"
	_ "github.com/ayflying/utility_go/internal/logic"
	"github.com/ayflying/utility_go/service2"

	"github.com/ayflying/utility_go/config"
)

var (
	Config = config.Cfg{}
)

func init() {
	service2.SystemCron().StartCron()

	//用户活动持久化
	service2.SystemCron().AddCron(v1.CronType_DAILY, func() error {
		service2.GameAct().Saves()
		return nil
	})
}
