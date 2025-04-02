package aycache

import (
	"github.com/ayflying/utility_go/pkg/aycache/drive"
	drive2 "github.com/ayflying/utility_go/pkg/aycache/drive"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Mod struct {
	client *gcache.Cache
}

var (
	QPSCount int
	QPS      = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "Cache_QPS",
			Help: "当前缓存QPS数量",
		},
	)
)

func New(_name ...string) gcache.Adapter {

	var cacheAdapterObj gcache.Adapter
	var name = "cache"
	if len(_name) > 0 {
		name = _name[0]
	}
	switch name {
	case "cache":
		cacheAdapterObj = drive2.NewAdapterMemory()
	case "redis":
		cacheAdapterObj = drive2.NewAdapterRedis()
	case "file":
		cacheAdapterObj = drive2.NewAdapterFile("runtime/cache")
	case "es":
		cacheAdapterObj = drive.NewAdapterElasticsearch(_name[1])
	}

	//var client = gcache.New()
	//client.SetAdapter(cacheAdapterObj)

	QPSCount++
	return cacheAdapterObj
}
