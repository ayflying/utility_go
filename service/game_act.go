// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/ayflying/utility_go/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IGameAct interface {
		// Info 获取活动信息
		//
		// @Description: 根据用户ID和活动ID获取活动信息
		// @receiver s *sGameAct: 代表活动操作的结构体实例
		// @param uid int64: 用户ID
		// @param actId int: 活动ID
		// @return data *v1.Act: 返回活动信息结构体指针
		// @return err error: 返回错误信息
		Info(uid int64, actId int) (data *g.Var, err error)
		// Set 将指定用户的活动信息存储到Redis缓存中。
		//
		// @Description:
		// @receiver s *sGameAct: 表示sGameAct类型的实例。
		// @param uid int64: 用户的唯一标识。
		// @param actId int: 活动的唯一标识。
		// @param data interface{}: 要存储的活动信息数据。
		// @return err error: 返回错误信息，如果操作成功，则返回nil。
		Set(uid int64, actId int, data interface{}) (err error)
		// Saves 保存游戏活动数据
		//
		// @Description: 保存游戏活动数据
		// @receiver s *sGameAct: 游戏活动服务结构体指针
		// @return err error: 返回错误信息
		// Deprecated: 该方法已被弃用，建议使用SavesV2方法
		Saves() (err error)
		// Save 保存游戏活动数据
		//
		// @Description: 保存游戏活动数据
		// @receiver s *sGameAct: 游戏活动服务结构体指针
		// @param ctx context.Context: 上下文对象
		// @param actId int: 活动ID
		// @return err error: 返回错误信息
		// deprecated: 该方法已被弃用，建议使用SaveV2方法
		Save(ctx context.Context, actId int) (err error)
		// SavesV2 保存游戏活动数据
		//
		// @Description: 保存游戏活动数据
		// @receiver s *sGameAct: 游戏活动服务结构体指针
		// @return err error: 返回错误信息
		SavesV2() (err error)
		// SaveV2 保存游戏活动数据
		//
		// @Description: 保存游戏活动数据
		// @receiver s *sGameAct: 游戏活动服务结构体指针
		// @param ctx context.Context: 上下文对象
		// @param cacheKey string: 缓存键
		// @param add []*entity.GameAct: 添加数据
		// @param update []*entity.GameAct: 更新数据
		// @return err error: 返回错误信息
		SaveV2(ctx context.Context, cacheKey string, addChan chan *entity.GameAct, updateChan chan *entity.GameAct) (err error)
		// Cache2Sql 缓存持久化到数据库
		// @Description: 缓存持久化到数据库
		// @receiver s *sGameAct: 游戏活动服务结构体指针
		// @param ctx context.Context: 上下文对象
		// @param add []*entity.GameAct: 添加数据
		// @param update []*entity.GameAct: 更新数据
		// @return err error: 返回错误信息
		Cache2Sql(ctx context.Context, add []*entity.GameAct, update []*entity.GameAct)
		// Cache2AddChan 批量添加数据库
		Cache2SqlChan(ctx context.Context, addChan chan *entity.GameAct, updateChan chan *entity.GameAct)
		// 删除缓存key
		DelCacheKey(ctx context.Context, aid int, uid int64)
		// 清空GetRedDot缓存
		RefreshGetRedDotCache(uid int64)
		Del(uid int64, actId int)
	}
)

var (
	localGameAct IGameAct
)

func GameAct() IGameAct {
	if localGameAct == nil {
		panic("implement not found for interface IGameAct, forgot register?")
	}
	return localGameAct
}

func RegisterGameAct(i IGameAct) {
	localGameAct = i
}
