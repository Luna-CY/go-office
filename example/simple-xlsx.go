package main

import (
	"github.com/Luna-CY/go-office/xlsx"
	"log"
)

func main() {
	document, err := xlsx.New()
	if nil != err {
		log.Fatalf("无法创建表格: %v\n", err)
	}

	sheet := document.NewSheetWithName("test-new-sheet")
	sheet.SetCellWithValue(1, 1, "hahahah")
	sheet.SetCellWithValue(2, 1, "hahahah")
	sheet.SetCellWithNumber(2, 2, 100)
	sheet.SetCellWithNumber(2, 3, 12.34)

	if err := document.Save("simple.xlsx"); nil != err {
		log.Fatalf("保存文件失败: %v\n", err)
	}
}
