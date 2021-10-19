package workbook

import (
	"github.com/Luna-CY/go-office/xlsx"
	"testing"
)

func TestWorkbookRelationship_FilePath(t *testing.T) {
	wrs := BookRelationship{}

	if "xl/_rels/workbook.xml.rels" != wrs.Filepath() {
		t.Fatal("测试失败")
	}
}

func TestWorkbookRelationship_Marshal(t *testing.T) {
	wrs := BookRelationship{}
	wrs.Relationships = append(wrs.Relationships, xlsx.Relationship{Id: "rId1", Type: "test-type", Target: "test-target"})

	content, err := wrs.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Relationships xmlns=""><BookRelationship Id="rId1" Type="test-type" Target="test-target"></BookRelationship></Relationships>` != string(content) {
		t.Fatal("测试失败")
	}
}
