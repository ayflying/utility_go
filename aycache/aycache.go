package aycache

import (
	"github.com/ayflying/utility_go/pkg/aycache/drive"
	"github.com/gogf/gf/v2/os/gcache"
)

type Mod struct {
	client *gcache.Cache
}

// Deprecated:弃用，改用 pkg.Cache()
func New(_name ...string) gcache.Adapter {

	var cacheAdapterObj gcache.Adapter
	var name = "cache"
	if len(_name) > 0 {
		name = _name[0]
	}
	switch name {
	case "cache":
		cacheAdapterObj = drive.NewAdapterMemory()
	case "redis":
		cacheAdapterObj = drive.NewAdapterRedis()
	case "file":
		cacheAdapterObj = drive.NewAdapterFile("runtime/cache")
		//case "es":
		//cacheAdapterObj = drive.NewAdapterElasticsearch("http://127.0.0.1:9200"})
	}

	//var client = gcache.New()
	//client.SetAdapter(cacheAdapterObj)
	return cacheAdapterObj
}
