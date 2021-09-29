package paragraph

import "testing"

func TestIdentity_SetFirstLine(t *testing.T) {
	i := Identity{}
	if true == i.isSet {
		t.Fatal("验证失败")
	}

	if nil != i.firstLine {
		t.Fatal("验证失败")
	}

	i.SetFirstLine(30)
	if false == i.isSet {
		t.Fatal("验证失败")
	}
	if nil == i.firstLine {
		t.Fatal("验证失败")
	}
	if 30 != *i.firstLine {
		t.Fatal("验证失败")
	}
}

func TestIdentity_SetHanging(t *testing.T) {
	i := Identity{}
	if true == i.isSet {
		t.Fatal("验证失败")
	}

	if nil != i.hanging {
		t.Fatal("验证失败")
	}

	i.SetHanging(30)
	if false == i.isSet {
		t.Fatal("验证失败")
	}
	if nil == i.hanging {
		t.Fatal("验证失败")
	}
	if 30 != *i.hanging {
		t.Fatal("验证失败")
	}
}

func TestIdentity_SetLeft(t *testing.T) {
	i := Identity{}
	if true == i.isSet {
		t.Fatal("验证失败")
	}

	if nil != i.left {
		t.Fatal("验证失败")
	}

	i.SetLeft(30)
	if false == i.isSet {
		t.Fatal("验证失败")
	}
	if nil == i.left {
		t.Fatal("验证失败")
	}
	if 30 != *i.left {
		t.Fatal("验证失败")
	}
}

func TestIdentity_SetRight(t *testing.T) {
	i := Identity{}
	if true == i.isSet {
		t.Fatal("验证失败")
	}

	if nil != i.right {
		t.Fatal("验证失败")
	}

	i.SetRight(30)
	if false == i.isSet {
		t.Fatal("验证失败")
	}
	if nil == i.right {
		t.Fatal("验证失败")
	}
	if 30 != *i.right {
		t.Fatal("验证失败")
	}
}

func TestIdentity_GetXmlBytes(t *testing.T) {
	i := Identity{}

	act, err := i.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	i.SetLeft(30)
	act, err = i.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp := `<w:ind w:left="30" w:start="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	i = Identity{}
	i.SetRight(30)
	act, err = i.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:ind w:right="30" w:end="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	i = Identity{}
	i.SetHanging(30)
	act, err = i.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:ind w:hanging="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	i = Identity{}
	i.SetFirstLine(30)
	act, err = i.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:ind w:firstLine="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
