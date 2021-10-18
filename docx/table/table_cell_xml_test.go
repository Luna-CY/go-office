package table

import (
	"testing"
)

func TestCell_GetXmlBytes(t *testing.T) {
	c := Cell{}

	act, err := c.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `<w:tc><w:tcPr><w:tcW w:w="0" w:type="auto"/></w:tcPr><w:p/></w:tc>` != string(act) {
		t.Fatal("验证失败")
	}
}
