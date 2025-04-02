package cmd

import (
	"github.com/ayflying/utility_go/controller/callback"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Load(s *ghttp.Server) {

	//开启prometheus监控
	s.Group("/metrics", func(group *ghttp.RouterGroup) {
		group.Bind(
			ghttp.WrapH(promhttp.Handler()),
		)
	})
}

// 注册游客方法
func RegistrationAnonymous(group *ghttp.RouterGroup) (res []interface{}) {
	group.Bind(
		callback.NewV1(),
	)
	return

}

// 注册用户方法
func RegistrationUser(group *ghttp.RouterGroup) (res []interface{}) {
	group.Bind()
	return
}
