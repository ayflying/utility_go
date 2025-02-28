// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemSetting is the golang structure for table system_setting.
type SystemSetting struct {
	Name  string `json:"name"  orm:"name"  description:"配置名称"` // 配置名称
	Value string `json:"value" orm:"value" description:"配置详情"` // 配置详情
	Type  int    `json:"type"  orm:"type"  description:"类型"`   // 类型
}
