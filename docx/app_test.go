package docx

import "testing"

func TestApp_GetXmlBytes(t *testing.T) {
    exp := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Properties xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"><Application>Test Application String</Application><DocSecurity>0</DocSecurity></Properties>`

    app := App{Application: "Test Application String", Security: 0}
    act, err := app.GetXmlBytes()
    if nil != err {
        t.Fatalf("生成XML失败: %v\n", err)
    }

    if exp != string(act) {
        t.Fatal("验证失败")
    }
}
