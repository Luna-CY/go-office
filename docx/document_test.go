package docx

import "testing"

func TestDocument_GetBackground(t *testing.T) {
    doc := Document{}

    if nil == doc.GetBackground() {
        t.Fatal("验证失败")
    }
}

func TestDocument_GetProperties(t *testing.T) {
    doc := Document{}

    if nil == doc.GetProperties() {
        t.Fatal("验证失败")
    }
}

func TestDocument_GetSection(t *testing.T) {
    doc := Document{}

    if nil == doc.GetSection() {
        t.Fatal("验证失败")
    }
}
