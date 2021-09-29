package table

import (
	"fmt"
	"testing"
)

func TestTable_AddCol(t *testing.T) {
	tab := Table{}
	if 0 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	r := tab.addCol()
	if nil == r {
		t.Fatal("验证失败")
	}
	if 1 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	if nil != tab.gridCols[0].w {
		t.Fatal("验证失败")
	}
}

func TestTable_AddColWithWidth(t *testing.T) {
	tab := Table{}
	if 0 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	r := tab.AddColWithWidth(30)
	if nil == r {
		t.Fatal("验证失败")
	}
	if 1 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	if nil == tab.gridCols[0].w {
		t.Fatal("验证失败")
	}
	if 30 != *tab.gridCols[0].w {
		t.Fatal("验证失败")
	}
}

func TestTable_AddColWithIndexAndWidth(t *testing.T) {
	tab := Table{}
	if 0 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	tab.addCol()
	r := tab.AddColWithIndexAndWidth(1, 30)
	if nil == r {
		t.Fatal("验证失败")
	}
	if 2 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	if nil == tab.gridCols[1].w {
		t.Fatal("验证失败")
	}
	if 30 != *tab.gridCols[1].w {
		t.Fatal("验证失败")
	}

	tab.AddColWithIndexAndWidth(0, 30)
	if 3 != len(tab.gridCols) {
		t.Fatal("验证失败")
	}

	if nil == tab.gridCols[0].w {
		t.Fatal("验证失败")
	}
	if 30 != *tab.gridCols[0].w {
		t.Fatal("验证失败")
	}
}

func TestTable_AddRow(t *testing.T) {
	tab := Table{}
	tab.addCol()
	tab.addCol()

	if 0 != len(tab.rows) {
		t.Fatal("验证失败")
	}

	r := tab.AddRow()
	if nil == r {
		t.Fatal("验证失败")
	}
	if 1 != len(tab.rows) {
		t.Fatal("验证失败")
	}

	if 2 != len(r.cells) {
		t.Fatal("验证失败")
	}
}

func TestTable_AddRowWithIndex(t *testing.T) {
	tab := Table{}
	tab.addCol()
	tab.addCol()

	if 0 != len(tab.rows) {
		t.Fatal("验证失败")
	}

	tab.AddRow()
	if 1 != len(tab.rows) {
		t.Fatal("验证失败")
	}

	r := tab.AddRowWithIndex(0)
	if nil == r {
		t.Fatal("验证失败")
	}
	if 2 != len(tab.rows) {
		t.Fatal("验证失败")
	}
	if 2 != len(r.cells) {
		t.Fatal("验证失败")
	}
	if fmt.Sprintf("%p", r) != fmt.Sprintf("%p", tab.rows[0]) {
		t.Fatal("验证失败")
	}

	r = tab.AddRowWithIndex(2)
	if nil == r {
		t.Fatal("验证失败")
	}
	if 3 != len(tab.rows) {
		t.Fatal("验证失败")
	}
	if 2 != len(r.cells) {
		t.Fatal("验证失败")
	}
	if fmt.Sprintf("%p", r) != fmt.Sprintf("%p", tab.rows[2]) {
		t.Fatal("验证失败")
	}

	r = tab.AddRowWithIndex(5)
	if nil == r {
		t.Fatal("验证失败")
	}
	if 4 != len(tab.rows) {
		t.Fatal("验证失败")
	}
	if fmt.Sprintf("%p", r) != fmt.Sprintf("%p", tab.rows[3]) {
		t.Fatal("验证失败")
	}
}

func TestTable_GetProperties(t *testing.T) {
	tab := Table{}
	if nil == tab.GetProperties() {
		t.Fatal("验证失败")
	}
}

func TestTable_GetCols(t *testing.T) {
	tab := Table{}
	tab.addCol()

	if 1 != len(tab.GetCols()) {
		t.Fatal("验证失败")
	}
}

func TestTable_GetCol(t *testing.T) {
	tab := Table{}
	tab.addCol()

	col, err := tab.GetCol(0)
	if nil != err {
		t.Fatal("验证失败")
	}

	if nil == col {
		t.Fatal("验证失败")
	}

	col, err = tab.GetCol(10)
	if nil == err {
		t.Fatal("验证失败")
	}

	if nil != col {
		t.Fatal("验证失败")
	}
}

func TestTable_GetRows(t *testing.T) {
	tab := Table{}
	tab.addCol()
	tab.AddRow()

	if 1 != len(tab.GetRows()) {
		t.Fatal("验证失败")
	}
}

func TestTable_GetRow(t *testing.T) {
	tab := Table{}
	tab.addCol()
	tab.AddRow()

	row, err := tab.GetRow(0)
	if nil != err {
		t.Fatal("验证失败")
	}

	if nil == row {
		t.Fatal("验证失败")
	}

	row, err = tab.GetRow(10)
	if nil == err {
		t.Fatal("验证失败")
	}

	if nil != row {
		t.Fatal("验证失败")
	}
}
