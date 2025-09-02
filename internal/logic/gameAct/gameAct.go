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
	var ctx = gctx.New()
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
	var ctx = gctx.New()
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

// Saves 保存游戏活动数据
//
// @Description: 保存游戏活动数据
// @receiver s *sGameAct: 游戏活动服务结构体指针
// @return err error: 返回错误信息
// Deprecated: 该方法已被弃用，建议使用SavesV2方法
func (s *sGameAct) Saves() (err error) {
	var ctx = gctx.New()
	g.Log().Debug(ctx, "开始执行游戏act数据保存了")
	//如果没有执行过，设置时间戳
	// 最大允许执行时间
	RunTimeMax = gtime.Now().Add(time.Minute * 30)
	//遍历执行
	ActList.Iterator(func(i interface{}) bool {
		//在时间内允许执行
		if gtime.Now().Before(RunTimeMax) {
			g.Log().Debugf(ctx, "开始执行游戏act数据保存:act=%v", i)
			err = s.Save(ctx, i.(int))
		} else {
			g.Log().Errorf(ctx, "游戏act数据保存超时:act=%v", i)
		}
		return true
	})
	return
}

// Save 保存游戏活动数据
//
// @Description: 保存游戏活动数据
// @receiver s *sGameAct: 游戏活动服务结构体指针
// @param ctx context.Context: 上下文对象
// @param actId int: 活动ID
// @return err error: 返回错误信息
// deprecated: 该方法已被弃用，建议使用SaveV2方法
func (s *sGameAct) Save(ctx context.Context, actId int) (err error) {

	cacheKey := fmt.Sprintf("act:%v:*", actId)
	var add = make([]*entity.GameAct, 0)
	var update = make([]*entity.GameAct, 0)
	//循环获取缓存数据
	err = tools.Redis.RedisScanV2(cacheKey, func(keys []string) (err error) {
		//判断是否超时
		if gtime.Now().After(RunTimeMax) {
			g.Log().Debug(ctx, "act执行超时了,停止执行！")
			err = errors.New("act执行超时了,停止执行！")
			return
		}

		for _, cacheKey = range keys {
			result := strings.Split(cacheKey, ":")
			actId, err = strconv.Atoi(result[1])
			var uid int64
			uid = gconv.Int64(result[2])
			//uid, err = strconv.ParseInt(result[2], 10, 64)
			if err != nil {
				g.Log().Error(ctx, err)
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
				g.Log().Errorf(ctx, "当前数据错误: %v", cacheKey)
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
		}

		//批量写入数据库
		updateCount := 0

		//g.Log().Debugf(ctx, "当前 %v 要更新的数据: %v 条", actId, len(update))
		if len(update) > 100 {
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
					continue
				}

				//删除缓存
				go s.DelCacheKey(ctx, v.ActId, v.Uid)

				updateCount++
				update = make([]*entity.GameAct, 0)
			}
			g.Log().Debugf(ctx, "当前 %v 更新数据库: %v 条", actId, updateCount)
			update = make([]*entity.GameAct, 0)
		}

		var count int64
		//g.Log().Debugf(ctx, "当前 %v 要添加的数据: %v 条", actId, len(add))
		if len(add) > 100 {
			dbRes, err2 := g.Model(Name).Data(add).Save()

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

			for _, v2 := range add {
				//删除缓存
				go s.DelCacheKey(ctx, v2.ActId, v2.Uid)
			}

			//g.Log().Debugf(ctx, "当前 %v 写入数据库: %v 条", actId, count)
			add = make([]*entity.GameAct, 0)
		}

		if err != nil {
			g.Log().Error(ctx, "当前临时数据入库失败: %v", err)
		}

		return err
	})

	return
}

// SavesV2 保存游戏活动数据
//
// @Description: 保存游戏活动数据
// @receiver s *sGameAct: 游戏活动服务结构体指针
// @return err error: 返回错误信息
func (s *sGameAct) SavesV2() (err error) {
	var ctx = gctx.New()
	g.Log().Debug(ctx, "开始执行游戏act数据保存了")
	//如果没有执行过，设置时间戳
	// 最大允许执行时间
	RunTimeMax = gtime.Now().Add(time.Minute * 30)

	//cacheKey := fmt.Sprintf("act:%v:*", actId)
	var add = make([]*entity.GameAct, 0)
	var update = make([]*entity.GameAct, 0)

	//循环获取缓存数据
	err = tools.Redis.RedisScanV2("act:*", func(keys []string) (err error) {
		for _, key := range keys {
			//格式化数据
			err = s.SaveV2(ctx, key, add, update)
			//持久化数据
			err = s.Cache2Sql(ctx, add, update)
		}
		return err
	})

	return
}

