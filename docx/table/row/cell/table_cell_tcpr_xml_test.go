package cell

import (
	"testing"
)

func TestTcPr_GetXmlBytes(t *testing.T) {
	tcpr := TcPr{}

	act, err := tcpr.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if `<w:tcPr><w:tcW w:w="0" w:type="auto"/></w:tcPr>` != string(act) {
		t.Fatal("验证失败")
	}

	tcpr.SetWidth(30).SetTcFitText(true).SetVMerge(VMergeTypeRestart).SetNoWrap(true).SetGridSpan(3)

	exp := `<w:tcPr><noWrap/><w:tcFitText w:val="true"/><w:gridSpan w:val="3"/><vMerge w:val="restart"/><w:tcW w:w="30" w:type="dxa"/></w:tcPr>`
	act, err = tcpr.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
