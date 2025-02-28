// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ILogData interface {
		Load()
		// UserSet 方法
		//
		// @Description: 设置用户信息。
		// @receiver s: sLogData 的实例，表示日志数据的结构体。
		// @param accountId: 账户ID，用于标识账户，是字符串格式。
		// @param uid: 用户的ID，是整型的唯一标识符。
		// @param data: 要设置的用户信息，以键值对的形式提供，是map[string]interface{}类型，支持多种用户属性。
		// @return err: 执行过程中可能出现的错误，如果执行成功，则返回nil。
		UserSet(accountId string, uid int64, data map[string]interface{}) (err error)
		// Track 函数记录特定事件。
		//
		// @Description: 用于跟踪和记录一个指定事件的发生，收集相关数据。
		// @receiver s: sLogData 的实例，代表日志数据的存储或处理实体。
		// @param accountId: 账户ID，用于标识事件所属的账户。
		// @param uid: 用户的ID，一个整型数值，用于区分不同的用户。
		// @param name: 事件名称，标识所记录的具体事件。
		// @param data: 事件相关的数据映射，包含事件的详细信息。
		// @return err: 错误信息，如果操作成功则为nil。
		Track(ctx context.Context, accountId string, uid int64, name string, data map[string]interface{})
	}
)

var (
	localLogData ILogData
)

func LogData() ILogData {
	if localLogData == nil {
		panic("implement not found for interface ILogData, forgot register?")
	}
	return localLogData
}

func RegisterLogData(i ILogData) {
	localLogData = i
}
