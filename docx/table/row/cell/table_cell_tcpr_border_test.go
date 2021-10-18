package cell

import "testing"

func TestBorderManager_SetBorder(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.top {
		t.Fatal("验证失败")
	}
	if nil != bm.right {
		t.Fatal("验证失败")
	}
	if nil != bm.bottom {
		t.Fatal("验证失败")
	}
	if nil != bm.left {
		t.Fatal("验证失败")
	}
	if nil != bm.insideH {
		t.Fatal("验证失败")
	}
	if nil != bm.insideV {
		t.Fatal("验证失败")
	}

	bm.SetBorder("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.top {
		t.Fatal("验证失败")
	}
	if nil == bm.right {
		t.Fatal("验证失败")
	}
	if nil == bm.bottom {
		t.Fatal("验证失败")
	}
	if nil == bm.left {
		t.Fatal("验证失败")
	}
	if nil == bm.insideH {
		t.Fatal("验证失败")
	}
	if nil == bm.insideV {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetBottom(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.bottom {
		t.Fatal("验证失败")
	}

	bm.SetBottom("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.bottom {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetLeft(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.left {
		t.Fatal("验证失败")
	}

	bm.SetLeft("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.left {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetRight(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.right {
		t.Fatal("验证失败")
	}

	bm.SetRight("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.right {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetTop(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.top {
		t.Fatal("验证失败")
	}

	bm.SetTop("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.top {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetInsideHorizontal(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.top {
		t.Fatal("验证失败")
	}

	bm.SetInsideHorizontal("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.insideH {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_SetInsideVertical(t *testing.T) {
	bm := BorderManager{}
	if true == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil != bm.top {
		t.Fatal("验证失败")
	}

	bm.SetInsideVertical("", "", 0, 0, false)
	if false == bm.isSet {
		t.Fatal("验证失败")
	}
	if nil == bm.insideV {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_getBorderBody(t *testing.T) {
	bm := BorderManager{}
	bm.SetTop("", "", 0, 0, false)

	act, err := bm.getBorderBody("w:top", bm.top)
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp := `<w:top w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestBorderManager_GetXmlBytes(t *testing.T) {
	bm := BorderManager{}

	act, err := bm.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	bm.SetBorder("", "", 0, 0, false)
	act, err = bm.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp := `<w:tcBorders><w:top w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/><w:right w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/><w:bottom w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/><w:left w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/><w:insideH w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/><w:insideV w:val="" w:sz="0" w:color="" w:space="0" w:shadow="false"/></w:tcBorders>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
