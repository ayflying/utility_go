// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IGameKv interface {
		// SavesV1 方法
		//
		// @Description: 保存用户KV数据列表。
		// @receiver s: sGameKv的实例。
		// @return err: 错误信息，如果操作成功，则为nil。
		SavesV1() (err error)
	}
)

var (
	localGameKv IGameKv
)

func GameKv() IGameKv {
	if localGameKv == nil {
		panic("implement not found for interface IGameKv, forgot register?")
	}
	return localGameKv
}

func RegisterGameKv(i IGameKv) {
	localGameKv = i
}
