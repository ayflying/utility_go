package cmd

import (
	"context"
	"github.com/ayflying/utility_go/pkg"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"os"
	"path"
	"time"
)

type serverCfg struct {
	Name    string `json:"name" dc:"服务名"`
	Address string `json:"address" dc:"服务地址"`
	Prod    bool   `json:"prod" dc:"是否生产环境"`
	S3      string `json:"s3" dc:"使用哪个对象储存中转"`
	Arch    string `json:"arch" dc:"架构"`
	System  string `json:"system" dc:"系统"`
}

type UpdateReq struct {
	File    *ghttp.UploadFile `json:"file" binding:"required" dc:"文件"`
	FileUrl string            `json:"file_url" dc:"文件地址"`
}

var (
	Update = gcmd.Command{
		Name:  "update",
		Usage: "update",
		Brief: "更新系统",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			g.Log().Info(ctx, "准备上传更新文件")
			//加载编辑配置文件
			g.Cfg("hack").GetAdapter().(*gcfg.AdapterFile).SetFileName("hack/config.yaml")
			//获取文件名
			getFileName, err := g.Cfg("hack").Get(ctx, "gfcli.build.name")
			Filename := getFileName.String()

			////获取架构
			//getArch, err := g.Cfg("hack").Get(ctx, "gfcli.build.arch")
			//Arch := getArch.String()
			//if Arch == "" {
			//	Arch = "amd64"
			//}
			//
			////获取操作系统
			//getSystem, err := g.Cfg("hack").Get(ctx, "gfcli.build.system")
			//System := getSystem.String()
			//
			//if System == "" {
			//	System = "linux"
			//}
			//var systemName = System + "_" + Arch

			//获取版本号
			getVersion, err := g.Cfg("hack").Get(ctx, "gfcli.build.version")
			Version := getVersion.String()

			var list []*serverCfg
			serverList := g.Cfg().MustGet(ctx, "server_list")
			serverList.Scan(&list)

			//如果有p或者prod参数，则删除prod字段为true的服务
			if parser.GetOpt("a").IsNil() {
				var temp []*serverCfg
				for _, v := range list {
					if v.Prod == false {
						temp = append(temp, v)
					}
				}
				list = temp
			} else {
				g.Dump("升级", parser.GetOpt("a"))
			}

			g.Dump("需要更新的服务器", list)
			//获取上传链接
			var url = make(map[string]string)
			var system = make(map[string]string)
			//filename := "linux_amd64/" + Filename
			//filename := path.Join(Version, "linux_amd64", Filename)

			client := g.Client()
			client.SetTimeout(time.Minute)
			client.SetDiscovery(nil)

			//循环服务器，推送更新
			for _, v := range list {
				address := v.Address
				if v.S3 == "" {
					v.S3 = "default"
				}

				//查询当前上传地址是否存在
				_, ok2 := system[v.System+v.Arch]
				if _, ok := url[v.S3]; !ok || !ok2 {

					var systemName = "linux_amd64"
					if v.Arch != "" && v.System != "" {
						systemName = v.System + "_" + v.Arch
					}
					var filename = path.Join(Version, systemName, Filename)
					g.Log().Debugf(ctx, "当前上传文件："+filename)

					url[v.S3], err = UploadS3(v.S3, filename)
					if err != nil {
						g.Log().Error(ctx, err)
						return
					}
					system[v.System+v.Arch] = filename

					if err != nil {
						g.Log().Error(ctx, err)
						return
					}
				}
				g.Log().Debugf(ctx, "准备同步服务器:%v,url=%v", v.Name, address+"/callback/update")
				get, err := client.Post(ctx, address+"/callback/update", &UpdateReq{
					FileUrl: url[v.S3],
				})
				if err != nil {
					g.Log().Debugf(ctx, "切换代理进行上传:err=%v", err)
					get, err = client.Proxy("http://192.168.50.114:10808").
						Post(ctx, address+"/callback/update", &UpdateReq{
							FileUrl: url[v.S3],
						})
				}
				if err != nil {
					g.Log().Error(ctx, err)
				}
				defer get.Close()
				g.Log().Debugf(ctx, "同步服务器:%v,完成=%v", v.Name, address)
			}

			return
		},
	}
)

func UploadS3(typ string, filename string) (res string, err error) {
	//updateServerS3Name, _ := g.Config().Get(ctx, "update_server_s3_name")

	var s3Mod = pkg.S3(typ)
	bucketName := s3Mod.GetCfg().BucketName
	obj, err := os.Open(filename)
	if err != nil {
		return
	}
	defer obj.Close()
	ff, err := obj.Stat()
	_, err = s3Mod.PutObject(obj, filename, bucketName, ff.Size())
	if err != nil {
		return
	}
	//上传当前文件
	get, err := s3Mod.GetFileUrl(filename, bucketName)
	g.Log().Debugf(gctx.New(), "下载地址:%s", get)

	res = get.String()
	return
}
