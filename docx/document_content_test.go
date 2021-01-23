package docx

import "testing"

func TestDocument_AddParagraph(t *testing.T) {
    doc := Document{}
    if 0 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    doc.AddParagraph()
    if 1 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    doc.AddParagraph()
    if 2 != len(doc.contents) {
        t.Fatal("验证失败")
    }
}

func TestDocument_AddTable(t *testing.T) {
    doc := Document{}
    if 0 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    doc.AddTable()
    if 1 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    doc.AddTable()
    if 2 != len(doc.contents) {
        t.Fatal("验证失败")
    }
}

func TestDocument_AddTableWithColumns(t *testing.T) {
    doc := Document{}
    if 0 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    tab := doc.AddTableWithColumns(3)
    if 1 != len(doc.contents) {
        t.Fatal("验证失败")
    }

    if 3 != len(tab.GetCols()) {
        t.Fatal("验证失败")
    }
}
