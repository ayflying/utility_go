// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameConfig is the golang structure for table game_config.
type GameConfig struct {
	Name      string      `json:"name"       orm:"name"       description:"配置名称"` // 配置名称
	Data      string      `json:"data"       orm:"data"       description:"配置内容"` // 配置内容
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"` // 更新时间
}
