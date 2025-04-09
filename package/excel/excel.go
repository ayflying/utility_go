package excel

import (
	"context"
	"github.com/ayflying/excel2json"
	"github.com/goccy/go-json"
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

var (
	shadiao = []string{",", ":"}
)

type FileItem struct {
	Name     string            `json:"name" dc:"配置文件名"`
	Filename string            `json:"filename" dc:"文件名"`
	Tabs     []string          `json:"tabs" dc:"页签"`
	Items    []string          `json:"items" dc:"道具字段"`
	ItemsMap []string          `json:"items_map" dc:"道具字段map格式"`
	Slice    map[string]string `json:"slice" dc:"切片"`
	Json     []string          `json:"json" dc:"json"`
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

		//排除注释行
		list = s.RemoveComments(list, fileItem.Json)

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
		//json格式转换
		if len(fileItem.Json) > 0 {
			list = s.jsonFormat(list, fileItem.Json)
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

//排除配置中注释行
func (s *Excel) RemoveComments(list []interface{}, json []string) []interface{} {
	var temp = make([]interface{}, 0)
	// 遍历列表中的每个元素
	for _, v2 := range list {
		var add = true
		// 遍历当前元素的每个键值对
		for _, v3 := range v2.(g.Map) {
			// 如果字符串中存在//则跳过不写入temp
			//if gstr.Contains(gconv.String(v3), "//") {
			//	//delKey = append(delKey, k2)
			//	add = false
			//	break
			//}
			if strings.HasPrefix(gconv.String(v3), "//") {
				add = false
				break
			}
		}
		if add {
			temp = append(temp, v2)
		}

	}
	return temp
}

// itemsFormat 格式化列表中的道具字段
// 参数:
//   - list: 待处理的列表，包含多个元素，每个元素是一个 g.Map 类型
//   - Items: 包含需要处理的道具字段名称的切片
// 返回值:
//   - 处理后的列表
func (s *Excel) itemsFormat(list []interface{}, Items []string) []interface{} {
	// 遍历列表中的每个元素
	for k2, v2 := range list {
		// 遍历当前元素的每个键值对
		for k3, v3 := range v2.(g.Map) {
			// 检查当前键是否在需要处理的道具字段列表中
			if gstr.InArray(Items, k3) {
				// 检查当前值是否为字符串类型
				if _, ok := v3.(string); ok {
					// 如果是字符串类型，调用 Spilt2Item 函数将其转换为 [][]int64 类型，并更新到列表中
					list[k2].(g.Map)[k3] = Spilt2Item(v3.(string))
				} else {
					// 如果不是字符串类型，记录错误日志
					g.Log().Errorf(gctx.New(), "当前类型断言失败:%v,list=%v", v3, v2)
				}
			}
		}
	}
	// 返回处理后的列表
	return list
}

// itemsMapFormat 将列表中指定字段的道具信息转换为映射格式
// 参数:
//   - list: 待处理的列表，包含多个元素，每个元素是一个 g.Map 类型
//   - Items: 包含需要处理的道具字段名称的切片
// 返回值:
//   - 处理后的列表
func (s *Excel) itemsMapFormat(list []interface{}, Items []string) []interface{} {
	// 遍历列表中的每个元素
	for k2, v2 := range list {
		// 遍历当前元素的每个键值对
		for k3, v3 := range v2.(g.Map) {
			// 检查当前键是否在需要处理的道具字段列表中
			if gstr.InArray(Items, k3) {
				// 检查当前值是否为字符串类型
				if _, ok := v3.(string); ok {
					// 如果是字符串类型，调用 Spilt2Item 函数将其转换为 [][]int64 类型
					get := Spilt2Item(v3.(string))
					// 调用 Items2Map 函数将 [][]int64 类型的数据转换为映射格式，并更新到列表中
					list[k2].(g.Map)[k3] = s.Items2Map(get)
				} else {
					// 如果不是字符串类型，记录错误日志
					g.Log().Errorf(gctx.New(), "当前类型断言失败:%v,list=%v", v3, v2)
				}
			}
		}
	}
	// 返回处理后的列表
	return list
}

// sliceFormat 格式化列表中指定字段为切片格式
// 参数:
//   - list: 待处理的列表，包含多个元素，每个元素是一个 g.Map 类型
//   - Slice: 一个映射，键为需要处理的字段名，值为目标类型（如 "int", "int64", "float64"）
// 返回值:
//   - 处理后的列表
func (s *Excel) sliceFormat(list []interface{}, Slice map[string]string) []interface{} {
	// 遍历 Slice 映射中的每个键值对
	for s1, s2 := range Slice {
		// 遍历列表中的每个元素
		for k2, v2 := range list {
			// 遍历当前元素的每个键值对
			for k3, v3 := range v2.(g.Map) {
				// 判断当前键是否与 Slice 中的键匹配
				if s1 != k3 {
					// 不匹配则跳过当前循环
					continue
				}
				// 检查当前值是否为空字符串
				if gconv.String(v3) == "" {
					// 若为空，则将该字段设置为空字符串切片
					list[k2].(g.Map)[k3] = []string{}
					// 跳过当前循环
					continue
				}
				// 用于存储分割后的字符串切片
				var parts []string
				// 断言当前值是否为字符串类型
				if get, ok := v3.(string); !ok {
					// 若断言失败，将当前值转换为字符串并作为唯一元素存入 parts
					parts = []string{gconv.String(v3)}
				} else {
					// 若为字符串类型，将字符串中的特殊字符替换为 "|"
					for _, v := range shadiao {
						get = strings.ReplaceAll(get, v, "|")
					}
					// 按 "|" 分割字符串
					parts = strings.Split(get, "|")
				}

				// 根据 Slice 映射中的值进行类型转换
				switch s2 {
				case "int":
					// 创建一个长度为 parts 的 int 切片
					var temp = make([]int, len(parts))
					// 遍历 parts 切片，将每个元素转换为 int 类型
					for k, v := range parts {
						temp[k], _ = strconv.Atoi(v)
					}
					// 将转换后的切片存入列表中
					list[k2].(g.Map)[k3] = temp
				case "int64":
					// 创建一个长度为 parts 的 int64 切片
					var temp = make([]int64, len(parts))
					// 遍历 parts 切片，将每个元素转换为 int64 类型
					for k, v := range parts {
						temp[k], _ = strconv.ParseInt(v, 10, 64)
					}
				case "float64":
					// 创建一个长度为 parts 的 float64 切片
					var temp = make([]float64, len(parts))
					// 遍历 parts 切片，将每个元素转换为 float64 类型
					for k, v := range parts {
						temp[k], _ = strconv.ParseFloat(v, 64)
					}
					// 将转换后的切片存入列表中
					list[k2].(g.Map)[k3] = temp
				default:
					// 若未匹配到指定类型，直接将 parts 存入列表中
					list[k2].(g.Map)[k3] = parts
				}
			}
		}
	}
	// 返回处理后的列表
	return list
}

// jsonFormat 将列表中指定字段的 JSON 字符串解析为 Go 数据结构
// 参数:
//   - list: 待处理的列表，包含多个元素，每个元素是一个 g.Map 类型
//   - Items: 包含需要处理的 JSON 字段名称的切片
// 返回值:
//   - 处理后的列表
func (s *Excel) jsonFormat(list []interface{}, Items []string) []interface{} {
	// 遍历列表中的每个元素
	for k2, v2 := range list {
		// 遍历当前元素的每个键值对
		for k3, v3 := range v2.(g.Map) {
			// 检查当前键是否在需要处理的 JSON 字段列表中
			if gstr.InArray(Items, k3) {
				// 检查当前值是否为字符串类型
				if _, ok := v3.(string); ok {
					// 用于存储解析后的 JSON 数据
					var get interface{}
					// 将字符串解析为 JSON 数据
					json.Unmarshal([]byte(v3.(string)), &get)
					// 将解析后的 JSON 数据更新到列表中
					list[k2].(g.Map)[k3] = get
				} else {
					// 如果不是字符串类型，记录错误日志
					g.Log().Errorf(gctx.New(), "当前类型断言失败:%v,list=%v", v3, v2)
				}
			}
		}
	}
	// 返回处理后的列表
	return list
}
