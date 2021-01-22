package table

import (
    "fmt"
    "testing"
)

func TestTblPr_GetInnerXmlBytes(t *testing.T) {
    tblpr := TblPr{}

    act, err := tblpr.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if `<w:tblPr><w:tblW w:w="0" w:type="auto"/></w:tblPr>` != string(act) {
        t.Fatal("验证失败")
    }

    tblpr.SetLayout(LayoutTypeAuto).SetHorizontalAlignment(HorizontalAlignmentCenter).SetCaption("TEST")
    tblpr.SetIndentation(30).SetWidth(30).SetCellSpacing(30)

    exp := `<w:tblPr><w:jc w:val="center"/><w:tblLayout w:type="autofit"/><w:tblCaption w:val="TEST"/><w:tblW w:w="30" w:type="dxa"/><w:tblCellSpacing w:w="30" w:type="dxa"/><w:tblInd w:w="30" w:val="30" w:type="dxa"/></w:tblPr>`
    act, err = tblpr.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }

    tblpr.GetBorder().SetBorder(BorderStyleSingle, "FF0000", 30, 30, false)

    exp = fmt.Sprintf(`<w:tblPr><w:tblStyle w:val="%p"/><w:jc w:val="center"/><w:tblLayout w:type="autofit"/><w:tblCaption w:val="TEST"/><w:tblW w:w="30" w:type="dxa"/><w:tblCellSpacing w:w="30" w:type="dxa"/><w:tblInd w:w="30" w:val="30" w:type="dxa"/></w:tblPr>`, &tblpr)
    act, err = tblpr.GetInnerXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}

func TestTblPr_GetExtraXmlBytes(t *testing.T) {
    tblpr := TblPr{}

    act, err := tblpr.GetExtraXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if `` != string(act) {
        t.Fatal("验证失败")
    }

    tblpr.GetBorder().SetTop(BorderStyleSingle, "FF0000", 30, 30, false)
    tblpr.GetCellMargin().SetTop(30)

    exp := `<w:tblPr><w:tblBorders><w:top w:val="single" w:sz="30" w:color="FF0000" w:space="30" w:shadow="false"/></w:tblBorders><w:tblCellMar><w:top w:w="30" w:type="dxa"/></w:tblCellMar></w:tblPr>`
    act, err = tblpr.GetExtraXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
