package systemLog

import (
	"context"

	v1 "github.com/ayflying/utility_go/api/admin/v1"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sSystemLog struct {
	ctx context.Context
}

var (
	ctx = gctx.New()
	//db   = dao.AdminSystemLog.Ctx(ctx)
	name = "system_log"
)

func init() {
	service.RegisterSystemLog(New())
}

func New() *sSystemLog {
	return &sSystemLog{}
}

func (s *sSystemLog) List(page int) (list []*v1.SystemLog, max int, err error) {

	//var list =   []*AdminSystemLog{}
	max, _ = g.Model(name).Count()
	g.Model(name).OrderDesc("created_at").Page(page, 100).Scan(&list)
	return
}

// 写入操作日志
func (s *sSystemLog) AddLog(uid int, url string, ip string, data g.Map) (id int64, err error) {
	//跳过空日志
	if data == nil {
		return
	}
	//如果存在这些值，直接跳过不写入日志
	paichu := []string{
		"/api/install",
		"/activity/url/log/add",
		"/system/update",
		"/api/cdkey",
	}

	for _, item := range paichu {
		if item == url {
			return
		}
	}

	var post v1.SystemLog
	//uid := g.RequestFromCtx(ctx).Header.Get("x-uid")
	post.Uid = uid
	post.Url = url
	post.Ip = ip
	post.Data = data

	id, err = g.Model(name).InsertAndGetId(post)
	return
}
