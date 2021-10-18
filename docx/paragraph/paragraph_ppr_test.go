package paragraph

import "testing"

func TestPPr_SetKeepNext(t *testing.T) {
	ppr := PPr{}
	if false != ppr.keepNext {
		t.Fatal("验证失败")
	}

	ppr.SetKeepNext(true)

	if false == ppr.keepNext {
		t.Fatal("验证失败")
	}

	ppr.SetKeepNext(false)

	if false != ppr.keepNext {
		t.Fatal("验证失败")
	}
}

func TestPPr_SetKeepLines(t *testing.T) {
	ppr := PPr{}
	if false != ppr.keepLines {
		t.Fatal("验证失败")
	}

	ppr.SetKeepLines(true)

	if false == ppr.keepLines {
		t.Fatal("验证失败")
	}

	ppr.SetKeepLines(false)

	if false != ppr.keepLines {
		t.Fatal("验证失败")
	}
}

func TestPPr_SetHorizontalAlignment(t *testing.T) {
	ppr := PPr{}
	if nil != ppr.horizontalAlignment {
		t.Fatal("验证失败")
	}

	ppr.SetHorizontalAlignment(HorizontalAlignmentCenter)
	if nil == ppr.horizontalAlignment {
		t.Fatal("验证失败")
	}

	if HorizontalAlignmentCenter != *ppr.horizontalAlignment {
		t.Fatal("验证失败")
	}
}

func TestPPr_GetBackground(t *testing.T) {
	ppr := PPr{}

	if nil == ppr.GetBackground() {
		t.Fatal("验证失败")
	}
}

func TestPPr_GetBorder(t *testing.T) {
	ppr := PPr{}

	if nil == ppr.GetBorder() {
		t.Fatal("验证失败")
	}
}

func TestPPr_GetIdentity(t *testing.T) {
	ppr := PPr{}

	if nil == ppr.GetIdentity() {
		t.Fatal("验证失败")
	}
}

func TestPPr_GetSection(t *testing.T) {
	ppr := PPr{}

	if nil == ppr.GetSection() {
		t.Fatal("验证失败")
	}
}

func TestPPr_GetSpacing(t *testing.T) {
	ppr := PPr{}

	if nil == ppr.GetSpacing() {
		t.Fatal("验证失败")
	}
}
