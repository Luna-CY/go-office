package sheet

import (
	"github.com/Luna-CY/go-office/xlsx/sheet/row"
)

// NewRow 添加一行空行并返回行对象
func (s *Sheet) NewRow() *row.Row {
	return s.Data.NewRow()
}
