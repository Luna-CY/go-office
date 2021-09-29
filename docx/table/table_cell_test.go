package table

import "testing"

func TestCell_AddParagraph(t *testing.T) {
	c := Cell{}
	if 0 != len(c.contents) {
		t.Fatal("验证失败")
	}

	p := c.AddParagraph()
	if 1 != len(c.contents) {
		t.Fatal("验证失败")
	}

	if nil == p {
		t.Fatal("验证失败")
	}

	if ContentTypeParagraph != c.contents[0].ct {
		t.Fatal("验证失败")
	}

	if nil == c.contents[0].paragraph {
		t.Fatal("验证失败")
	}
}

func TestCell_AddTable(t *testing.T) {
	c := Cell{}
	if 0 != len(c.contents) {
		t.Fatal("验证失败")
	}

	tab := c.AddTable()
	if 1 != len(c.contents) {
		t.Fatal("验证失败")
	}

	if nil == tab {
		t.Fatal("验证失败")
	}

	if ContentTypeTable != c.contents[0].ct {
		t.Fatal("验证失败")
	}

	if nil == c.contents[0].table {
		t.Fatal("验证失败")
	}
}

func TestCell_GetContents(t *testing.T) {
	c := Cell{}

	if 0 != len(c.GetContents()) {
		t.Fatal("验证失败")
	}

	c.AddParagraph()

	if 1 != len(c.GetContents()) {
		t.Fatal("验证失败")
	}
}

func TestCell_GetProperties(t *testing.T) {
	c := Cell{}

	if nil == c.GetProperties() {
		t.Fatal("验证失败")
	}
}
