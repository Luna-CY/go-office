package main

import (
    "github.com/Luna-CY/go-office/docx"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/run"
    "log"
)

func main() {
    doc, err := docx.New()
    if nil != err {
        log.Fatal(err)
    }

    doc.GetStyle().GetDefaultRunProperties().SetSize(32)
    doc.GetStyle().GetDefaultParagraphProperties().GetSpacing().SetSpace(360)

    p1 := doc.AddParagraph()
    p1.GetParagraphProperties().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)

    r1 := p1.AddRun()
    r1.AddText("hi")
    r1.GetRunProperties().SetHighlight(run.HighlightYellow)

    r2 := p1.AddRun()
    r2.AddText("没有背景色")

    p2 := doc.AddParagraph()
    p2.GetParagraphProperties().GetIdentity().SetFirstLine(400)
    p2r1 := p2.AddRun()
    p2r1.AddText("第二段内容足够长到换行第二段内容足够长到换行第二段内容足够长到换行第二段内容足够长到换行第二段内容足够长到换行")

    p3 := doc.AddParagraph()
    p3r1 := p3.AddRun()
    p3r1.AddText("第三段，没有样式")

    if err := doc.Save("example.docx"); nil != err {
        log.Fatal(err)
    }

    log.Println("生成成功")
}
