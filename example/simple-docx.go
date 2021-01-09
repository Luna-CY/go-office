package main

import (
    "github.com/Luna-CY/go-office/docx"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/section"
    "log"
)

func main() {
    doc, err := docx.New()
    if nil != err {
        log.Fatal(err)
    }

    doc.GetSection().GetPageBorder().SetDisplay(section.PageBorderDisplayAll)
    doc.GetSection().GetPageBorder().SetTop(section.PageBorderStyleSingle, "000000", 10, 24, false)

    p1 := doc.AddParagraph()
    p1.GetPPr().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)

    r1 := p1.AddRun()
    r1.SetText("hi")

    if err := doc.Save("example.docx"); nil != err {
        log.Fatal(err)
    }

    log.Println("生成成功")
}
