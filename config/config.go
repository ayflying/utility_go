package config

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	apolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/contrib/config/apollo/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
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
//func (c *Cfg) GetApollo(name string, obj Load) (jsonObj *gjson.Json, err error) {
//	jsonObj, err = c.GetApolloV2(name, obj)
//	return
//
//	//c.Lock.Lock()
//	//defer c.Lock.Unlock()
//	//
//	//Item2Obj[name+".json"] = obj
//	//var cfg = apolloConfig.AppConfig{
//	//	AppID:             ApolloCfg.AppID,
//	//	Cluster:           ApolloCfg.Cluster,
//	//	IP:                ApolloCfg.IP,
//	//	NamespaceName:     name + ".json",
//	//	Secret:            ApolloCfg.Secret,
//	//	IsBackupConfig:    ApolloCfg.IsBackupConfig,
//	//	BackupConfigPath:  ApolloCfg.BackupConfigPath,
//	//	SyncServerTimeout: 60,
//	//	MustStart:         true,
//	//}
//	////cfg.NamespaceName = name + ".json"
//	//
//	//client, err := agollo.StartWithConfig(func() (*apolloConfig.AppConfig, error) {
//	//	return ApolloCfg, nil
//	//})
//	//if client == nil {
//	//	return
//	//}
//	//var getStr string
//	//var getApollo *storage.Config
//	//for range 5 {
//	//	getApollo = client.GetConfig(cfg.NamespaceName)
//	//	if getApollo != nil {
//	//		break
//	//	}
//	//	time.Sleep(time.Second * 5)
//	//}
//	//
//	//if getApollo != nil {
//	//	getStr = getApollo.GetValue("content")
//	//	if getStr != "" {
//	//		//写入配置
//	//		gfile.PutContents(path.Join("manifest", "game", name+".json"), getStr)
//	//	}
//	//} else {
//	//	jsonObj, err = c.GetFile(name)
//	//}
//	//jsonObj, err = gjson.DecodeToJson(getStr)
//	////首次运行加入监听器
//	//if !gstr.InArray(ApolloListener, name) {
//	//	c2 := &CustomChangeListener{}
//	//	client.AddChangeListener(c2)
//	//	ApolloListener = append(ApolloListener, name)
//	//}
//	//return
//}

func (c *Cfg) GetApollo(name string, obj Load) (jsonObj *gjson.Json, err error) {
	Item2Obj[name+".json"] = obj

	// 接入阿波罗配置
	ApolloCfg.NamespaceName = name + ".json"
	adapter, err := apollo.New(nil, *ApolloCfg)
	if err != nil {
		g.Log().Fatalf(nil, `%+v`, err)
	}
	// Change the adapter of default configuration instance.
	g.Cfg(name).SetAdapter(adapter)

	//首次运行加入监听器
	if !gstr.InArray(ApolloListener, name+".json") {
		//放置监听器
		client, _ := agollo.StartWithConfig(func() (*apolloConfig.AppConfig, error) {
			return &apolloConfig.AppConfig{
				AppID:             ApolloCfg.AppID,
				Cluster:           ApolloCfg.Cluster,
				NamespaceName:     ApolloCfg.NamespaceName,
				IP:                ApolloCfg.IP,
				IsBackupConfig:    ApolloCfg.IsBackupConfig,
				BackupConfigPath:  ApolloCfg.BackupConfigPath,
				Secret:            ApolloCfg.Secret,
				SyncServerTimeout: ApolloCfg.SyncServerTimeout,
				MustStart:         ApolloCfg.MustStart,
			}, nil
		})
		c2 := &CustomChangeListener{}
		client.AddChangeListener(c2)
		ApolloListener = append(ApolloListener, name+".json")
	}

	cfg, err := g.Cfg(name).Get(nil, "content")
	cfg.Scan(&jsonObj)
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
