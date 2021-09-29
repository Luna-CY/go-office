package table

import (
	"testing"
)

func TestRow_GetXmlBytes(t *testing.T) {
	row := Row{}

	act, err := row.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `<w:tr></w:tr>` != string(act) {
		t.Fatal("验证失败")
	}

	row.addCell()
	act, err = row.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp := `<w:tr><w:tc><w:tcPr><w:tcW w:w="0" w:type="auto"/></w:tcPr><w:p/></w:tc></w:tr>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
