package gameAct

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ayflying/utility_go/internal/model/do"
	"github.com/ayflying/utility_go/internal/model/entity"
	"github.com/ayflying/utility_go/pkg"
	service2 "github.com/ayflying/utility_go/service"
	"github.com/ayflying/utility_go/tools"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ctx        = gctx.New()
	Name       = "game_act"
	ActList    = gset.New(true)
	RunTimeMax *gtime.Time
)

type sGameAct struct {
}

func new() *sGameAct {
	return &sGameAct{}
}

func init() {
	service2.RegisterGameAct(new())
}

// Info 获取活动信息
//
// @Description: 根据用户ID和活动ID获取活动信息
// @receiver s *sGameAct: 代表活动操作的结构体实例
// @param uid int64: 用户ID
// @param actId int: 活动ID
// @return data *v1.Act: 返回活动信息结构体指针
// @return err error: 返回错误信息
func (s *sGameAct) Info(uid int64, actId int) (data *g.Var, err error) {
	if uid == 0 || actId == 0 {
		g.Log().Error(ctx, "当前参数为空")
		return
	}

	// 构造缓存键名
	keyCache := fmt.Sprintf("act:%v:%v", actId, uid)
	// 尝试从Redis缓存中获取活动信息
	get, err := g.Redis().Get(ctx, keyCache)
	if !get.IsEmpty() {
		// 如果缓存中存在，将数据扫描到data结构体中并返回
		data = get
		return
	}
	// 从数据库中查询活动信息
	getDb, err := g.Model(Name).Where(do.GameAct{
		Uid:   uid,
		ActId: actId,
	}).Fields("action").OrderDesc("updated_at").Value()
	getDb.Scan(&data)

	if data == nil || data.IsEmpty() {
		return
	}
	// 将查询到的活动信息保存到Redis缓存中
	_, err = g.Redis().Set(ctx, keyCache, data)

	var CacheKey = fmt.Sprintf("act:update:%d", uid)
	pkg.Cache("redis").Set(ctx, CacheKey, uid, time.Hour*24*1)

	return
}

// Set 将指定用户的活动信息存储到Redis缓存中。
//
// @Description:
// @receiver s *sGameAct: 表示sGameAct类型的实例。
// @param uid int64: 用户的唯一标识。
// @param actId int: 活动的唯一标识。
// @param data interface{}: 要存储的活动信息数据。
// @return err error: 返回错误信息，如果操作成功，则返回nil。
func (s *sGameAct) Set(uid int64, actId int, data interface{}) (err error) {
	if uid == 0 || actId == 0 {
		g.Log().Error(ctx, "当前参数为空")
		return
	}
	// 构造缓存键名
	keyCache := fmt.Sprintf("act:%v:%v", actId, uid)
	if data == nil {
		_, err = g.Redis().Del(ctx, keyCache)
		return
	}

	// 将活动信息保存到Redis缓存，并将用户ID添加到活动索引集合中
	_, err = g.Redis().Set(ctx, keyCache, data)

	//插入集合
	ActList.Add(actId)

	return
}

func (s *sGameAct) Saves(ctx context.Context) (err error) {
	getCache, _ := pkg.Cache("redis").Get(nil, "cron:game_act")
	g.Log().Debug(ctx, "开始执行游戏act数据保存了")
	//如果没有执行过，设置时间戳
	if getCache.Int64() > 0 {
		return
	} else {
		pkg.Cache("redis").Set(nil, "cron:game_act", gtime.Now().Unix(), time.Hour)
	}
	// 最大允许执行时间
	RunTimeMax = gtime.Now().Add(time.Minute * 30)
	//遍历执行
	ActList.Iterator(func(i interface{}) bool {
		//在时间内允许执行
		if gtime.Now().Before(RunTimeMax) {
			g.Log().Debug(ctx, "开始执行游戏act数据保存: act%v", i)
			err = s.Save(ctx, i.(int))
		} else {
			g.Log().Errorf(ctx, "游戏act数据保存超时: act=%v", i)
		}
		return true
	})
	return
}

