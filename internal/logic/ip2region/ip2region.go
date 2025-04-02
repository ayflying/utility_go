package ip2region

import (
	"github.com/ayflying/utility_go/internal/boot"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net"
	"strings"
)

var (
	ctx = gctx.New()
)

type sIp2region struct {
	searcher *xdb.Searcher
}

func New() *sIp2region {

	return &sIp2region{}
}

func init() {
	service.RegisterIp2Region(New())

	boot.AddFunc(func() {
		service.Ip2Region().Load()
	})
}

// Load 加载到内存中
//
//	@Description: 加载ip2region数据库到内存中。
//	@receiver s *sIp2region: sIp2region的实例。
func (s *sIp2region) Load() {
	var err error
	var dbPath = "/runtime/library/ip2region.xdb"

	if gfile.IsEmpty(dbPath) {
		g.Log().Debug(ctx, "等待下载ip库文件")
		//下载文件
		putData, err2 := g.Client().Discovery(nil).
			Get(ctx, "https://resource.luoe.cn/attachment/ip2region.xdb")
		if err2 != nil {
			return
		}
		err = gfile.PutBytes(dbPath, putData.ReadAll())
	}
	cBuff := gfile.GetBytes(dbPath)
	/*
		var cBuff []byte
		if gres.Contains(dbPath) {
			cBuff = gres.GetContent(dbPath)
		} else {
			cBuff = gfile.GetBytes(dbPath)
		}
	*/

	// 基于读取的内容，创建查询对象
	s.searcher, err = xdb.NewWithBuffer(cBuff)
	if err != nil {
		g.Log().Errorf(ctx, "无法创建内容为的搜索器： %s", err)
		return
	}

}

func (s *sIp2region) GetIp(ip string) (res []string) {
	res = make([]string, 5)
	if s.searcher == nil {
		return
	}

	//如果是ipv6直接跳过
	if s.isIPv6(ip) {
		return
	}

	region, err := s.searcher.SearchByStr(ip)
	if err != nil {
		return
	}
	res = strings.Split(region, "|")
	return
}

// isIPv6 判断输入字符串是否为IPv6地址
//
// @Description: 通过解析输入的IP字符串判断其是否为IPv6地址。
// @receiver s *sIp2region: 代表`sIp2region`类型的实例，本函数中未使用，可忽略。
// @param ipStr string: 待判断的IP地址字符串。
// @return bool: 返回true表示是IPv6地址，返回false表示不是IPv6地址。
func (s *sIp2region) isIPv6(ipStr string) bool {
	// 尝试将输入字符串解析为IP地址
	ip := net.ParseIP(ipStr)
	// 尝试将IP地址转换为IPv4格式
	ipv4 := ip.To4()
	// 如果转换为IPv4格式不为nil，则说明是IPv4地址，返回false
	if ipv4 != nil {
		return false
	}
	// 如果无法转换为IPv4格式，则说明是IPv6地址，返回true
	return true
}
