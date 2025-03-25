package pkg

import (
	v1 "github.com/ayflying/utility_go/api/pgk/v1"
	"github.com/ayflying/utility_go/pkg/aycache"
	"github.com/ayflying/utility_go/pkg/config"
	"github.com/ayflying/utility_go/pkg/notice"
	"github.com/ayflying/utility_go/pkg/rank"
	"github.com/ayflying/utility_go/pkg/s3"
	"github.com/ayflying/utility_go/pkg/websocket"
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

func Websocket() *websocket.SocketV1 {
	return websocket.NewV1()
}

func Config() *config.Cfg {
	return config.NewV1()
}
