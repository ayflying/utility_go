package gameAct

import (
	"fmt"
	"github.com/ayflying/utility_go/internal/model/do"
	"github.com/ayflying/utility_go/internal/model/entity"
	"github.com/ayflying/utility_go/pgk"
	"github.com/ayflying/utility_go/pkg"
	service2 "github.com/ayflying/utility_go/service"
	"github.com/ayflying/utility_go/tools"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
	"strings"
	"time"
)

var (
	ctx     = gctx.New()
	Name    = "game_act"
	ActList = gset.New(true)
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

func (s *sGameAct) Saves() (err error) {
	//遍历执行
	ActList.Iterator(func(i interface{}) bool {
		err = s.Save(i.(int))
		return true
	})
	return
}

func (s *sGameAct) Save(actId int) (err error) {

	cacheKey := fmt.Sprintf("act:%v:*", actId)
	//获取当前用户的key值
	//keys, err := utils.RedisScan(cacheKey)
	//if len(keys) > 10000 {
	//	keys = keys[:10000]
	//}

	//循环获取缓存数据
	err = tools.Redis.RedisScanV2(cacheKey, func(keys []string) (err error) {
		var add []interface{}
		var delKey []string
		for _, cacheKey = range keys {
			result := strings.Split(cacheKey, ":")
			actId, err = strconv.Atoi(result[1])
			var uid int64
			uid, err = strconv.ParseInt(result[2], 10, 64)
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
			if getBool, _ := pkg.Cache("redis").Contains(ctx,
				fmt.Sprintf("act:update:%d", uid)); getBool {
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
				//data =
				add = append(add, &do.GameAct{
					ActId:  actId,
					Uid:    uid,
					Action: actionData,
				})
			} else {
				//覆盖数据
				data.Action = actionData
				add = append(add, data)
			}
			//最后删除key
			delKey = append(delKey, cacheKey)
		}

		//批量写入数据库
		if len(add) > 0 {
			dbRes, err2 := g.Model(Name).Batch(30).Data(add).Save()
			add = make([]interface{}, 0)
			if err2 != nil {
				g.Log().Error(ctx, err2)
				return
			}

			for _, v := range delKey {
				_, err2 = g.Redis().Del(ctx, v)
				if err2 != nil {
					g.Log().Error(ctx, err2)
					return
				}
			}
			delKey = make([]string, 0)

			count, _ := dbRes.RowsAffected()
			g.Log().Debugf(ctx, "当前 %v 写入数据库: %v 条", actId, count)
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
	_, err := pgk.Cache("redis").Remove(gctx.New(), cacheKey)
	if err != nil {
		g.Log().Error(ctx, err)
		g.Dump(err)
	}
}
