// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GamePay is the golang structure for table game_pay.
type GamePay struct {
	OrderId        string      `json:"order_id"         orm:"order_id"         description:"订单编号"`        // 订单编号
	Uid            int64       `json:"uid"              orm:"uid"              description:""`            //
	TerraceOrderId string      `json:"terrace_order_id" orm:"terrace_order_id" description:"平台订单id"`      // 平台订单id
	Device         string      `json:"device"           orm:"device"           description:"设备名称"`        // 设备名称
	Channel        string      `json:"channel"          orm:"channel"          description:"支付渠道"`        // 支付渠道
	ShopId         int64       `json:"shop_id"          orm:"shop_id"          description:"商品id"`        // 商品id
	Cent           int         `json:"cent"             orm:"cent"             description:"美分(不要使用小数点)"` // 美分(不要使用小数点)
	PackageName    string      `json:"package_name"     orm:"package_name"     description:"包名"`          // 包名
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`        // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"更新时间"`        // 更新时间
	PayTime        *gtime.Time `json:"pay_time"         orm:"pay_time"         description:"支付时间"`        // 支付时间
	Status         int         `json:"status"           orm:"status"           description:"状态"`          // 状态
	Token          string      `json:"token"            orm:"token"            description:"支付标识"`        // 支付标识
	Ip             string      `json:"ip"               orm:"ip"               description:"ip地址"`        // ip地址
}
