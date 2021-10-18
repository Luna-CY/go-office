package docx

import "testing"

func TestBackground_SetBackground(t *testing.T) {
	b := Background{}
	b.SetBackground("COLOR", "THEME_COLOR", "THEME_SHADE", "THEME_TINT")

	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if "COLOR" != b.color {
		t.Fatal("验证失败")
	}

	if "THEME_COLOR" != b.themeColor {
		t.Fatal("验证失败")
	}

	if "THEME_SHADE" != b.themeShade {
		t.Fatal("验证失败")
	}

	if "THEME_TINT" != b.themeTint {
		t.Fatal("验证失败")
	}

	exp := `<w:background w:color="COLOR" w:themeColor="THEME_COLOR" w:themeShade="THEME_SHADE" w:themeTint="THEME_TINT"/>`
	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetColor(t *testing.T) {
	b := Background{}
	b.SetColor("COLOR")

	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if "COLOR" != b.color {
		t.Fatal("验证失败")
	}

	exp := `<w:background w:color="COLOR"/>`
	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetThemeColor(t *testing.T) {
	b := Background{}
	b.SetThemeColor("COLOR")

	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if "COLOR" != b.themeColor {
		t.Fatal("验证失败")
	}

	exp := `<w:background w:themeColor="COLOR"/>`
	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetThemeShade(t *testing.T) {
	b := Background{}
	b.SetThemeShade("COLOR")

	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if "COLOR" != b.themeShade {
		t.Fatal("验证失败")
	}

	exp := `<w:background w:themeShade="COLOR"/>`
	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestBackground_SetThemeTint(t *testing.T) {
	b := Background{}
	b.SetThemeTint("COLOR")

	if false == b.isSet {
		t.Fatal("验证失败")
	}

	if "COLOR" != b.themeTint {
		t.Fatal("验证失败")
	}

	exp := `<w:background w:themeTint="COLOR"/>`
	act, err := b.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
