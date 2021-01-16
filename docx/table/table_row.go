package table

import (
    "errors"
    "fmt"
    "github.com/Luna-CY/go-office/docx/table/row"
    "sync"
)

// Row 表格行结构定义
type Row struct {
    // pr 表格行的属性指针
    pr *row.TrPr

    cm sync.RWMutex
    // 单元格列表
    cells []*Cell
}

// GetProperties 获取属性配置结构
func (r *Row) GetProperties() *row.TrPr {
    if nil == r.pr {
        r.pr = new(row.TrPr)
    }

    return r.pr
}

// addCell 添加一个单元格
func (r *Row) addCell() {
    r.cm.Lock()
    defer r.cm.Unlock()

    cell := new(Cell)
    cell.GetProperties().SetWidth(2000)
    r.cells = append(r.cells, cell)
}

// GetCells 获取全部单元格
func (r *Row) GetCells() []*Cell {
    r.cm.RLock()
    r.cm.RUnlock()

    return r.cells
}

// GetCell 获取指定位置的单元格
func (r *Row) GetCell(index uint) (*Cell, error) {
    if index > uint(len(r.cells)) {
        return nil, errors.New(fmt.Sprintf("索引溢出"))
    }

    r.cm.RLock()
    defer r.cm.RUnlock()

    return r.cells[index], nil
}

// addCellWithIndex 添加一列到指定位置
func (r *Row) addCellWithIndex(index uint) {
    if index == uint(len(r.cells)) {
        r.cells = append(r.cells, new(Cell))

        return
    }

    r.cm.Lock()
    defer r.cm.Unlock()

    before := r.cells[:index]
    after := r.cells[index:]

    r.cells = make([]*Cell, len(r.cells)+1)
    r.cells = append(r.cells, before...)
    r.cells[index] = new(Cell)
    r.cells = append(r.cells, after...)
}
