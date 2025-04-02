package aycache

import (
	"github.com/ayflying/utility_go/pkg"
	"github.com/gogf/gf/v2/os/gcache"
)

type Mod struct {
	client *gcache.Cache
}

// Deprecated:弃用，改用 pkg.Cache()
func New(_name ...string) gcache.Adapter {
	return pkg.Cache(_name...)
}
