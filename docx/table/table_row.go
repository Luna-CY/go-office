package table

import (
    "errors"
    "fmt"
)

// Row 表格行结构定义
type Row struct {
    // 单元格列表
    cells []*Cell
}

// addCell 添加一个单元格
func (r *Row) addCell() {
    r.cells = append(r.cells, new(Cell))
}

// GetCells 获取全部单元格
func (r *Row) GetCells() []*Cell {
    return r.cells
}

// GetCell 获取指定位置的单元格
func (r *Row) GetCell(index uint) (*Cell, error) {
    if index > uint(len(r.cells)) {
        return nil, errors.New(fmt.Sprintf("索引溢出"))
    }

    return r.cells[index], nil
}

// addCellWithIndex 添加一列到指定位置
func (r *Row) addCellWithIndex(index uint) {
    if index == uint(len(r.cells)) {
        r.cells = append(r.cells, new(Cell))

        return
    }

    before := r.cells[:index]
    after := r.cells[index:]

    r.cells = make([]*Cell, len(r.cells)+1)
    r.cells = append(r.cells, before...)
    r.cells[index] = new(Cell)
    r.cells = append(r.cells, after...)
}
