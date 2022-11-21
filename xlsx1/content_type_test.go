package xlsx1

import "testing"

func TestContentType_FileName(t *testing.T) {
	ct := ContentType{}
	if "[Content_Types].xml" != ct.Filepath() {
		t.Fatal("测试失败")
	}
}

func TestContentType_Marshal(t *testing.T) {
	ct := ContentType{RootNamespace: "test-namespace"}
	ct.Defaults = append(ct.Defaults, Default{Extension: "test", ContentType: "123"})
	ct.Overrides = append(ct.Overrides, Override{PartName: "test", ContentType: "123"})

	content, err := ct.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Types xmlns="test-namespace"><Default Extension="test" ContentType="123"></Default><Override PartName="test" ContentType="123"></Override></Types>` != string(content) {
		t.Fatal("测试失败")
	}
}
