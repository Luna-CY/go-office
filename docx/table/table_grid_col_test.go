package table

import (
	"testing"
)

func TestGridCol_GetWidth(t *testing.T) {
	grid := GridCol{}

	if 0 != grid.GetWidth() {
		t.Fatal("验证失败")
	}

	width := 30
	grid.w = &width

	if 30 != grid.GetWidth() {
		t.Fatal("验证失败")
	}
}

func TestGridCol_SetWidth(t *testing.T) {
	grid := GridCol{}

	if nil != grid.w {
		t.Fatal("验证失败")
	}

	grid.SetWidth(30)
	if nil == grid.w {
		t.Fatal("验证失败")
	}

	if 30 != *grid.w {
		t.Fatal("验证失败")
	}
}

func TestGridCol_GetXmlBytes(t *testing.T) {
	grid := GridCol{}

	act, err := grid.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `<w:gridCol/>` != string(act) {
		t.Fatal("验证失败")
	}

	grid.SetWidth(30)
	act, err = grid.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `<w:gridCol w:w="30"/>` != string(act) {
		t.Fatal("验证失败")
	}
}
