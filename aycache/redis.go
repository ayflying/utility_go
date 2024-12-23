package aycache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var adapterRedisClient gcache.Adapter
var adapterRedisCache = gcache.New()

func NewAdapterRedis() gcache.Adapter {

	if adapterRedisClient == nil {
		adapterRedisClient = gcache.NewAdapterRedis(g.Redis("cache"))
		adapterRedisCache.SetAdapter(adapterRedisClient)
	}
	return adapterRedisCache
}
