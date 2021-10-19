package sheet

import (
	"github.com/Luna-CY/go-office/xlsx/sheet/row"
	"strconv"
)

// AddRow 添加一行数据
func (s *Sheet) AddRow() *row.Row {
	r := new(row.Row)
	r.Id = strconv.Itoa(len(s.rows) + 1)

	s.rm.Lock()
	defer s.rm.Unlock()

	s.rows = append(s.rows, r)

	return r
}
