package cmd

import (
	"context"
	"github.com/ayflying/utility_go/s3"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"time"
)

type serverCfg struct {
	Name    string `json:"name" dc:"服务名"`
	Address string `json:"address" dc:"服务地址"`
	Prod    bool   `json:"prod" dc:"是否生产环境"`
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
			getFileName, err := g.Cfg("hack").Get(ctx, "gfcli.build.name")

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

			filename := "linux_amd64/shining_u_server"
			obj, err := os.Open(filename)
			ff, err := obj.Stat()
			var s3Mod *s3.Mod
			//切换s3目标
			updateServerS3Name, _ := g.Config().Get(ctx, "update_server_s3_name")
			if updateServerS3Name != nil {
				s3Mod = s3.New(updateServerS3Name.String())
			} else {
				s3Mod = s3.New("default")
			}
			bucketName := s3Mod.GetCfg().BucketName

			_, err = s3Mod.PutObject(obj, getFileName.String(), bucketName, ff.Size())
			if err != nil {
				return err
			}

			//上传当前文件
			url, err := s3Mod.GetFileUrl(getFileName.String(), bucketName)
			g.Log().Debugf(ctx, "下载地址:%v", url)

			client := g.Client()
			client.SetTimeout(time.Minute)
			client.SetDiscovery(nil)

			//循环服务器，推送更新
			for _, v := range list {
				address := v.Address
				g.Log().Debugf(ctx, "准备同步服务器:%v,url=%v", v.Name, address+"/callback/update")
				get, err := client.Post(ctx, address+"/callback/update", &UpdateReq{
					FileUrl: url.String(),
				})
				if err != nil {
					get, err = client.Proxy("http://127.0.0.1:10809").
						Post(ctx, address+"/callback/update", &UpdateReq{
							FileUrl: url.String(),
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
