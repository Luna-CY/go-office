package table

import "testing"

func TestRow_AddCellText(t *testing.T) {
	row := Row{}

	row.addCell()
	row.addCell()
	row.addCell()

	err := row.AddCellText(1, 2, 3)
	if nil != err {
		t.Fatal("验证失败")
	}

	if 1 != len(row.cells[0].GetContents()) {
		t.Fatal("验证失败")
	}
	if 1 != len(row.cells[1].GetContents()) {
		t.Fatal("验证失败")
	}
	if 1 != len(row.cells[2].GetContents()) {
		t.Fatal("验证失败")
	}

	err = row.AddCellText(1, 2, 3)
	if nil != err {
		t.Fatal("验证失败")
	}

	if 2 != len(row.cells[0].GetContents()) {
		t.Fatal("验证失败")
	}
	if 2 != len(row.cells[1].GetContents()) {
		t.Fatal("验证失败")
	}
	if 2 != len(row.cells[2].GetContents()) {
		t.Fatal("验证失败")
	}

	err = row.AddCellText(1, 2, 3, 4)
	if nil == err {
		t.Fatal("验证失败")
	}

	if 2 != len(row.cells[0].GetContents()) {
		t.Fatal("验证失败")
	}
	if 2 != len(row.cells[1].GetContents()) {
		t.Fatal("验证失败")
	}
	if 2 != len(row.cells[2].GetContents()) {
		t.Fatal("验证失败")
	}
}

func TestRow_GetCell(t *testing.T) {
	row := Row{}

	c, err := row.GetCell(0)
	if nil == err {
		t.Fatal("验证失败")
	}

	if nil != c {
		t.Fatal("验证失败")
	}

	row.addCell()
	c, err = row.GetCell(0)
	if nil != err {
		t.Fatal("验证失败")
	}

	if nil == c {
		t.Fatal("验证失败")
	}
}

func TestRow_GetProperties(t *testing.T) {
	row := Row{}

	if nil == row.GetProperties() {
		t.Fatal("验证失败")
	}
}

func TestRow_GetCells(t *testing.T) {
	row := Row{}

	if 0 != len(row.GetCells()) {
		t.Fatal("验证失败")
	}

	row.addCell()
	if 1 != len(row.GetCells()) {
		t.Fatal("验证失败")
	}
}

func TestRow_addCell(t *testing.T) {
	row := Row{}

	if 0 != len(row.cells) {
		t.Fatal("验证失败")
	}

	row.addCell()
	if 1 != len(row.cells) {
		t.Fatal("验证失败")
	}

	if 0 != row.cells[0].GetProperties().GetWidth() {
		t.Fatal("验证失败")
	}
}

func TestRow_addCellWithWidth(t *testing.T) {
	row := Row{}

	if 0 != len(row.cells) {
		t.Fatal("验证失败")
	}

	row.addCellWithWidth(30)
	if 1 != len(row.cells) {
		t.Fatal("验证失败")
	}

	if 30 != row.cells[0].GetProperties().GetWidth() {
		t.Fatal("验证失败")
	}
}

func TestRow_addCellWithIndexAndWidth(t *testing.T) {
	row := Row{}

	if 0 != len(row.cells) {
		t.Fatal("验证失败")
	}

	row.addCellWithIndexAndWidth(1, 30)
	if 1 != len(row.cells) {
		t.Fatal("验证失败")
	}

	if 30 != row.cells[0].GetProperties().GetWidth() {
		t.Fatal("验证失败")
	}

	row.addCellWithIndexAndWidth(1, 50)
	if 2 != len(row.cells) {
		t.Fatal("验证失败")
	}

	if 50 != row.cells[1].GetProperties().GetWidth() {
		t.Fatal("验证失败")
	}

	row.addCellWithIndexAndWidth(1, 80)
	if 3 != len(row.cells) {
		t.Fatal("验证失败")
	}

	if 80 != row.cells[1].GetProperties().GetWidth() {
		t.Fatal("验证失败")
	}
}
