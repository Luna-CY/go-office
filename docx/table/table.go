package table

import (
    "errors"
    "fmt"
    "sync"
)

type Table struct {
    gm sync.RWMutex
    // gridCols 列定义
    gridCols []*GridCol

    // tblPr 表格属性配置结构
    tblPr *TblPr

    rm sync.RWMutex
    // rows 行列表
    rows []*Row
}

func (t *Table) GetProperties() *TblPr {
    if nil == t.tblPr {
        t.tblPr = new(TblPr)
    }

    return t.tblPr
}

// AddCol 添加新的列
func (t *Table) AddCol() {
    defaultWidth := 2000
    t.addColWithWidth(&defaultWidth)
}

// AddColWithWidth 添加一列并设置列宽
func (t *Table) AddColWithWidth(width int) {
    t.addColWithWidth(&width)
}

func (t *Table) addColWithWidth(width *int) {
    col := new(GridCol)
    col.w = width

    t.gm.Lock()
    t.gridCols = append(t.gridCols, col)
    t.gm.Unlock()

    for _, r := range t.GetRows() {
        r.addCell()
    }
}

// AddColWithIndex 插入到指定索引
func (t *Table) AddColWithIndex(index uint) {
    t.addColWithIndexAndWidth(index, nil)
}

func (t *Table) AddColWithIndexAndWidth(index uint, width int) {
    t.addColWithIndexAndWidth(index, &width)
}

func (t *Table) addColWithIndexAndWidth(index uint, width *int) {
    if index > uint(len(t.gridCols)) {
        index = uint(len(t.gridCols))
    }

    col := new(GridCol)
    col.w = width

    if index == uint(len(t.gridCols)) {
        t.gm.Lock()
        t.gridCols = append(t.gridCols, col)
        t.gm.Unlock()

        for _, r := range t.rows {
            r.addCellWithIndex(index)
        }

        return
    }

    t.gm.Lock()
    before := t.gridCols[:index]
    after := t.gridCols[index:]

    t.gridCols = make([]*GridCol, len(t.gridCols)+1)
    t.gridCols = append(t.gridCols, before...)
    t.gridCols[index] = col
    t.gridCols = append(t.gridCols, after...)
    t.gm.Unlock()

    for _, r := range t.rows {
        r.addCellWithIndex(index)
    }
}

// GetCol 获取指定列结构
func (t *Table) GetCol(index uint) (*GridCol, error) {
    if index > uint(len(t.gridCols)) {
        return nil, errors.New(fmt.Sprintf("索引溢出"))
    }

    t.gm.RLock()
    defer t.gm.RUnlock()

    return t.gridCols[index], nil
}

// GetCols 获取所有列结构
func (t *Table) GetCols() []*GridCol {
    t.gm.RLock()
    defer t.gm.RUnlock()

    return t.gridCols
}

// AddRow 添加一行
func (t *Table) AddRow() *Row {
    r := new(Row)
    for i := 0; i < len(t.gridCols); i++ {
        r.addCell()
    }

    t.rm.Lock()
    t.rows = append(t.rows, r)
    t.rm.Unlock()

    return r
}

// AddRowWithIndex 添加一行到指定位置
func (t *Table) AddRowWithIndex(index uint) *Row {
    r := new(Row)
    for i := 0; i < len(t.gridCols); i++ {
        r.addCell()
    }

    if index > uint(len(t.rows)) {
        index = uint(len(t.rows))
    }

    if index == uint(len(t.rows)) {
        t.rm.Lock()
        t.rows = append(t.rows, r)
        t.rm.Unlock()

        return r
    }

    before := t.rows[:index]
    after := t.rows[index:]

    t.rows = make([]*Row, len(t.rows)+1)
    t.rows = append(t.rows, before...)
    t.rows[index] = r
    t.rows = append(t.rows, after...)

    return r
}

// GetRow 获取指定行结构
func (t *Table) GetRow(index uint) (*Row, error) {
    if index > uint(len(t.rows)) {
        return nil, errors.New(fmt.Sprintf("索引溢出"))
    }

    t.rm.RLock()
    defer t.rm.RUnlock()

    return t.rows[index], nil
}

func (t *Table) GetRows() []*Row {
    t.rm.RLock()
    defer t.rm.RUnlock()

    return t.rows
}
