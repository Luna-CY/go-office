package cell

import "testing"

func TestMargin_GetBottom(t *testing.T) {
	m := Margin{}
	if nil != m.bottom {
		t.Fatal("验证失败")
	}

	if 0 != m.GetBottom() {
		t.Fatal("验证失败")
	}
}

func TestMargin_GetLeft(t *testing.T) {
	m := Margin{}
	if nil != m.left {
		t.Fatal("验证失败")
	}

	if 0 != m.GetLeft() {
		t.Fatal("验证失败")
	}
}

func TestMargin_GetRight(t *testing.T) {
	m := Margin{}
	if nil != m.right {
		t.Fatal("验证失败")
	}

	if 0 != m.GetRight() {
		t.Fatal("验证失败")
	}
}

func TestMargin_GetTop(t *testing.T) {
	m := Margin{}
	if nil != m.top {
		t.Fatal("验证失败")
	}

	if 0 != m.GetTop() {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetBottom(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.bottom {
		t.Fatal("验证失败")
	}

	m.SetBottom(10)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.bottom {
		t.Fatal("验证失败")
	}
	if 10 != *m.bottom {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetLeft(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.left {
		t.Fatal("验证失败")
	}

	m.SetLeft(10)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.left {
		t.Fatal("验证失败")
	}
	if 10 != *m.left {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetRight(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.right {
		t.Fatal("验证失败")
	}

	m.SetRight(10)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.right {
		t.Fatal("验证失败")
	}
	if 10 != *m.right {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetTop(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.top {
		t.Fatal("验证失败")
	}

	m.SetTop(10)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.top {
		t.Fatal("验证失败")
	}
	if 10 != *m.top {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetMargin(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.bottom {
		t.Fatal("验证失败")
	}
	if nil != m.right {
		t.Fatal("验证失败")
	}
	if nil != m.left {
		t.Fatal("验证失败")
	}
	if nil != m.top {
		t.Fatal("验证失败")
	}

	m.SetMargin(10)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.bottom {
		t.Fatal("验证失败")
	}
	if 10 != *m.bottom {
		t.Fatal("验证失败")
	}
	if nil == m.right {
		t.Fatal("验证失败")
	}
	if 10 != *m.right {
		t.Fatal("验证失败")
	}
	if nil == m.left {
		t.Fatal("验证失败")
	}
	if 10 != *m.left {
		t.Fatal("验证失败")
	}
	if nil == m.top {
		t.Fatal("验证失败")
	}
	if 10 != *m.top {
		t.Fatal("验证失败")
	}
}

func TestMargin_SetGroup(t *testing.T) {
	m := Margin{}
	if true == m.isSet {
		t.Fatal("验证失败")
	}
	if nil != m.bottom {
		t.Fatal("验证失败")
	}
	if nil != m.right {
		t.Fatal("验证失败")
	}
	if nil != m.left {
		t.Fatal("验证失败")
	}
	if nil != m.top {
		t.Fatal("验证失败")
	}

	m.SetGroup(10, 30)
	if false == m.isSet {
		t.Fatal("验证失败")
	}
	if nil == m.bottom {
		t.Fatal("验证失败")
	}
	if 10 != *m.bottom {
		t.Fatal("验证失败")
	}
	if nil == m.right {
		t.Fatal("验证失败")
	}
	if 30 != *m.right {
		t.Fatal("验证失败")
	}
	if nil == m.left {
		t.Fatal("验证失败")
	}
	if 30 != *m.left {
		t.Fatal("验证失败")
	}
	if nil == m.top {
		t.Fatal("验证失败")
	}
	if 10 != *m.top {
		t.Fatal("验证失败")
	}
}

func TestMargin_GetXmlBytes(t *testing.T) {
	m := Margin{}

	act, err := m.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `` != string(act) {
		t.Fatal("验证失败")
	}

	m.SetMargin(10)

	exp := `<w:tcMar><w:top w:w="10" w:type="dxa"/><w:right w:w="10" w:type="dxa"/><w:bottom w:w="10" w:type="dxa"/><w:left w:w="10" w:type="dxa"/></w:tcMar>`
	act, err = m.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
