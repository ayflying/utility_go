package drive

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

var adapterRedisClient = make(map[string]gcache.Adapter)
var adapterRedisCache = make(map[string]*gcache.Cache)

func NewAdapterRedis(name string) gcache.Adapter {
	if adapterRedisClient[name] == nil {
		_cfg, err := g.Cfg().Get(gctx.New(), "redis."+name)
		if err != nil {
			panic("当前redis配置不存在")
		}
		var cfg *gredis.Config
		_cfg.Scan(&cfg)
		redisObj, _ := gredis.New(cfg)
		//adapterRedisClient[name] = gcache.NewAdapterRedis(g.Redis(name))
		adapterRedisClient[name] = gcache.NewAdapterRedis(redisObj)
		adapterRedisCache[name] = gcache.New()
		adapterRedisCache[name].SetAdapter(adapterRedisClient[name])
	}
	return adapterRedisCache[name]
}
