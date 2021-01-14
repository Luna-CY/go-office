package paragraph

import (
    "github.com/Luna-CY/go-office/docx/run"
    "testing"
)

func TestParagraph_GetProperties(t *testing.T) {
    p := Paragraph{}

    if nil == p.GetProperties() {
        t.Fatal("验证失败")
    }
}

func TestParagraph_GetRunProperties(t *testing.T) {
    p := Paragraph{}

    if nil == p.GetRunProperties() {
        t.Fatal("验证失败")
    }
}

func TestParagraph_GetRuns(t *testing.T) {
    p := Paragraph{}

    if 0 != len(p.GetRuns()) {
        t.Fatal("验证失败")
    }
}

func TestParagraph_AddRun(t *testing.T) {
    p := Paragraph{}
    p.AddRun()

    if 1 != len(p.GetRuns()) {
        t.Fatal("验证失败")
    }
}

func TestParagraph_AddBreakLine(t *testing.T) {
    p := Paragraph{}
    p.AddBreakLine(run.BreakLineTypeDefault, run.BreakLineClearTypeAll)

    if 1 != len(p.GetRuns()) {
        t.Fatal("验证失败")
    }
}
