package workbook

import "testing"

func TestWorkbook_Filepath(t *testing.T) {
	wb := Workbook{}

	if "xl/workbook.xml" != wb.Filepath() {
		t.Fatal("测试失败")
	}
}

func TestWorkbook_Marshal(t *testing.T) {
	wb := Workbook{}

	wb.RootNamespace = "root-namespace"
	wb.RelationshipNamespace = "relationship-namespace"

	wb.Views = new(BookViews)
	wb.Sheets = new(Sheets)

	wb.AddView(BookView{XWindow: 100, YWindow: 100, WindowWidth: 1000, WindowHeight: 1000})
	wb.NewSheet("rId0")

	content, err := wb.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	exp := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><workbook xmlns="root-namespace" xmlns:r="relationship-namespace"><bookViews><workbookView xWindow="100" yWindow="100" windowWidth="1000" windowHeight="1000"></workbookView></bookViews><sheets><sheet name="Sheet1" sheetId="1" r:Id="rId0"></sheet></sheets></workbook>`
	if string(content) != exp {
		t.Fatal("测试失败")
	}
}
