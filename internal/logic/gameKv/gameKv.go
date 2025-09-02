package gameKv

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ayflying/utility_go/pkg"
	"github.com/ayflying/utility_go/service"
	"github.com/ayflying/utility_go/tools"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Name       = "game_kv"
	RunTimeMax *gtime.Time
)

type sGameKv struct {
	Lock sync.Mutex
}

func New() *sGameKv {
	return &sGameKv{}
}

func init() {
	service.RegisterGameKv(New())
}

// SavesV1 方法
//
// @Description: 保存用户KV数据列表。
// @receiver s: sGameKv的实例。
// @return err: 错误信息，如果操作成功，则为nil。
func (s *sGameKv) SavesV1() (err error) {
	var ctx = gctx.New()
	// 最大允许执行时间
	RunTimeMax = gtime.Now().Add(time.Minute * 30)
	g.Log().Debug(ctx, "开始执行游戏kv数据保存")

	// 定义用于存储用户数据的结构体
	type ListData struct {
		Uid int64       `json:"uid"`
		Kv  interface{} `json:"kv"`
	}
	var list []*ListData
	// 初始化列表，长度与keys数组一致
	list = make([]*ListData, 0)

	// 从Redis列表中获取所有用户KV索引的键
	//keys, err := utils.RedisScan("user:kv:*")
	err = tools.Redis.RedisScanV2("user:kv:*", func(keys []string) (err error) {
		//判断是否超时
		if gtime.Now().After(RunTimeMax) {
			g.Log().Error(ctx, "kv执行超时了,停止执行！")
			err = errors.New("kv执行超时了,停止执行！")
			return
		}

		//需要删除的key

		// 遍历keys，获取每个用户的数据并填充到list中
		for _, cacheKey := range keys {
			//g.Log().Infof(ctx, "保存用户kv数据%v", v)
			//uid := v.Int64()
			//cacheKey = "user:kv:" + strconv.FormatInt(uid, 10)
			result := strings.Split(cacheKey, ":")
			var uid int64
			uid, err = strconv.ParseInt(result[2], 10, 64)
			if err != nil {
				g.Log().Error(ctx, err)
				g.Redis().Del(ctx, cacheKey)
				continue
			}

			////如果1天没有活跃，跳过
			//user, _ := service.MemberUser().Info(uid)
			//if user.UpdatedAt.Seconds < gtime.Now().Add(consts.ActSaveTime).Unix() {
			//	continue
			//}
			//如果有活跃，跳过持久化
			if getBool, _ := pkg.Cache("redis").
				Contains(ctx, fmt.Sprintf("act:update:%d", uid)); getBool {
				continue
			}

			get, _ := g.Redis().Get(ctx, cacheKey)
			var data interface{}
			get.Scan(&data)
			list = append(list, &ListData{
				Uid: uid,
				Kv:  data,
			})
		}

		// 将列表数据保存到数据库
		if len(list) > 100 {
			_, err2 := g.Model("game_kv").Data(list).Save()

			if err2 != nil {
				g.Log().Error(ctx, err2)
				return
			}
			//删除当前key
			for _, v := range list {
				go s.DelCacheKey(ctx, v.Uid)
			}
			list = make([]*ListData, 0)
		}
		if err != nil {
			g.Log().Error(ctx, "当前kv数据入库失败: %v", err)
		}

		return
	})

	return
}

// 删除缓存key
func (s *sGameKv) DelCacheKey(ctx context.Context, uid int64) {
	cacheKey := fmt.Sprintf("user:kv:%v", uid)
	_, err := g.Redis().Del(ctx, cacheKey)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
