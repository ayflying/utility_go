package excel

import (
	"context"
	"github.com/ayflying/excel2json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileItem struct {
	Name     string            `json:"name" dc:"配置文件名"`
	Filename string            `json:"filename" dc:"文件名"`
	Tabs     []string          `json:"tabs" dc:"页签"`
	Items    []string          `json:"items" dc:"道具字段"`
	ItemsMap []string          `json:"items_map" dc:"道具字段map格式"`
	Slice    map[string]string `json:"slice" dc:"切片"`
}

type Excel struct {
	Header int //表头行数
	Key    int //key列
}

func New(header int, key int) *Excel {
	return &Excel{
		Header: header,
		Key:    key,
	}
}

func (s *Excel) ExcelLoad(ctx context.Context, fileItem *FileItem, mainPath string) (runTime time.Duration) {
	startTime := gtime.Now()
	filepath := path.Join("manifest/game", fileItem.Name)

	//如果filepath文件不存在，跳过
	if !gfile.Exists(path.Join(mainPath, fileItem.Filename)) {
		return
	}

	//假设我们有一个命令行工具，比如：dir（Windows环境下列出目录内容）
	var tempJson []interface{}
	for k, v2 := range fileItem.Tabs {
		sheet := v2
		if k == 0 {
			sheet = v2
		}

		//导出json
		excel2json.Excel(path.Join(mainPath, fileItem.Filename),
			filepath, s.Header, s.Key, sheet)

		//如果配置了道具字段,则进行转换
		//g.Log().Info(ctx, "当前任务表=%v,items=%v", v.Name, v.Items)
		fileBytes := gfile.GetBytes(filepath)
		arr, _ := gjson.DecodeToJson(fileBytes)
		list := arr.Array()
		//格式化item格式
		if len(fileItem.Items) > 0 {
			list = s.itemsFormat(list, fileItem.Items)
		}
		if len(fileItem.ItemsMap) > 0 {
			list = s.itemsMapFormat(list, gconv.Strings(fileItem.ItemsMap))
		}
		//格式化切片修改
		if len(fileItem.Slice) > 0 {
			list = s.sliceFormat(list, fileItem.Slice)
		}

		//拼接json
		tempJson = append(tempJson, list...)
		fileBytes, _ = gjson.MarshalIndent(tempJson, "", "\t")
		err := gfile.PutBytes(filepath, fileBytes)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}

	runTime = gtime.Now().Sub(startTime)
	return
}

func (s *Excel) itemsFormat(list []interface{}, Items []string) []interface{} {
	for k2, v2 := range list {
		for k3, v3 := range v2.(g.Map) {
			if gstr.InArray(Items, k3) {
				if _, ok := v3.(string); ok {
					list[k2].(g.Map)[k3] = s.Spilt2Item(v3.(string))

				} else {
					g.Log().Errorf(gctx.New(), "当前类型断言失败:%v,list=%v", v3, v2)
				}
			}
		}
	}
	return list
}

func (s *Excel) itemsMapFormat(list []interface{}, Items []string) []interface{} {
	for k2, v2 := range list {
		for k3, v3 := range v2.(g.Map) {
			if gstr.InArray(Items, k3) {
				if _, ok := v3.(string); ok {
					get := s.Spilt2Item(v3.(string))
					list[k2].(g.Map)[k3] = s.Items2Map(get)
				} else {
					g.Log().Errorf(gctx.New(), "当前类型断言失败:%v,list=%v", v3, v2)
				}
			}
		}
	}
	return list
}

func (s *Excel) sliceFormat(list []interface{}, Slice map[string]string) []interface{} {
	for s1, s2 := range Slice {
		for k2, v2 := range list {
			for k3, v3 := range v2.(g.Map) {
				//判断是否存在
				if s1 != k3 {
					continue
				}
				if gconv.String(v3) == "" {
					list[k2].(g.Map)[k3] = []string{}
					continue
				}

				var parts []string
				//断言是否成功
				if get, ok := v3.(string); !ok {
					//g.Log().Errorf(gctx.New(), "当前类型断言失败:%v", v3)
					parts = []string{gconv.String(v3)}
					continue
				} else {
					parts = strings.Split(get, "|") // 分割字符串
				}

				switch s2 {
				case "int":
					var temp = make([]int, len(parts))
					for k, v := range parts {
						temp[k], _ = strconv.Atoi(v)
					}
					list[k2].(g.Map)[k3] = temp
				case "int64":
					var temp = make([]int64, len(parts))
					for k, v := range parts {
						temp[k], _ = strconv.ParseInt(v, 10, 64)
					}
				case "float64":
					var temp = make([]float64, len(parts))
					for k, v := range parts {
						temp[k], _ = strconv.ParseFloat(v, 64)
					}
					list[k2].(g.Map)[k3] = temp
				default:
					list[k2].(g.Map)[k3] = parts
				}

			}

		}
	}

	return list
}
