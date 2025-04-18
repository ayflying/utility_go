// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameConfig is the golang structure of table shiningu_game_config for DAO operations like Where/Data.
type GameConfig struct {
	g.Meta    `orm:"table:shiningu_game_config, do:true"`
	Name      interface{} // 配置名称
	Data      interface{} // 配置内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
