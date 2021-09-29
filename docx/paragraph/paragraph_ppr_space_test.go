package paragraph

import "testing"

func TestSpacing_SetBefore(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if nil != s.before {
		t.Fatal("验证失败")
	}

	s.SetBefore(30)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if nil == s.before {
		t.Fatal("验证失败")
	}

	if 30 != *s.before {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetAfter(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if nil != s.after {
		t.Fatal("验证失败")
	}

	s.SetAfter(30)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if nil == s.after {
		t.Fatal("验证失败")
	}

	if 30 != *s.after {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetSpace(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if nil != s.before {
		t.Fatal("验证失败")
	}
	if nil != s.after {
		t.Fatal("验证失败")
	}

	s.SetSpace(30)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if nil == s.before {
		t.Fatal("验证失败")
	}
	if nil == s.after {
		t.Fatal("验证失败")
	}

	if 30 != *s.before {
		t.Fatal("验证失败")
	}
	if 30 != *s.after {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetLine(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if nil != s.line {
		t.Fatal("验证失败")
	}

	s.SetLine(30)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if nil == s.line {
		t.Fatal("验证失败")
	}

	if 30 != *s.line {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetLineRule(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if nil != s.lineRule {
		t.Fatal("验证失败")
	}

	s.SetLineRule("")
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if nil == s.lineRule {
		t.Fatal("验证失败")
	}

	if "" != *s.lineRule {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetBeforeAutoSpacing(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if false != s.beforeAutoSpacing {
		t.Fatal("验证失败")
	}

	s.SetBeforeAutoSpacing(true)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if false == s.beforeAutoSpacing {
		t.Fatal("验证失败")
	}

	s.SetBeforeAutoSpacing(false)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if true == s.beforeAutoSpacing {
		t.Fatal("验证失败")
	}
}

func TestSpacing_SetAfterAutoSpacing(t *testing.T) {
	s := Spacing{}
	if true == s.isSet {
		t.Fatal("验证失败")
	}

	if false != s.afterAutoSpacing {
		t.Fatal("验证失败")
	}

	s.SetAfterAutoSpacing(true)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if false == s.afterAutoSpacing {
		t.Fatal("验证失败")
	}

	s.SetAfterAutoSpacing(false)
	if false == s.isSet {
		t.Fatal("验证失败")
	}

	if true == s.afterAutoSpacing {
		t.Fatal("验证失败")
	}
}

func TestSpacing_GetXmlBytes(t *testing.T) {
	s := Spacing{}

	act, err := s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	s.SetAfter(30)
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp := `<w:spacing w:after="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	s.SetBefore(30)
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:spacing w:after="30" w:before="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	s.SetLine(30)
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:spacing w:after="30" w:before="30" w:line="30"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	s.SetLineRule("")
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:spacing w:after="30" w:before="30" w:line="30" w:lineRule=""/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	s.SetBeforeAutoSpacing(true)
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:spacing w:after="30" w:before="30" w:line="30" w:lineRule="" w:beforeAutospacing="true"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}

	s.SetAfterAutoSpacing(true)
	act, err = s.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	exp = `<w:spacing w:after="30" w:before="30" w:line="30" w:lineRule="" w:beforeAutospacing="true" w:afterAutospacing="true"/>`
	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
