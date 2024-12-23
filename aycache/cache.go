package aycache

import (
	"github.com/ayflying/utility_go/ayCache/drive"
	"github.com/gogf/gf/v2/os/gcache"
)

type Mod struct {
	client *gcache.Cache
}

//func NewV1(_name ...string) *cache.Mod {
//	return pgk.Cache
//}

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
	}

	//var client = gcache.New()
	//client.SetAdapter(cacheAdapterObj)
	return cacheAdapterObj
}
