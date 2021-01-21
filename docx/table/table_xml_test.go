package table

import (
    "testing"
)

func TestTable_GetXmlBytes(t *testing.T) {
    tab := Table{}

    act, err := tab.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if "" != string(act) {
        t.Fatal("验证失败")
    }

    tab.addCol()
    act, err = tab.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if "" != string(act) {
        t.Fatal("验证失败")
    }

    tab.AddRow()
    act, err = tab.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    exp := `<w:tbl><w:tblPr><w:tblW w:w="0" w:type="auto"/></w:tblPr><w:tblGrid><w:gridCol/></w:tblGrid><w:tr><w:tc><w:tcPr><w:tcW w:w="0" w:type="auto"/></w:tcPr><w:p/></w:tc></w:tr></w:tbl>`
    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
