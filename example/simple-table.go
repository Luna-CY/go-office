package main

import (
    "github.com/Luna-CY/go-office/docx"
    "log"
)

func main() {
    doc, err := docx.New()
    if nil != err {
        log.Fatal(err)
    }

    // 初始化创建三列
    t1 := doc.AddTableWithColumns(3)

    // 单独增加一列
    t1.AddCol()

    // 增加一行
    t1r1 := t1.AddRow()

    // 取到第一行的第一列
    // 每行的列数与col数量一致
    t1r1c1, err := t1r1.GetCell(0)
    if nil != err {
        log.Fatal(err)
    }

    // 设置单元格内容
    t1r1c1.AddParagraph().AddRun().AddText("第一行第一列")

    t1r1c2, err := t1r1.GetCell(1)
    if nil != err {
        log.Fatal(err)
    }

    t1r1c2.AddParagraph().AddRun().AddText("第一行第二列")

    // 第二行
    t1r2 := t1.AddRow()

    // 获取第二行第一列
    t1r2c1, err := t1r2.GetCell(0)
    if nil != err {
        log.Fatal(err)
    }

    t1r2c1.AddParagraph().AddRun().AddText("第二行第一列")

    if err := doc.Save("example-table.docx"); nil != err {
        log.Fatal(err)
    }

    log.Println("生成成功")
}
