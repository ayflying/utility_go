package drive

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

var adapterRedisClient gcache.Adapter
var adapterRedisCache = gcache.New()

func NewAdapterRedis() gcache.Adapter {

	if adapterRedisClient == nil {
		_cfg, _ := g.Cfg().Get(gctx.New(), "redis.default")
		var cfg *gredis.Config
		_cfg.Scan(&cfg)
		redisObj, _ := gredis.New(cfg)
		//adapterRedisClient = gcache.NewAdapterRedis(g.Redis("default"))
		adapterRedisClient = gcache.NewAdapterRedis(redisObj)

		adapterRedisCache.SetAdapter(adapterRedisClient)
	}
	return adapterRedisCache
}
