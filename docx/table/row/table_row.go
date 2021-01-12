package row

import (
    "errors"
    "fmt"
    "github.com/Luna-CY/go-office/docx/table/row/cell"
    "sync"
)

// Row 表格行结构定义
type Row struct {
    // pr 表格行的属性指针
    pr *TrPr

    cm sync.RWMutex
    // 单元格列表
    cells []*cell.Cell
}

// GetProperties 获取属性配置指针
func (r *Row) GetProperties() *TrPr {
    if nil == r.pr {
        r.pr = new(TrPr)
    }

    return r.pr
}

// AddCell 添加一个单元格
func (r *Row) AddCell() {
    r.cm.Lock()
    defer r.cm.Unlock()

    r.cells = append(r.cells, new(cell.Cell))
}

// GetCells 获取全部单元格
func (r *Row) GetCells() []*cell.Cell {
    r.cm.RLock()
    r.cm.RUnlock()

    return r.cells
}

// GetCell 获取指定位置的单元格
func (r *Row) GetCell(index uint) (*cell.Cell, error) {
    if index > uint(len(r.cells)) {
        return nil, errors.New(fmt.Sprintf("索引溢出"))
    }

    r.cm.RLock()
    defer r.cm.RUnlock()

    return r.cells[index], nil
}

// AddCellWithIndex 添加一列到指定位置
func (r *Row) AddCellWithIndex(index uint) {
    if index == uint(len(r.cells)) {
        r.cells = append(r.cells, new(cell.Cell))

        return
    }

    r.cm.Lock()
    defer r.cm.Unlock()

    before := r.cells[:index]
    after := r.cells[index:]

    r.cells = make([]*cell.Cell, len(r.cells)+1)
    r.cells = append(r.cells, before...)
    r.cells[index] = new(cell.Cell)
    r.cells = append(r.cells, after...)
}
