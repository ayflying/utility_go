package excel

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"
)

// Excel2Slice 读取excel文件导入为切片
func Excel2Slice(filePath string, _sheet ...string) [][]string {
	excelObj, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("无法打开Excel文件: %v", err)
	}
	defer excelObj.Close()
	var sheet string
	if len(_sheet) == 0 {
		sheet = excelObj.GetSheetList()[0]
	} else {
		sheet = _sheet[0]
	}
	res, err := excelObj.GetRows(sheet)

	return res
}

// 字符串转道具类型
func (s *Excel) Spilt2Item(str string) (result [][]int64) {
	var shadiao = []string{","}
	for _, v := range shadiao {
		str = strings.ReplaceAll(str, v, "|")
		//parts = append(parts, strings.Split(str, v)...) // 分割字符串
	}

	//var parts []string
	parts := strings.Split(str, "|") // 分割字符串
	if parts == nil {
		parts = []string{str}
	}

	//var parts []string
	//for _, v := range parts1 {
	//	parts = append(parts, strings.Split(v, ",")...) // 分割字符串
	//}

	//for _, v := range parts1 {
	//	parts2 := strings.Split(v, ",") // 分割字符串
	//	if parts2 == nil {
	//		parts = append(parts, v)
	//	} else {
	//		parts = append(parts, parts2...)
	//	}
	//}

	for i := 0; i < len(parts)-1; i += 2 {
		num1, _ := strconv.ParseInt(parts[i], 10, 64)
		num2, _ := strconv.ParseInt(parts[i+1], 10, 64)

		pair := []int64{num1, num2}
		result = append(result, pair)
	}
	return
}

// 道具格式转map
func (s *Excel) Items2Map(items [][]int64) (list map[int64]int64) {
	list = make(map[int64]int64)
	for _, v := range items {
		list[v[0]] = v[1]
	}
	return
}
