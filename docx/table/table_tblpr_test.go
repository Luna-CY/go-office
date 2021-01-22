package table

import "testing"

func TestTblPr_SetWidth(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.width {
        t.Fatal("验证失败")
    }

    tblpr.SetWidth(30)
    if nil == tblpr.width {
        t.Fatal("验证失败")
    }
    if 30 != *tblpr.width {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetCaption(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.caption {
        t.Fatal("验证失败")
    }

    tblpr.SetCaption("TEST")
    if nil == tblpr.caption {
        t.Fatal("验证失败")
    }
    if "TEST" != *tblpr.caption {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetCellSpacing(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.cellSpacing {
        t.Fatal("验证失败")
    }

    tblpr.SetCellSpacing(30)
    if nil == tblpr.cellSpacing {
        t.Fatal("验证失败")
    }
    if 30 != *tblpr.cellSpacing {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetWidthAuto(t *testing.T) {
    tblpr := TblPr{}
    tblpr.SetWidth(30)

    tblpr.SetWidthAuto()
    if nil != tblpr.width {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetIndentation(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.indentation {
        t.Fatal("验证失败")
    }

    tblpr.SetIndentation(30)
    if nil == tblpr.indentation {
        t.Fatal("验证失败")
    }
    if 30 != *tblpr.indentation {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetHorizontalAlignment(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.horizontalAlignment {
        t.Fatal("验证失败")
    }

    tblpr.SetHorizontalAlignment(HorizontalAlignmentCenter)
    if nil == tblpr.horizontalAlignment {
        t.Fatal("验证失败")
    }
    if HorizontalAlignmentCenter != *tblpr.horizontalAlignment {
        t.Fatal("验证失败")
    }
}

func TestTblPr_SetLayout(t *testing.T) {
    tblpr := TblPr{}
    if nil != tblpr.layout {
        t.Fatal("验证失败")
    }

    tblpr.SetLayout(LayoutTypeAuto)
    if nil == tblpr.layout {
        t.Fatal("验证失败")
    }
    if LayoutTypeAuto != *tblpr.layout {
        t.Fatal("验证失败")
    }
}

func TestTblPr_GetBorder(t *testing.T) {
    tblpr := TblPr{}
    if nil == tblpr.GetBorder() {
        t.Fatal("验证失败")
    }
}

func TestTblPr_GetCellMargin(t *testing.T) {
    tblpr := TblPr{}
    if nil == tblpr.GetCellMargin() {
        t.Fatal("验证失败")
    }
}

func TestTblPr_GetWidth(t *testing.T) {
    tblpr := TblPr{}
    if 0 != tblpr.GetWidth() {
        t.Fatal("验证失败")
    }

    tblpr.SetWidth(30)
    if 30 != tblpr.GetWidth() {
        t.Fatal("验证失败")
    }
}
