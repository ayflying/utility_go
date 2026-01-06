package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/apolloconfig/agollo/v4"
	apolloConfig "github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/gogf/gf/contrib/config/apollo/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/contrib/config/consul/v2"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
)

var (
	//ApolloCfg      *apolloConfig.AppConfig
	ApolloCfg *apollo.Config
	// ApolloListener 存储需要监听的 Apollo 配置项名称
	ApolloListener []string
	// Item2Obj 存储配置项名称和对应的加载器对象的映射
	Item2Obj = map[string]Load{}

	ConsulIsWatcher = map[string]bool{}
)

// Load 接口定义了 Load 方法，用于加载数据
type Load interface {
	// Load 方法用于加载配置数据，支持传入可选的配置参数
	Load(cfg ...string)
}

// NewV1 创建一个新的 Cfg 实例
func NewV1() *Cfg {
	return &Cfg{}
}

// Cfg 结构体包含配置操作的相关方法
type Cfg struct {
	// Lock 用于保证并发安全的互斥锁
	Lock sync.Mutex
}

// GetDbFile 从数据库中获取指定名称的配置文件
// 参数:
//   name - 配置文件的名称
// 返回值:
//   *g.Var - 存储配置数据的变量
//   error  - 操作过程中遇到的错误
func (c *Cfg) GetDbFile(name string) (res *g.Var, err error) {
	// 从数据库的 game_config 表中查询指定名称的配置数据
	get2, err := g.Model("game_config").
		Where("name", name).Master().Value("data")
	// 将查询结果扫描到 res 变量中
	err = get2.Scan(&res)
	if res == nil {
		res = &gvar.Var{}
	}
	return
}

// GetFile 从文件系统或资源文件中加载 JSON 配置
// 参数:
//   filename   - 需要加载的配置文件名（不带扩展名）
//   _pathStr   - 可选参数，指定配置文件目录路径，默认"manifest/game/"
// 返回值:
//   *gjson.Json - 解析后的 JSON 对象
//   error       - 文件加载或解析过程中遇到的错误
func (c *Cfg) GetFile(filename string, _pathStr ...string) (jsonObj *gjson.Json, err error) {
	// 处理路径参数，使用默认路径或传入参数
	pathStr := "manifest/game/"
	if len(_pathStr) > 0 {
		pathStr = _pathStr[0]
	}
	// 拼接完整的文件路径
	filePath := pathStr + filename + ".json"

	// 载入静态资源到文件对象
	err = gres.Load(filePath)
	var bytes []byte

	// 优先从文件系统读取，不存在时从资源文件读取
	if gfile.IsFile(filePath) {
		bytes = gfile.GetBytes(filePath) // 读取物理文件内容
	} else {
		bytes = gres.GetContent(filePath) // 从打包资源中获取内容
	}

	for range 5 {
		//如果还是没有读取到配置，从当前目录返回上级读取
		if bytes == nil {
			// 上级拼接完整的文件路径
			filePath = "../" + filePath
			if gfile.IsFile(filePath) {
				bytes = gfile.GetBytes(filePath) // 读取物理文件内容
				//找到配置了，跳过
				break
			}
		}
	}
	if bytes == nil {
		g.Log().Errorf(gctx.New(), "未读取到配置文件：%v", filePath)
	}

	// 解析 JSON 内容并返回结果
	jsonObj, err = gjson.DecodeToJson(bytes)
	return
}

// GetUrlFile 获取远程配置
// 参数:
//   name - 配置文件的名称
// 返回值:
//   *gjson.Json - 解析后的 JSON 对象
//   error       - 请求或解析过程中遇到的错误
func (c *Cfg) GetUrlFile(name string) (jsonObj *gjson.Json, err error) {
	// 拼接远程配置文件的 URL
	urlStr := fmt.Sprintf("http://sdf.sdfs.sdf/%s.json", name)
	// 发送 HTTP 请求获取远程配置数据
	getUrl, err := g.Client().Get(nil, urlStr)
	// 读取响应内容
	bytes := getUrl.ReadAll()
	// 解析 JSON 内容并返回结果
	jsonObj, err = gjson.DecodeToJson(bytes)
	return
}

