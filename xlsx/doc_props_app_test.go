package xlsx

import "testing"

func TestDocPropsApp_FilePath(t *testing.T) {
	dpa := DocPropsApp{}

	if "/docProps/app.xml" != dpa.FilePath() {
		t.Fatal("测试失败")
	}
}

func TestDocPropsApp_Marshal(t *testing.T) {
	dpa := DocPropsApp{AppVersion: "1.0.0", Application: "Test Application", DocSecurity: 1, RootNameSpace: "test-space"}

	content, err := dpa.Marshal()
	if nil != err {
		t.Fatal(err)
	}

	if `<Properties xmlns="test-space"><Application>Test Application</Application><DocSecurity>1</DocSecurity><AppVersion>1.0.0</AppVersion></Properties>` != string(content) {
		t.Fatal("测试失败")
	}
}
