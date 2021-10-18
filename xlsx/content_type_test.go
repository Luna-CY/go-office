package xlsx

import "testing"

func TestContentType_FileName(t *testing.T) {
	ct := ContentType{}
	if "/[Content_Types].xml" != ct.FilePath() {
		t.Fatal("测试失败")
	}
}

func TestContentType_Marshal(t *testing.T) {
	ct := ContentType{RootNameSpace: "test-namespace"}
	ct.Defaults = append(ct.Defaults, Default{Extension: "test", ContentType: "123"})
	ct.Overrides = append(ct.Overrides, Override{PartName: "test", ContentType: "123"})

	content, err := ct.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<Types xmlns="test-namespace"><Default Extension="test" ContentType="123"></Default><Override PartName="test" ContentType="123"></Override></Types>` != string(content) {
		t.Fatal("测试失败")
	}
}
