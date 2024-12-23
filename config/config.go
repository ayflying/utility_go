package config

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
	"sync"
)

var (
	ApolloCfg      *config.AppConfig
	ApolloListener []string
	Item2Obj       = map[string]Load{}
)

// load接口定义了Load方法，用于加载数据
type Load interface {
	Load()
}

type Cfg struct {
}

func (c *Cfg) GetDbFile(name string) (res *g.Var, err error) {
	get2, err := g.Model("game_config").
		Where("name", name).Master().Value("data")
	err = get2.Scan(&res)
	if res == nil {
		res = &gvar.Var{}
	}
	return
}

func (c *Cfg) GetFile(filename string, obj ...Load) (jsonObj *gjson.Json, err error) {
	pathStr := "manifest/game/"
	filePath := pathStr + filename + ".json"
	//err := gres.Load(pathStr + filename)

	//载入静态资源到文件对象
	err = gres.Load(filePath)
	var bytes []byte

	if gfile.IsFile(filePath) {
		bytes = gfile.GetBytes(filePath)
	} else {
		bytes = gres.GetContent(filePath)
	}

	jsonObj, err = gjson.DecodeToJson(bytes)
	//g.Dump(filePath, jsonObj)
	return
}

// getUrlFile 获取远程配置
func (c *Cfg) GetUrlFile(name string) (jsonObj *gjson.Json, err error) {
	urlStr := fmt.Sprintf("http://sdf.sdfs.sdf/%s.json", name)
	getUrl, err := g.Client().Discovery(nil).Get(nil, urlStr)
	bytes := getUrl.ReadAll()
	jsonObj, err = gjson.DecodeToJson(bytes)
	return
}

// 获取阿波罗
func (c *Cfg) GetApollo(name string, obj Load) (jsonObj *gjson.Json, err error) {
	Item2Obj[name] = obj

	//c2 := &config.AppConfig{
	//	AppID:          consts.ApolloAppID,
	//	Cluster:        consts.ApolloCluster,
	//	IP:             consts.ApolloIP,
	//	NamespaceName:  consts.ApolloNamespaceName,
	//	IsBackupConfig: consts.ApolloIsBackupConfig,
	//	Secret:         consts.ApolloSecret,
	//}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return ApolloCfg, nil
	})

	get := client.GetConfigCache(ApolloCfg.NamespaceName)
	getStr, err := get.Get("content")
	jsonObj, err = gjson.DecodeToJson(getStr)

	//首次运行加入监听器
	if !gstr.InArray(ApolloListener, name) {
		c2 := &CustomChangeListener{}
		client.AddChangeListener(c2)
		ApolloListener = append(ApolloListener, name)
	}

	return
}

// 阿波罗监听器
type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
	fmt.Println(changeEvent.Changes)
	//for key, value := range changeEvent.Changes {
	//	fmt.Println("change key : ", key, ", value :", value)
	//}
	//fmt.Println(changeEvent.Namespace)
	//c.wg.Done()

	g.Log().Debugf(nil, "当前Namespace变化了：%v", changeEvent.Namespace)
	filename := changeEvent.Namespace
	if obj, ok := Item2Obj[filename]; ok {
		obj.Load()
	}

	//cfgList := consts.ExcelFileList
	//for _, v := range cfgList {
	//	if v.Name == changeEvent.Namespace {
	//
	//	}
	//}

}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here

}
