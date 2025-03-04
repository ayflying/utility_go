package pgk

import (
	v1 "github.com/ayflying/utility_go/api/pgk/v1"
	"github.com/ayflying/utility_go/pgk/aycache"
	"github.com/ayflying/utility_go/pgk/notice"
	"github.com/ayflying/utility_go/pgk/rank"
	"github.com/ayflying/utility_go/pgk/s3"
	"github.com/gogf/gf/v2/os/gcache"
)

var ()

// 统一调用
func Notice(typ v1.NoticeType, host string) notice.MessageV1 {
	return notice.New(typ, host)
}

// 统一调用cache
func Cache(_name ...string) gcache.Adapter {
	return aycache.New(_name...)
}

func S3(_name ...string) *s3.Mod {
	return s3.New(_name...)
}

func Rank() *rank.Mod {
	return rank.New()
}
