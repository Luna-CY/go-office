package run

import (
    "fmt"
    "testing"
)

func TestRPr_GetInnerXmlBytes(t *testing.T) {
    r := RPr{}

    act, err := r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp := ``
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetBold(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetItalics(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetEmboss(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:emboss w:val="true"/><w:color w:val="FFF"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetImprint(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:imprint w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetShadow(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetVanish(true)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetColor("000000")
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetUnderlineWithColor(UnderlineSingle, "000000")
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetDeleteLine(DeleteLineStrike)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:strike w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetDeleteLine(DeleteLineDoubleStrike)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:dstrike w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetSize(30)
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:dstrike w:val="true"/><w:sz w:val="30"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.GetBackground().SetBackgroundColor("FF0000")
    act, err = r.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:rStyle w:val="%p"/><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:dstrike w:val="true"/><w:sz w:val="30"/></w:rPr>`, &r)
    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestRPr_GetExtraXmlBytes(t *testing.T) {
    r := RPr{}

    r.GetBackground().SetBackgroundColor("FF0000")
    act, err := r.GetExtraXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp := fmt.Sprintf(`<w:rPr><w:shd w:fill="FF0000" w:color="auto"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestRPr_GetDefaultXmlBytes(t *testing.T) {
    r := RPr{}

    act, err := r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp := fmt.Sprintf(`<w:rPr></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetBold(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetItalics(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetEmboss(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:emboss w:val="true"/><w:color w:val="FFF"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetImprint(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:imprint w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetShadow(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetVanish(true)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetColor("000000")
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetUnderlineWithColor(UnderlineSingle, "000000")
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetDeleteLine(DeleteLineStrike)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:strike w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetDeleteLine(DeleteLineDoubleStrike)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:dstrike w:val="true"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }

    r.SetSize(30)
    act, err = r.GetDefaultXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp = fmt.Sprintf(`<w:rPr><w:b w:val="true"/><w:i w:val="true"/><w:shadow w:val="true"/><vanish/><w:color w:val="000000"/><w:u w:color="000000" w:val="single"/><w:dstrike w:val="true"/><w:sz w:val="30"/></w:rPr>`)
    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
