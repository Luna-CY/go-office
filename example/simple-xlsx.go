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

	if err := document.Save("simple.xlsx"); nil != err {
		log.Fatalf("保存文件失败: %v\n", err)
	}
}
