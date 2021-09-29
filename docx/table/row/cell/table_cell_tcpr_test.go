package cell

import "testing"

func TestTcPr_GetBackground(t *testing.T) {
	tcpr := TcPr{}
	if nil == tcpr.GetBackground() {
		t.Fatal("验证失败")
	}
}

func TestTcPr_GetBorder(t *testing.T) {
	tcpr := TcPr{}
	if nil == tcpr.GetBorder() {
		t.Fatal("验证失败")
	}
}

func TestTcPr_GetMargin(t *testing.T) {
	tcpr := TcPr{}
	if nil == tcpr.GetMargin() {
		t.Fatal("验证失败")
	}
}

func TestTcPr_GetWidth(t *testing.T) {
	tcpr := TcPr{}
	if 0 != tcpr.GetWidth() {
		t.Fatal("验证失败")
	}

	width := 30
	tcpr.width = &width

	if 30 != tcpr.GetWidth() {
		t.Fatal("验证失败")
	}
}

func TestTcPr_SetGridSpan(t *testing.T) {
	tcpr := TcPr{}

	if nil != tcpr.gridSpan {
		t.Fatal("验证失败")
	}

	tcpr.SetGridSpan(3)
	if nil == tcpr.gridSpan {
		t.Fatal("验证失败")
	}
	if 3 != *tcpr.gridSpan {
		t.Fatal("验证失败")
	}
}

func TestTcPr_SetNoWrap(t *testing.T) {
	tcpr := TcPr{}

	if false != tcpr.noWrap {
		t.Fatal("验证失败")
	}

	tcpr.SetNoWrap(true)
	if true != tcpr.noWrap {
		t.Fatal("验证失败")
	}
	tcpr.SetNoWrap(false)
	if false != tcpr.noWrap {
		t.Fatal("验证失败")
	}
}

func TestTcPr_SetTcFitText(t *testing.T) {
	tcpr := TcPr{}

	if false != tcpr.tcFitText {
		t.Fatal("验证失败")
	}

	tcpr.SetTcFitText(true)
	if true != tcpr.tcFitText {
		t.Fatal("验证失败")
	}
	tcpr.SetTcFitText(false)
	if false != tcpr.tcFitText {
		t.Fatal("验证失败")
	}
}

func TestTcPr_SetWidth(t *testing.T) {
	tcpr := TcPr{}

	if nil != tcpr.width {
		t.Fatal("验证失败")
	}

	tcpr.SetWidth(30)
	if nil == tcpr.width {
		t.Fatal("验证失败")
	}
	if 30 != *tcpr.width {
		t.Fatal("验证失败")
	}
}

func TestTcPr_SetVMerge(t *testing.T) {
	tcpr := TcPr{}

	if nil != tcpr.vMerge {
		t.Fatal("验证失败")
	}

	tcpr.SetVMerge(VMergeTypeRestart)
	if nil == tcpr.vMerge {
		t.Fatal("验证失败")
	}

	if VMergeTypeRestart != *tcpr.vMerge {
		t.Fatal("验证失败")
	}
}
