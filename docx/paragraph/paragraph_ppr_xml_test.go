package paragraph

import (
    "fmt"
    "testing"
)

func TestPPr_GetStyleXmlBytes(t *testing.T) {
    p := PPr{}

    exp := fmt.Sprintf(`<w:pPr><w:pStyle w:val="%v"/></w:pPr>`, p.GetId())
    act, err := p.GetStyleXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestPPr_GetXmlBytes(t *testing.T) {
    p := PPr{}

    exp := `<w:pPr></w:pPr>`
    act, err := p.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    p.SetHorizontalAlignment("")
    exp = `<w:pPr><w:jc w:val=""/></w:pPr>`
    act, err = p.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    p.SetKeepLines(true)
    exp = `<w:pPr><w:jc w:val=""/><w:keeplines/></w:pPr>`
    act, err = p.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    p.SetKeepNext(true)
    exp = `<w:pPr><w:jc w:val=""/><w:keeplines/><w:keepNext/></w:pPr>`
    act, err = p.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
