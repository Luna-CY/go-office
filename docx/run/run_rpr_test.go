package run

import "testing"

func TestRPr_GetBackground(t *testing.T) {
	r := RPr{}

	if nil == r.GetBackground() {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetBold(t *testing.T) {
	r := RPr{}

	if false != r.bold {
		t.Fatal("验证失败")
	}

	r.SetBold(true)
	if true != r.bold {
		t.Fatal("验证失败")
	}

	r.SetBold(false)
}

func TestRPr_SetColor(t *testing.T) {
	r := RPr{}

	if nil != r.color {
		t.Fatal("验证失败")
	}

	r.SetColor("000000")
	if nil == r.color {
		t.Fatal("验证失败")
	}

	if "000000" != *r.color {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetDeleteLine(t *testing.T) {
	r := RPr{}

	if nil != r.deleteLine {
		t.Fatal("验证失败")
	}

	r.SetDeleteLine(DeleteLineStrike)
	if nil == r.deleteLine {
		t.Fatal("验证失败")
	}

	if DeleteLineStrike != *r.deleteLine {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetEmboss(t *testing.T) {
	r := RPr{}

	if false != r.emboss {
		t.Fatal("验证失败")
	}

	r.imprint = true
	r.shadow = true
	r.color = nil

	r.SetEmboss(true)
	if true != r.emboss {
		t.Fatal("验证失败")
	}
	if false != r.imprint {
		t.Fatal("验证失败")
	}
	if false != r.shadow {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetImprint(t *testing.T) {
	r := RPr{}

	if false != r.imprint {
		t.Fatal("验证失败")
	}

	r.emboss = true
	r.shadow = true

	r.SetImprint(true)
	if true != r.imprint {
		t.Fatal("验证失败")
	}
	if false != r.emboss {
		t.Fatal("验证失败")
	}
	if false != r.shadow {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetItalics(t *testing.T) {
	r := RPr{}

	if false != r.italics {
		t.Fatal("验证失败")
	}

	r.SetItalics(true)
	if true != r.italics {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetShadow(t *testing.T) {
	r := RPr{}

	if false != r.shadow {
		t.Fatal("验证失败")
	}

	r.imprint = true
	r.emboss = true

	r.SetShadow(true)
	if true != r.shadow {
		t.Fatal("验证失败")
	}
	if false != r.imprint {
		t.Fatal("验证失败")
	}
	if false != r.emboss {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetSize(t *testing.T) {
	r := RPr{}

	if nil != r.size {
		t.Fatal("验证失败")
	}

	r.SetSize(30)
	if nil == r.size {
		t.Fatal("验证失败")
	}

	if 30 != *r.size {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetUnderline(t *testing.T) {
	r := RPr{}

	if nil != r.underline {
		t.Fatal("验证失败")
	}

	r.SetUnderline(UnderlineSingle)
	if nil == r.underline {
		t.Fatal("验证失败")
	}

	if UnderlineSingle != r.underline.lineType {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetUnderlineWithColor(t *testing.T) {
	r := RPr{}

	if nil != r.underline {
		t.Fatal("验证失败")
	}

	r.SetUnderlineWithColor(UnderlineSingle, "000000")
	if nil == r.underline {
		t.Fatal("验证失败")
	}
	if nil == r.underline.color {
		t.Fatal("验证失败")
	}

	if UnderlineSingle != r.underline.lineType {
		t.Fatal("验证失败")
	}
	if "000000" != *r.underline.color {
		t.Fatal("验证失败")
	}
}

func TestRPr_SetVanish(t *testing.T) {
	r := RPr{}

	if false != r.vanish {
		t.Fatal("验证失败")
	}

	r.SetVanish(true)
	if true != r.vanish {
		t.Fatal("验证失败")
	}
}
