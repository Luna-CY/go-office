package paragraph

import (
	"fmt"
	"testing"
)

func TestParagraph_GetXmlBytes(t *testing.T) {
	p := Paragraph{}

	act, err := p.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "<w:p/>" != string(act) {
		t.Fatal("验证失败")
	}

	p.GetProperties().SetKeepLines(true)
	exp := fmt.Sprintf(`<w:p><w:pPr><w:keeplines/></w:pPr></w:p>`)
	act, err = p.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	p.GetRunProperties().SetBold(true)
	exp = fmt.Sprintf(`<w:p><w:pPr><w:keeplines/></w:pPr><w:rPr><w:b w:val="true"/></w:rPr></w:p>`)
	act, err = p.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	r := p.AddRun()
	r.AddText("TEST")
	exp = fmt.Sprintf(`<w:p><w:pPr><w:keeplines/></w:pPr><w:rPr><w:b w:val="true"/></w:rPr><w:r><w:t>TEST</w:t></w:r></w:p>`)
	act, err = p.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
