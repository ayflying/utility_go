package aycache

import (
	"context"
	"math"

	v1 "github.com/ayflying/utility_go/api/system/v1"
	"github.com/ayflying/utility_go/internal/boot"
	"github.com/ayflying/utility_go/pkg/aycache/drive"
	drive2 "github.com/ayflying/utility_go/pkg/aycache/drive"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Mod 定义缓存模块结构体，包含一个 gcache.Cache 客户端实例
type Mod struct {
	client *gcache.Cache
}

// QPSCount 记录缓存的 QPS 计数
var QPSCount int

// QPS 是一个 Prometheus 指标，用于记录当前缓存的 QPS 数量
var QPS = promauto.NewGauge(
	prometheus.GaugeOpts{
		Name: "Cache_QPS",
		Help: "当前缓存QPS数量",
	},
)

// init 函数在包被导入时执行，用于初始化定时任务以更新 QPS 指标
func init() {
	boot.AddFunc(func() {
		// 初始化指标，每分钟计算一次平均 QPS 并重置计数器
		service.SystemCron().AddCronV2(v1.CronType_MINUTE, func(context.Context) error {
			QPS.Set(math.Round(float64(QPSCount) / 60))
			QPSCount = 0
			return nil
		})
	})
}

// New 根据传入的名称创建不同类型的缓存适配器
// 如果未传入名称，默认使用 "cache" 类型
// 支持的类型包括 "cache"（内存缓存）、"redis"（Redis 缓存）、"file"（文件缓存）和 "es"（Elasticsearch 缓存）
func New(_name ...string) gcache.Adapter {
	var cacheAdapterObj gcache.Adapter
	var name = "cache"
	if len(_name) > 0 {
		name = _name[0]
	}
	switch name {
	case "cache":
		// 创建内存缓存适配器
		cacheAdapterObj = drive2.NewAdapterMemory()
	case "redis":
		//第二个参数为配置名称，默认为default
		var typ = "default"
		if len(_name) >= 2 {
			typ = _name[1]
		}
		// 创建 Redis 缓存适配器
		cacheAdapterObj = drive2.NewAdapterRedis(typ)
	case "file":
		// 创建文件缓存适配器，指定缓存目录为 "runtime/cache"
		cacheAdapterObj = drive2.NewAdapterFile("runtime/cache")
	case "es":
		// 创建 Elasticsearch 缓存适配器，需要传入额外参数
		cacheAdapterObj = drive.NewAdapterElasticsearch(_name[1])
	}

	//var client = gcache.New()
	//client.SetAdapter(cacheAdapterObj)

	// 每次创建适配器时，QPS 计数加 1
	QPSCount++
	return cacheAdapterObj
}
