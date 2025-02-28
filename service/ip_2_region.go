// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IIp2Region interface {
		// Load 加载到内存中
		//
		//	@Description: 加载ip2region数据库到内存中。
		//	@receiver s *sIp2region: sIp2region的实例。
		Load()
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
