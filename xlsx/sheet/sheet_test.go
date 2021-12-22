package sheet

import "testing"

func TestSheet_Filepath(t *testing.T) {
	sh := Sheet{id: "1", name: "sheet1", relationshipId: "rId1"}

	if "xl/worksheets/sheet1.xml" != sh.Filepath() {
		t.Fatal("测试失败")
	}
}

func TestSheet_Marshal(t *testing.T) {
	sh := Sheet{id: "1", name: "sheet1", relationshipId: "rId1"}

	sh.RootNamespace = "root-namespace"
	sh.Views = new(Views)
	sh.Views.AddView(View{TabSelected: 1, WorkbookViewId: 0})

	sh.Data = new(Data)

	content, err := sh.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	exp := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><worksheet xmlns="root-namespace"><sheetViews><sheetView tabSelected="1" workbookViewId="0"></sheetView></sheetViews><sheetData></sheetData></worksheet>`
	if string(content) != exp {
		t.Fatal("测试失败")
	}
}