// SaveV2 保存游戏活动数据
//
// @Description: 保存游戏活动数据
// @receiver s *sGameAct: 游戏活动服务结构体指针
// @param ctx context.Context: 上下文对象
// @param cacheKey string: 缓存键
// @param add []*entity.GameAct: 添加数据
// @param update []*entity.GameAct: 更新数据
// @return err error: 返回错误信息
func (s *sGameAct) SaveV2(ctx context.Context, cacheKey string, add, update []*entity.GameAct) (err error) {

	result := strings.Split(cacheKey, ":")
	actId := gconv.Int(result[1])
	if actId == 0 {
		return
	}
	var uid int64
	uid = gconv.Int64(result[2])
	if uid == 0 {
		//跳过为空的用户缓存
		return
	}

	//获取缓存数据
	cacheGet, _ := g.Redis().Get(ctx, cacheKey)

	if cacheGet.IsEmpty() {
		//空数据也不保存
		return
	}

	//如果有活跃，跳过持久化
	if getBool, _ := pkg.Cache("redis").
		Contains(ctx, fmt.Sprintf("act:update:%d", uid)); getBool {
		return
	}

	//获取数据库数据
	var data *entity.GameAct
	// 从数据库中查询活动信息
	err = g.Model(Name).Where(do.GameAct{
		Uid:   uid,
		ActId: actId,
	}).Fields("uid,act_id").Scan(&data)
	if err != nil {
		g.Log().Errorf(ctx, "当前数据错误: %v", cacheKey)
		return
	}
	//如果没有数据，添加
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

	return
}

// Cache2Sql 缓存持久化到数据库
// @Description: 缓存持久化到数据库
// @receiver s *sGameAct: 游戏活动服务结构体指针
// @param ctx context.Context: 上下文对象
// @param add []*entity.GameAct: 添加数据
// @param update []*entity.GameAct: 更新数据
// @return err error: 返回错误信息
func (s *sGameAct) Cache2Sql(ctx context.Context, add, update []*entity.GameAct) (err error) {
	//批量写入数据库
	updateCount := 0
	if len(update) > 100 {
		for _, v := range update {
			v.UpdatedAt = gtime.Now()
			updateRes, err2 := g.Model(Name).Where(do.GameAct{
				Uid:   v.Uid,
				ActId: v.ActId,
			}).Data(v).Update()
			if err2 != nil {
				g.Log().Error(ctx, err2)
				continue
			}

			if row, _ := updateRes.RowsAffected(); row == 0 {
				g.Log().Error(ctx, "本次更新为0，更新数据失败: %v", v)
				continue
			}

			//删除缓存
			go s.DelCacheKey(ctx, v.ActId, v.Uid)

			updateCount++
			update = make([]*entity.GameAct, 0)
		}

		g.Log().Debugf(ctx, "act当前更新数据库: %v 条", updateCount)
		update = make([]*entity.GameAct, 0)
	}

	var addCount int64
	if len(add) > 100 {
		for _, v := range add {
			addRes, err2 := g.Model(Name).Data(v).Insert()
			if err2 != nil {
				g.Log().Error(ctx, err2)
				continue
			}
			if row, _ := addRes.RowsAffected(); row == 0 {
				g.Log().Error(ctx, "本次新增为0，新增数据失败: %v", v)
				continue
			}
			addCount++
			//删除缓存
			go s.DelCacheKey(ctx, v.ActId, v.Uid)
		}
		g.Log().Debugf(ctx, "act当前写入数据库: %v 条", addCount)
		add = make([]*entity.GameAct, 0)
	}
	return
}

// 删除缓存key
func (s *sGameAct) DelCacheKey(ctx context.Context, aid int, uid int64) {
	cacheKey := fmt.Sprintf("act:%v:%v", aid, uid)
	_, err := g.Redis().Del(ctx, cacheKey)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// 清空GetRedDot缓存
func (s *sGameAct) RefreshGetRedDotCache(uid int64) {
	cacheKey := fmt.Sprintf("gameAct:GetRedDot:%s:%d", gtime.Now().Format("d"), uid)
	_, err := pkg.Cache("redis").Remove(gctx.New(), cacheKey)
	if err != nil {
		g.Log().Error(gctx.New(), err)
		g.Dump(err)
	}
}

func (s *sGameAct) Del(uid int64, actId int) {
	var ctx = gctx.New()
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
