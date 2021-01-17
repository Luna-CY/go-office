package run

import (
    "fmt"
    "testing"
)

func TestRun_GetXmlBytes(t *testing.T) {
    r := Run{}

    act, err := r.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if "<w:r></w:r>" != string(act) {
        t.Fatal("验证失败")
    }

    r.AddText("HI")
    act, err = r.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if "<w:r><w:t>HI</w:t></w:r>" != string(act) {
        t.Fatal("验证失败")
    }

    r.GetProperties().SetBold(true)
    act, err = r.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp := fmt.Sprintf(`<w:r><w:rPr><w:rStyle w:val="%p"/></w:rPr><w:t>HI</w:t></w:r>`, r.GetProperties())

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
