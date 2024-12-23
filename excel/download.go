package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xuri/excelize/v2"
)

func Data2Excel(ctx context.Context, data []map[string]interface{}) (file string) {

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	var sheetName = "Sheet1"
	f.NewSheet(sheetName)

	//// 准备数据
	//data = []map[string]interface{}{
	//	{"Name": "Alice", "Age": 30, "City": "New York"},
	//	{"Name": "Bob", "Age": 25, "City": "Los Angeles"},
	//	{"Name": "Charlie", "Age": 35, "City": "Chicago"},
	//}

	//写入头部
	var colIndex = 0
	var headers []string
	for header := range data[0] {
		//追加头部名字
		headers = append(headers, header)
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, 1) // 表头在第一行
		f.SetCellValue(sheetName, cell, header)
		colIndex++
	}

	// 写入数据
	for rowIndex, record := range data {
		for colIndex, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2) // 数据从第二行开始
			f.SetCellValue(sheetName, cell, record[header])                   // 通过表头获取数据
		}
	}

	// 保存 Excel 文件
	saveName := fmt.Sprintf("runtime/uploads/out_%v.xlsx", gtime.Now().Nanosecond())
	if err := f.SaveAs(saveName); err != nil {
		g.Log().Fatal(ctx, err)
	}

	//下载excel
	//g.RequestFromCtx(ctx).Response.ServeFileDownload(saveName)
	return saveName
}
