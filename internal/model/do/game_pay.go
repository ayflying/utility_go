// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GamePay is the golang structure of table shiningu_game_pay for DAO operations like Where/Data.
type GamePay struct {
	g.Meta         `orm:"table:shiningu_game_pay, do:true"`
	OrderId        interface{} // 订单编号
	Uid            interface{} //
	TerraceOrderId interface{} // 平台订单id
	Device         interface{} // 设备名称
	Channel        interface{} // 支付渠道
	ShopId         interface{} // 商品id
	Cent           interface{} // 美分(不要使用小数点)
	PackageName    interface{} // 包名
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	PayTime        *gtime.Time // 支付时间
	Status         interface{} // 状态
	Token          interface{} // 支付标识
	Ip             interface{} // ip地址
}
