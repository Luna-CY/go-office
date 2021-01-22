package row

import (
    "testing"
)

func TestTrPr_GetXmlBytes(t *testing.T) {
    trpr := TrPr{}

    act, err := trpr.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if `` != string(act) {
        t.Fatal("验证失败")
    }

    trpr.SetHorizontalAlignment(HorizontalAlignmentCenter).SetHidden(true).SetHeight(100)
    trpr.SetHeightWithRule(30, HeightRuleTypeAuto).SetCellSpacing(30).SetHeader(true)
    trpr.SetCantSplit(true)

    exp := `<w:trPr><w:cantSplit/><w:hidden/><w:tblHeader/><w:tblCellSpacing w:type="dxa" w:val="30"/><w:jc w:val="center"/><w:trHeight w:val="30" w:hRule="auto"/></w:trPr>`
    act, err = trpr.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