func (c *Cfg) GetConsul(name string, obj Load) (jsonObj *gjson.Json, err error) {
	ctx := gctx.New()
	// 加载在线配置
	getCfg := g.Cfg().MustGet(ctx, "cfg")
	if getCfg.IsNil() {
		return
	}
	g.Log().Info(ctx, " - 初始化consul在线配置")
	var cfg = getCfg.MapStrStr()
	var consulConfig = api.Config{
		Address:    cfg["address"],                     // Consul server address
		Scheme:     "http",                             // Connection scheme (http/https)
		Datacenter: "dc1",                              // Datacenter name
		Transport:  cleanhttp.DefaultPooledTransport(), // HTTP transport with connection pooling
		Token:      cfg["token"],                       // ACL token for authentication
	}
	var configPath = name + ".json"

	adapter, err := consul.New(ctx, consul.Config{
		ConsulConfig: consulConfig, // Consul client configuration
		Path:         configPath,   // Configuration path in KV store
		Watch:        true,         // Enable configuration watching for updates
	})
	// 更改默认配置实例的适配器
	g.Cfg(name).SetAdapter(adapter)

	getCfg, err = g.Cfg(name).Get(nil, ".")
	// 将配置值扫描到 jsonObj 中
	getCfg.Scan(&jsonObj)

	// 添加观察者到配置实例的适配器
	if ok, _ := ConsulIsWatcher[name]; !ok {
		if adapter2, ok2 := g.Cfg(name).GetAdapter().(gcfg.WatcherAdapter); ok2 {
			ConsulIsWatcher[name] = true
			adapter2.AddWatcher("config-watcher", func(ctx context.Context) {
				fmt.Println("配置发生了变化")
				//cfg3 := g.Cfg(name).MustGet(ctx, ".")
				obj.Load()
			})
		}
	}

	return
}

// GetApollo 从 Apollo 配置中心获取指定名称的配置
// 参数:
//   name - 配置文件的名称
//   obj  - 实现了 Load 接口的加载器对象
// 返回值:
//   *gjson.Json - 解析后的 JSON 对象
//   error       - 操作过程中遇到的错误
func (c *Cfg) GetApollo(name string, obj Load) (jsonObj *gjson.Json, err error) {
	// 将配置项名称和对应的加载器对象存入映射
	Item2Obj[name+".json"] = obj

	// 接入 Apollo 配置
	ApolloCfg.NamespaceName = name + ".json"
	// 创建 Apollo 配置适配器
	adapter, err := apollo.New(nil, *ApolloCfg)
	if err != nil {
		// 配置适配器创建失败，记录致命错误日志
		g.Log().Fatalf(nil, `%+v`, err)
	}
	// 更改默认配置实例的适配器
	g.Cfg(name).SetAdapter(adapter)

	// 首次运行加入监听器
	if !gstr.InArray(ApolloListener, name+".json") {
		// 启动 Apollo 客户端
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
		// 创建自定义监听器实例
		c2 := &CustomChangeListener{}
		// 为 Apollo 客户端添加监听器
		client.AddChangeListener(c2)
		// 将配置项名称添加到监听器列表
		ApolloListener = append(ApolloListener, name+".json")
	}

	// 从配置中心获取指定配置项的值
	cfg, err := g.Cfg(name).Get(nil, "content")
	// 将配置值扫描到 jsonObj 中
	cfg.Scan(&jsonObj)
	return
}

// CustomChangeListener 是 Apollo 配置变化的自定义监听器
type CustomChangeListener struct {
	// wg 用于等待所有处理任务完成
	wg sync.WaitGroup
}

// OnChange 当 Apollo 配置发生变化时触发
func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	// 记录配置变化的日志
	g.Log().Debugf(nil, "当前Namespace变化了：%v", changeEvent.Namespace)
	// 获取变化的配置项名称
	filename := changeEvent.Namespace
	if obj, ok := Item2Obj[filename]; ok {
		// 重载配置文件
		obj.Load(changeEvent.Changes["content"].NewValue.(string))
	}
}

// OnNewestChange 当获取到最新配置时触发，当前为空实现
func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here
}
