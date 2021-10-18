package row

import "testing"

func TestTrPr_SetCantSplit(t *testing.T) {
	trpr := TrPr{}

	if false != trpr.cantSplit {
		t.Fatal("验证失败")
	}

	trpr.SetCantSplit(true)
	if true != trpr.cantSplit {
		t.Fatal("验证失败")
	}

	trpr.SetCantSplit(false)
	if false != trpr.cantSplit {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetCellSpacing(t *testing.T) {
	trpr := TrPr{}

	if nil != trpr.cellSpacing {
		t.Fatal("验证失败")
	}

	trpr.SetCellSpacing(30)
	if nil == trpr.cellSpacing {
		t.Fatal("验证失败")
	}

	if 30 != *trpr.cellSpacing {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetHeader(t *testing.T) {
	trpr := TrPr{}

	if false != trpr.header {
		t.Fatal("验证失败")
	}

	trpr.SetHeader(true)
	if true != trpr.header {
		t.Fatal("验证失败")
	}

	trpr.SetHeader(false)
	if false != trpr.header {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetHeight(t *testing.T) {
	trpr := TrPr{}

	if nil != trpr.height {
		t.Fatal("验证失败")
	}

	trpr.SetHeight(30)
	if nil == trpr.height {
		t.Fatal("验证失败")
	}

	if 30 != *trpr.height {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetHeightWithRule(t *testing.T) {
	trpr := TrPr{}

	if nil != trpr.height {
		t.Fatal("验证失败")
	}

	trpr.SetHeightWithRule(30, HeightRuleTypeAuto)
	if nil == trpr.height {
		t.Fatal("验证失败")
	}

	if 30 != *trpr.height {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetHidden(t *testing.T) {
	trpr := TrPr{}

	if false != trpr.hidden {
		t.Fatal("验证失败")
	}

	trpr.SetHidden(true)
	if true != trpr.hidden {
		t.Fatal("验证失败")
	}

	trpr.SetHidden(false)
	if false != trpr.hidden {
		t.Fatal("验证失败")
	}
}

func TestTrPr_SetHorizontalAlignment(t *testing.T) {
	trpr := TrPr{}

	if nil != trpr.horizontalAlignment {
		t.Fatal("验证失败")
	}

	trpr.SetHorizontalAlignment(HorizontalAlignmentCenter)
	if nil == trpr.horizontalAlignment {
		t.Fatal("验证失败")
	}

	if HorizontalAlignmentCenter != *trpr.horizontalAlignment {
		t.Fatal("验证失败")
	}
}
