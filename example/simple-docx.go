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

    doc.GetSection().GetPageNumber().SetPageNumber(section.PageNumberFmtLowerLetter).SetStart(10)

    p1 := doc.AddParagraph()
    p1.GetPPr().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)
    p1.GetPPr().GetBackground().SetBackgroundColor("FF0000")

    r1 := p1.AddRun()
    r1.AddText("hi")

    if err := doc.Save("example.docx"); nil != err {
        log.Fatal(err)
    }

    log.Println("生成成功")
}
