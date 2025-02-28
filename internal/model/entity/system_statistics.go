// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemStatistics is the golang structure for table system_statistics.
type SystemStatistics struct {
	Id    int    `json:"id"     orm:"id"     description:"流水号"`     // 流水号
	AppId int    `json:"app_id" orm:"app_id" description:"应用编号"`    // 应用编号
	Key   string `json:"key"    orm:"key"    description:"唯一缓存key"` // 唯一缓存key
	Data  string `json:"data"   orm:"data"   description:"数据"`      // 数据
}
