package table

import (
	"testing"
)

func TestBackground_GetBackgroundColor(t *testing.T) {
	b := Background{}

	if "" != b.GetBackgroundColor() {
		t.Fatal("验证失败")
	}

	color := "000000"
	b = Background{fill: &color}

	if "000000" != b.GetBackgroundColor() {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetBackgroundColor(t *testing.T) {
	b := Background{}
	if true == b.isSet {
		t.Fatal("验证失败")
	}

	if nil != b.fill {
		t.Fatal("验证失败")
	}

	b.SetBackgroundColor("000000")
	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if nil == b.fill {
		t.Fatal("验证失败")
	}

	if "000000" != *b.fill {
		t.Fatal("验证失败")
	}
}

func TestBackground_GetColor(t *testing.T) {
	b := Background{}

	if "auto" != b.GetColor() {
		t.Fatal("验证失败")
	}

	color := "FFFFFF"
	b = Background{color: &color}

	if "FFFFFF" != b.GetColor() {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetColor(t *testing.T) {
	b := Background{}

	if true == b.isSet {
		t.Fatal("验证失败")
	}

	if nil != b.color {
		t.Fatal("验证失败")
	}

	b.SetColor("FFFFFF")
	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if nil == b.color {
		t.Fatal("验证失败")
	}

	if "FFFFFF" != *b.color {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetMask(t *testing.T) {
	b := Background{}

	if true == b.isSet {
		t.Fatal("验证失败")
	}

	if nil != b.mask {
		t.Fatal("验证失败")
	}

	b.SetMask("FF0000")
	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if nil == b.mask {
		t.Fatal("验证失败")
	}

	if "FF0000" != *b.mask {
		t.Fatal("验证失败")
	}
}

func TestBackground_GetXmlBytes(t *testing.T) {
	b := Background{}

	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	b = Background{}
	b.SetBackgroundColor("000000")
	exp := `<w:shd w:fill="000000" w:color="auto"/>`
	act, err = b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	b.SetColor("FF0000")
	exp = `<w:shd w:fill="000000" w:color="FF0000"/>`
	act, err = b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	b.SetMask("FFFFFF")
	exp = `<w:shd w:fill="000000" w:color="FF0000" w:val="FFFFFF"/>`
	act, err = b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
