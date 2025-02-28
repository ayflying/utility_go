// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service2

import (
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
		Saves() (err error)
		Save(actId int) (err error)
		// 清空GetRedDot缓存
		RefreshGetRedDotCache(uid int64)
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
