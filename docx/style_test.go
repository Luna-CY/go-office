package docx

import (
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "testing"
)

func TestStyleConfig_GetDefaultParagraphProperties(t *testing.T) {
    style := StyleConfig{}

    if nil == style.GetDefaultParagraphProperties() {
        t.Fatal("验证失败")
    }

    exp := fmt.Sprintf(`%v%v<w:docDefaults></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
    act, err := style.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    style.GetDefaultParagraphProperties().SetKeepNext(true)
    exp = fmt.Sprintf(`%v%v<w:docDefaults><w:pPrDefault><w:pPr><w:keepNext/></w:pPr></w:pPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
    act, err = style.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestStyleConfig_GetDefaultRunProperties(t *testing.T) {
    style := StyleConfig{}

    if nil == style.GetDefaultRunProperties() {
        t.Fatal("验证失败")
    }

    exp := fmt.Sprintf(`%v%v<w:docDefaults></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
    act, err := style.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    style.GetDefaultRunProperties().SetSize(30)
    exp = fmt.Sprintf(`%v%v<w:docDefaults><w:rPrDefault><w:rPr><w:sz w:val="30"/></w:rPr></w:rPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
    act, err = style.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestStyleConfig_GetXmlBytes(t *testing.T) {
    style := StyleConfig{}
    if 0 != len(style.styleList) {
        t.Fatal("验证失败")
    }

    exp := fmt.Sprintf(`%v%v%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
    act, err := style.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestStyleConfig_AddParagraphStyle(t *testing.T) {
    style := StyleConfig{}
    if 0 != len(style.styleList) {
        t.Fatal("验证失败")
    }

    style.AddParagraphStyle("TEST_ID", nil, nil)
    if 1 != len(style.styleList) {
        t.Fatal("验证失败")
    }

    // TODO
}
