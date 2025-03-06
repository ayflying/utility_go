package pgk

import (
	"github.com/gogf/gf/v2/os/gcache"
	v1 "new-gitlab.adesk.com/public_project/utility_go/api/pgk/v1"
	"new-gitlab.adesk.com/public_project/utility_go/pgk/aycache"
	"new-gitlab.adesk.com/public_project/utility_go/pgk/notice"
	"new-gitlab.adesk.com/public_project/utility_go/pgk/rank"
	"new-gitlab.adesk.com/public_project/utility_go/pgk/s3"
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
