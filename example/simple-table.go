package main

import (
	"github.com/Luna-CY/go-office/docx"
	"github.com/Luna-CY/go-office/docx/paragraph"
	"github.com/Luna-CY/go-office/docx/table"
	"log"
)

func main() {
	doc, err := docx.New()
	if nil != err {
		log.Fatal(err)
	}

	hdr1 := doc.NewHeader()
	hp1 := hdr1.AddParagraph()
	hp1.GetProperties().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)
	hp1.AddRun().AddText("this is header")
	doc.UseHeader(docx.HeaderTypeDefault, hdr1)

	ftr1 := doc.NewFooter()
	fp1 := ftr1.AddParagraph()
	fp1.GetProperties().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)
	fp1.AddRun().AddText("this is footer")
	doc.UseFooter(docx.FooterTypeDefault, ftr1)

	// 初始化创建三列
	t1 := doc.AddTableWithColumns(6)
	t1.GetProperties().GetCellMargin().SetMargin(300)
	t1.GetProperties().GetBorder().SetBorder(table.BorderStyleSingle, "000000", 10, 0, false)

	// 单独增加一列
	//t1.AddColWithWidth(2000)

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

	// 设置整行单元格
	if err := t1r2.AddCellText("第二行第一列", "第二行第二列"); nil != err {
		log.Fatalf("设置行内容失败: %v\n", err)
	}

	if err := doc.Save("example-table.docx"); nil != err {
		log.Fatal(err)
	}

	log.Println("生成成功")
}
