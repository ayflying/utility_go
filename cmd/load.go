package cmd

import (
	"github.com/ayflying/utility_go/controller/callback"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 注册游客方法
func RegistrationAnonymous(group *ghttp.RouterGroup) (res []interface{}) {
	group.Bind(
		callback.NewV1(),
	)
	return

}

// 注册用户方法
func RegistrationUser(group *ghttp.RouterGroup) (res []interface{}) {
	group.Bind(
		callback.NewV1(),
	)
	return
}
