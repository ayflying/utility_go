// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SystemStatistics is the golang structure of table shiningu_system_statistics for DAO operations like Where/Data.
type SystemStatistics struct {
	g.Meta `orm:"table:shiningu_system_statistics, do:true"`
	Id     interface{} // 流水号
	AppId  interface{} // 应用编号
	Key    interface{} // 唯一缓存key
	Data   interface{} // 数据
}
