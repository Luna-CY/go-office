package xlsx

import (
	"testing"
)

func TestDocPropsCore_FilePath(t *testing.T) {
	dpc := DocPropsCore{}

	if "/docProps/core.xml" != dpc.FilePath() {
		t.Fatal("测试失败")
	}
}

func TestDocPropsCore_Marshal(t *testing.T) {
	dpc := DocPropsCore{}

	dpc.CpNameSpace = "test-cp-name-space"
	dpc.DcNameSpace = "test-dc-name-space"
	dpc.DctNameSpace = "test-dct-name-space"
	dpc.XsiNameSpace = "test-xsi-name-space"

	dpc.Creator = "Test Creator"
	dpc.LastModify = "Last Modify"
	dpc.Created = DocPropsCoreCreated{Type: "test-xsi-type", Created: "test-time-now"}
	dpc.Modified = DocPropsCoreModified{Type: "test-xsi-type", Created: "test-time-now"}

	content, err := dpc.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	exp := `<cp:coreProperties xmlns:cp="test-cp-name-space" xmlns:dc="test-dc-name-space" xmlns:dcterms="test-dct-name-space" xmlns:xsi="test-xsi-name-space"><dc:creator>Test Creator</dc:creator><cp:lastModifiedBy>Last Modify</cp:lastModifiedBy><dcterms:created xsi:type="test-xsi-type">test-time-now</dcterms:created><dcterms:modified xsi:type="test-xsi-type">test-time-now</dcterms:modified></cp:coreProperties>`
	if string(content) != exp {
		t.Fatal("测试失败")
	}
}
