package run

import "testing"

func TestRun_AddBreakLine(t *testing.T) {
    r := Run{}
    if "" != r.body.String() {
        t.Fatal("验证失败")
    }

    r.AddBreakLine("", "")
    if `<w:br w:type="" w:clear=""/>` != r.body.String() {
        t.Fatal("验证失败")
    }
}

func TestRun_AddText(t *testing.T) {
    r := Run{}
    if "" != r.body.String() {
        t.Fatal("验证失败")
    }

    r.AddText("TEST")
    if `<w:t>TEST</w:t>` != r.body.String() {
        t.Fatal("验证失败")
    }
}

func TestRun_AddTextSpace(t *testing.T) {
    r := Run{}
    if "" != r.body.String() {
        t.Fatal("验证失败")
    }

    r.AddTextSpace("TEST")
    if `<w:t xml:space="preserve">TEST</w:t>` != r.body.String() {
        t.Fatal("验证失败")
    }
}

func TestRun_GetProperties(t *testing.T) {
    r := Run{}
    if nil == r.GetProperties() {
        t.Fatal("验证失败")
    }
}
