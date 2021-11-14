package row

import (
	"encoding/xml"
	"fmt"
	"github.com/Luna-CY/go-office/xlsx/sheet/cell"
	"sync"
)

// Row 行结构定义
type Row struct {
	Name xml.Name `xml:"row"`

	R            uint32  `xml:"r,attr"`            // 行索引， 无符号整数
	S            uint32  `xml:"s,attr"`            // 样式索引，无符号整数
	Ht           float32 `xml:"ht,attr"`           // 行高
	Hidden       bool    `xml:"hidden,attr"`       // 隐藏
	CustomFormat bool    `xml:"customFormat,attr"` // 自定义样式，当此项为true时 S 字段有效
	CustomHeight bool    `xml:"customHeight,attr"` // 自定义行高，当此项为true时 Ht 字段有效

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
