package v1

import "github.com/gogf/gf/v2/frame/g"

type IpReq struct {
	g.Meta `path:"/callback/ip/{ip}" tags:"回调响应" method:"get" summary:"获取ip"`
	Ip     string `json:"ip" dc:"ip"`
}
type IpRes struct {
	g.Meta  `mime:"application/json" example:"string"`
	Address []string `json:"address" dc:"地区名"`
}

type Ip struct {
	Country  string `json:"country" dc:"国家"`  //国家
	Region   string `json:"region" dc:"地区"`   //地区
	Province string `json:"province" dc:"省份"` //省份
	City     string `json:"city" dc:"城市"`     //城市
	District string `json:"district" dc:"区县"` //区县
}
