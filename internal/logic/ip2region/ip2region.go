package ip2region

import (
	"net"
	"path"
	"strings"

	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	ctx  = gctx.New()
	wait = false
)

const IpDbPath = "runtime/library"

type sIp2region struct {
	//searcher  *xdb.Searcher
	searchers map[string]*xdb.Searcher
}

func New() *sIp2region {
	return &sIp2region{}
}

func init() {
	service.RegisterIp2Region(New())

	//boot.AddFunc(func() {
	//	service.Ip2Region().Load()
	//})
}

//func (s *sIp2region) New() *xdb.Searcher {
//
//	return nil
//}

// Load 加载到内存中
//
//	@Description: 加载ip2region数据库到内存中。

//	@receiver s *sIp2region: sIp2region的实例。
func (s *sIp2region) Load(t *xdb.Version) {
	var err error
	var url string
	//var dbPath = "runtime/library/ip2region.xdb"
	switch t {
	case xdb.IPv4:
		//url = "https://github.com/ayflying/resource/raw/refs/heads/master/attachment/ip2region_v4.xdb"
		url = "https://github.com/lionsoul2014/ip2region/raw/refs/heads/master/data/ip2region_v4.xdb"
	case xdb.IPv6:
		url = "https://github.com/lionsoul2014/ip2region/raw/refs/heads/master/data/ip2region_v6.xdb"
	}

	if wait {
		return
	}
	filename := gfile.Basename(url)
	var IpDbFile = path.Join(IpDbPath, filename)
	g.Log().Debugf(ctx, "加载ip库文件:%v", filename)
	if gfile.IsEmpty(IpDbFile) {
		wait = true
		defer func() {
			wait = false
		}()
		g.Log().Debug(ctx, "等待下载ip库文件")
		//下载文件
		putData, err2 := g.Client().Get(ctx, url)
		if err2 != nil {
			return
		}
		err = gfile.PutBytes(IpDbFile, putData.ReadAll())
	}

	err = xdb.VerifyFromFile(IpDbFile)
	if err != nil {
		// err 包含的验证的错误
		gfile.RemoveFile(IpDbFile)
		g.Log().Errorf(ctx, "ip库xdb file verify: %v", err)
		return
	}

	// 1、从 dbPath 加载 VectorIndex 缓存，把下述 vIndex 变量全局到内存里面。
	vIndex, err := xdb.LoadVectorIndexFromFile(IpDbFile)
	if err != nil {
		g.Log().Errorf(ctx, "failed to load vector index from `%s`: %s\n", IpDbFile, err)
		return
	}
	// 2、用全局的 vIndex 创建带 VectorIndex 缓存的查询对象。
	if s.searchers == nil {
		s.searchers = make(map[string]*xdb.Searcher)
	}
	s.searchers[t.Name], err = xdb.NewWithVectorIndex(t, IpDbFile, vIndex)
	if err != nil {
		g.Log().Errorf(ctx, "failed to create searcher with vector index: %s\n", err)
		return
	}

	//cBuff := gfile.GetBytes(IpDbFile)
	//// 基于读取的内容，创建查询对象
	//s.searchers[t.Name], err = xdb.NewWithBuffer(t, cBuff)
	//if err != nil {
	//	g.Log().Errorf(ctx, "无法创建内容为的搜索器： %s", err)
	//	return
	//}

}

func (s *sIp2region) GetIp(ip string) (res []string) {
	//初始化加载
	//if s.searcher == nil {
	//	s.Load(xdb.IPv6)
	//}
	var searchers *xdb.Searcher
	//区分ipv6与ipv4
	if s.isIPv6(ip) {
		if s.searchers[xdb.IPv6.Name] == nil {
			s.Load(xdb.IPv6)
		}
		searchers = s.searchers[xdb.IPv6.Name]
	} else {
		if s.searchers[xdb.IPv4.Name] == nil {
			s.Load(xdb.IPv4)
		}
		searchers = s.searchers[xdb.IPv4.Name]
	}

	res = make([]string, 5)
	if searchers == nil {
		return
	}

	region, err := searchers.SearchByStr(ip)
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
