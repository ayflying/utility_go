package config

import (
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/ayflying/utility_go/pkg"
	"github.com/gogf/gf/contrib/config/apollo/v2"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

var (
	//ApolloCfg      *apolloConfig.AppConfig
	ApolloCfg      *apollo.Config
	ApolloListener []string
	Item2Obj       = map[string]Load{}
)

// load接口定义了Load方法，用于加载数据
type Load interface {
	Load(cfg ...string)
}

type Cfg struct {
	Lock sync.Mutex
}

// Deprecated : pkg.Config().GetDbFile(name)
func (c *Cfg) GetDbFile(name string) (res *g.Var, err error) {
	pkg.Config().GetDbFile(name)
	return
}

// Deprecated : pkg.Config().GetFile(name, obj...)
func (c *Cfg) GetFile(filename string, obj ...Load) (jsonObj *gjson.Json, err error) {
	pkg.Config().GetFile(filename)
	return
}

// getUrlFile 获取远程配置
// Deprecated : pkg.Config().GetUrlFile(name)
func (c *Cfg) GetUrlFile(name string) (jsonObj *gjson.Json, err error) {
	pkg.Config().GetUrlFile(name)
	return
}

// Deprecated : pkg.Config().GetApollo(name, obj)
func (c *Cfg) GetApollo(name string, obj Load) (jsonObj *gjson.Json, err error) {
	pkg.Config().GetApollo(name, obj)
	return
}

// 阿波罗监听器
type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	g.Log().Debugf(nil, "当前Namespace变化了：%v", changeEvent.Namespace)
	filename := changeEvent.Namespace
	if obj, ok := Item2Obj[filename]; ok {
		//重载配置文件
		obj.Load(changeEvent.Changes["content"].NewValue.(string))
	}
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here

}
