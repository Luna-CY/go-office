package main

import (
    "github.com/Luna-CY/go-office/docx"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/run"
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
    p1.GetParagraphProperties().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)
    p1.GetParagraphProperties().GetBackground().SetBackgroundColor("000000")
    p1.GetRunProperties().SetColor("FFFFFF").SetSize(300)

    r1 := p1.AddRun()
    r1.AddText("hi")
    r1.GetRunProperties().SetHighlight(run.HighlightYellow)

    r2 := p1.AddRun()
    r2.AddText("没有背景色")

    if err := doc.Save("example.docx"); nil != err {
        log.Fatal(err)
    }

    log.Println("生成成功")
}
