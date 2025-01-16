package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Make = gcmd.Command{
		Name:  "make",
		Usage: "make",
		Brief: "创建模块文件",
		Arguments: []gcmd.Argument{
			{Name: "model", Short: "m", Brief: "模块名"},
			{Name: "id", Short: "i", Brief: "活动id"},
			{Name: "name", Short: "n", Brief: "服务文件名"},
		},
		Examples: "make -m=act -id=1:     创建活动1的接口与服务文件 \n" +
			"make -m=logic -n=test: 创建test的服务文件",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			//g.Dump(parser.GetOptAll(), parser.GetArgAll())
			//return
			var model = parser.GetOpt("model").String()
			//var name = parser.GetOpt("n").String()
			this := cMake{}
			switch model {
			case "act":
				var id = parser.GetOpt("id").Int()
				if id == 0 {
					return
				}
				err = this.Act(id)
			case "logic":
				var name = parser.GetOpt("name").String()
				if name == "" {
					return
				}
				err = this.Logic(name)
			}

			return
		},
	}
)

type cMake struct{}

func (c *cMake) Api() {

}

func (c *cMake) Act(id int) (err error) {
	filePath := fmt.Sprintf("api/act/v1/act%v.go", id)
	err = gfile.PutContents(filePath, "package v1\n")

	filePath = fmt.Sprintf("internal/game/act/act%d/act%d.go", id, id)
	fileStr := gfile.GetContents("internal/cmd/make/act")
	fileStr = gstr.Replace(fileStr, "{id}", gconv.String(id))
	err = gfile.PutContents(filePath, fileStr)

	return
}

func (c *cMake) Logic(name string) (err error) {
	var filePath = fmt.Sprintf("internal/logic/%s/%s.go", name, name)
	fileStr := gfile.GetContents("internal/cmd/make/logic")
	fileStr = gstr.Replace(fileStr, "{package}", name)
	fileStr = gstr.Replace(fileStr, "{name}", gstr.CaseCamel(name))

	err = gfile.PutContents(filePath, fileStr)
	return
}
