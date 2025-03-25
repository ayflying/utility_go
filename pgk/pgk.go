package pgk

import (
	v1 "github.com/ayflying/utility_go/api/pkg/v1"
	"github.com/ayflying/utility_go/pkg/aycache"
	"github.com/ayflying/utility_go/pkg/notice"
	"github.com/ayflying/utility_go/pkg/rank"
	"github.com/ayflying/utility_go/pkg/s3"
	"github.com/gogf/gf/v2/os/gcache"
)

var ()

// 统一调用
// Deprecated: 请使用 pkg.Notice() 方法替代。
func Notice(typ v1.NoticeType, host string) notice.MessageV1 {
	return notice.New(typ, host)
}

// 统一调用cache
// Deprecated: 请使用 pkg.Cache() 方法替代。
func Cache(_name ...string) gcache.Adapter {
	return aycache.New(_name...)
}

// Deprecated: 请使用 pkg.S3() 方法替代。
func S3(_name ...string) *s3.Mod {
	return s3.New(_name...)
}

// Deprecated: 请使用 pkg.Rank() 方法替代。
func Rank() *rank.Mod {
	return rank.New()
}
