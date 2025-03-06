package act

import (
	"fmt"
	"github.com/ayflying/utility_go/package/aycache"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

var (
	Cache           = aycache.New()
	ActIdListIsShow map[int]func(uid int64) bool
	RedDotList      map[string]func(uid int64) int32
)

func GetCacheKey(uid int64, actId int) string {
	return fmt.Sprintf("actRedDot:%s:%d:%d", gtime.Now().Format("Ymd"), actId, uid)
}

// 刷新缓存
func RefreshCache(uid int64, actId int) {

	Cache.Remove(gctx.New(), GetCacheKey(uid, actId))
	service.GameAct().RefreshGetRedDotCache(uid)
}

func GetRedDot(uid int64, actId int) *gvar.Var {
	get, _ := Cache.Get(nil, GetCacheKey(uid, actId))
	return get
}

func SetRedDot(uid int64, actId int, redDot int32) {
	Cache.Set(nil, GetCacheKey(uid, actId), redDot, time.Hour)
}

// 注册隐藏活动接口
func AddIsShowRegistrar(actId int, isShow func(uid int64) bool) {
	if ActIdListIsShow == nil {
		ActIdListIsShow = make(map[int]func(uid int64) bool)
	}
	ActIdListIsShow[actId] = isShow
}

// 注册红点接口
func AddRedDotRegistrar(key string, redDot func(uid int64) int32) {
	if RedDotList == nil {
		RedDotList = make(map[string]func(uid int64) int32)
	}
	RedDotList[key] = redDot
}
