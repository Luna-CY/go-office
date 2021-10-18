package xlsx

import "testing"

func TestWorkbookRelationship_FilePath(t *testing.T) {
	wrs := WorkbookRelationship{}

	if "/xl/_rels/workbook.xml.rels" != wrs.FilePath() {
		t.Fatal("测试失败")
	}
}

func TestWorkbookRelationship_Marshal(t *testing.T) {
	wrs := WorkbookRelationship{}
	wrs.Relationships = append(wrs.Relationships, Relationship{Id: "rId1", Type: "test-type", Target: "test-target"})

	content, err := wrs.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<Relationships xmlns=""><Relationship Id="rId1" Type="test-type" Target="test-target"></Relationship></Relationships>` != string(content) {
		t.Fatal("测试失败")
	}
}
