package row

import (
	"fmt"
	"github.com/Luna-CY/go-office/xlsx/sheet/cell"
	"sync"
)

// Row 行结构定义
type Row struct {
	// Id 行ID
	Id string

	cm sync.RWMutex
	// cells 单元格
	cells []*cell.Cell
}

// AddCell 添加一个单元格
func (r *Row) AddCell() *cell.Cell {
	return r.AddCellWithContent("")
}

// AddCellWithContent 添加一个单元格并设置内容
func (r *Row) AddCellWithContent(content interface{}) *cell.Cell {
	c := new(cell.Cell)
	c.SetContent(content)

	r.cm.Lock()
	defer r.cm.Unlock()

	c.Id = fmt.Sprintf("%c%v", 'A'+len(r.cells), r.Id)

	r.cells = append(r.cells, c)

	return c
}

// AddCellsWithContent 添加任意数量的单元格
func (r *Row) AddCellsWithContent(contents ...interface{}) *Row {
	for _, content := range contents {
		r.AddCellWithContent(content)
	}

	return r
}
