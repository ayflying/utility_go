// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type (
	IIp2Region interface {
		//	@receiver s *sIp2region: sIp2region的实例。
		Load(t *xdb.Version)
		GetIp(ip string) (res []string)
	}
)

var (
	localIp2Region IIp2Region
)

func Ip2Region() IIp2Region {
	if localIp2Region == nil {
		panic("implement not found for interface IIp2Region, forgot register?")
	}
	return localIp2Region
}

func RegisterIp2Region(i IIp2Region) {
	localIp2Region = i
}
