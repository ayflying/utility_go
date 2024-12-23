package drive

import (
	"github.com/gogf/gf/v2/os/gcache"
)

var adapterMemoryClient = gcache.New()

// NewAdapterMemory 创建并返回一个新的内存缓存对象。
func NewAdapterMemory() gcache.Adapter {
	//if adapterMemoryClient == nil {
	//	adapterMemoryClient = gcache.New()
	//}
	return adapterMemoryClient
}
