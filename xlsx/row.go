package xlsx

import (
	"fmt"
	"sync"
)

// Row 行结构定义
type Row struct {
	// id 行ID
	id string

	cm sync.RWMutex
	// cells 单元格
	cells []*Cell
}

// AddCell 添加一个单元格
func (r *Row) AddCell() *Cell {
	return r.AddCellWithContent("")
}

// AddCellWithContent 添加一个单元格并设置内容
func (r *Row) AddCellWithContent(content interface{}) *Cell {
	c := new(Cell)
	c.SetContent(content)

	r.cm.Lock()
	defer r.cm.Unlock()

	c.id = fmt.Sprintf("%c%v", 'A'+len(r.cells), r.id)

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