func (s *sGameAct) Save(ctx context.Context, actId int) (err error) {

	cacheKey := fmt.Sprintf("act:%v:*", actId)
	//获取当前用户的key值
	//keys, err := utils.RedisScan(cacheKey)
	//if len(keys) > 10000 {
	//	keys = keys[:10000]
	//}

	//循环获取缓存数据
	err = tools.Redis.RedisScanV2(cacheKey, func(keys []string) (err error) {
		//判断是否超时
		if gtime.Now().After(RunTimeMax) {
			g.Log().Debug(ctx, "act执行超时了,停止执行！")
			err = errors.New("act执行超时了,停止执行！")
			return
		}
		var add = make([]*entity.GameAct, 0)
		var update = make([]*entity.GameAct, 0)
		var delKey []string
		for _, cacheKey = range keys {
			result := strings.Split(cacheKey, ":")
			actId, err = strconv.Atoi(result[1])
			var uid int64
			uid = gconv.Int64(result[2])
			//uid, err = strconv.ParseInt(result[2], 10, 64)
			if err != nil {
				continue
			}

			cacheGet, _ := g.Redis().Get(ctx, cacheKey)

			if uid == 0 {
				//跳过为空的用户缓存
				continue
			}
			if cacheGet.IsEmpty() {
				//空数据也不保存
				continue
			}

			//如果有活跃，跳过持久化
			if getBool, _ := pkg.Cache("redis").
				Contains(ctx, fmt.Sprintf("act:update:%d", uid)); getBool {
				continue
			}

			//获取数据库数据
			var data *entity.GameAct
			// 从数据库中查询活动信息
			err = g.Model(Name).Where(do.GameAct{
				Uid:   uid,
				ActId: actId,
			}).Fields("uid,act_id").Scan(&data)
			if err != nil {
				g.Log().Debugf(ctx, "当前数据错误: %v", cacheKey)
				continue
			}
			actionData := cacheGet.String()
			if data == nil {
				add = append(add, &entity.GameAct{
					ActId:  actId,
					Uid:    uid,
					Action: actionData,
				})
			} else {
				//覆盖数据
				data.ActId = actId
				data.Uid = uid
				data.Action = actionData
				update = append(update, data)
			}
			//最后删除key
			delKey = append(delKey, cacheKey)
		}

		//批量写入数据库
		updateCount := 0
		if len(delKey) > 200 {
			for _, v := range update {
				v.UpdatedAt = gtime.Now()
				updateRes, err2 := g.Model(Name).Where(do.GameAct{
					Uid:   v.Uid,
					ActId: v.ActId,
				}).Data(v).Update()
				if err2 != nil {
					g.Log().Error(ctx, err2)
					return
				}
				if row, _ := updateRes.RowsAffected(); row == 0 {
					g.Log().Error(ctx, "本次更新为0，更新数据失败: %v", v)
					return
				}
				updateCount++
			}
			g.Log().Debugf(ctx, "当前 %v 更新数据库: %v 条", actId, updateCount)

			update = make([]*entity.GameAct, 0)
			var count int64

			if len(add) > 0 {
				dbRes, err2 := g.Model(Name).Data(add).Save()
				add = make([]*entity.GameAct, 0)
				err = err2
				if err != nil {
					g.Log().Error(ctx, err2)
					return
				}
				count, _ = dbRes.RowsAffected()
				if count == 0 {
					g.Log().Error(ctx, "当前 %v 写入数据库: %v 条", actId, count)
					for _, vTemp := range add {
						g.Log().Debugf(ctx, "当前act：%v，add写入数据: %v,内容：%v", vTemp.ActId, vTemp.Uid, vTemp.Action)
					}
					return

				}
				//g.Log().Debugf(ctx, "当前 %v 写入数据库: %v 条", actId, count)
			}

			for _, v := range delKey {
				_, err = g.Redis().Del(ctx, v)
				if err != nil {
					g.Log().Error(ctx, err)
				}
			}
			delKey = make([]string, 0)

		}

		if err != nil {
			g.Log().Error(ctx, "当前临时数据入库失败: %v", err)
		}

		return err
	})

	return
}

// 清空GetRedDot缓存
func (s *sGameAct) RefreshGetRedDotCache(uid int64) {
	cacheKey := fmt.Sprintf("gameAct:GetRedDot:%s:%d", gtime.Now().Format("d"), uid)
	_, err := pkg.Cache("redis").Remove(gctx.New(), cacheKey)
	if err != nil {
		g.Log().Error(ctx, err)
		g.Dump(err)
	}
}

func (s *sGameAct) Del(uid int64, actId int) {
	if uid == 0 || actId == 0 {
		g.Log().Error(ctx, "当前参数为空")
		return
	}
	// 构造缓存键名
	keyCache := fmt.Sprintf("act:%v:%v", actId, uid)

	//删除活动缓存
	g.Redis().Del(ctx, keyCache)

	//删除当前活动储存
	g.Model(Name).Where(do.GameAct{
		Uid:   uid,
		ActId: actId,
	}).Delete()

}
