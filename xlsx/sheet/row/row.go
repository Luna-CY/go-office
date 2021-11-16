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

	Reference    uint64  `xml:"r,attr"`            // 行索引， 无符号整数
	Style        uint32  `xml:"s,attr"`            // 样式索引，无符号整数
	Height       float32 `xml:"ht,attr"`           // 行高
	Hidden       bool    `xml:"hidden,attr"`       // 隐藏
	CustomFormat bool    `xml:"customFormat,attr"` // 自定义样式，当此项为true时 Style 字段有效
	CustomHeight bool    `xml:"customHeight,attr"` // 自定义行高，当此项为true时 Height 字段有效

	cm    sync.RWMutex
	Cells []*cell.Cell
}

// AddCell 添加一个单元格并设置单元格内容
// 该方法接收任意类型并将单元格类型设置为纯文本进行展示
func (r *Row) AddCell(value interface{}) *Row {
	r.cm.Lock()
	defer r.cm.Unlock()

	c := cell.NewCellWithTextForInline(r.Reference, uint64(len(r.Cells)+1), fmt.Sprintf("%v", value))
	r.Cells = append(r.Cells, c)

	return r
}

// AddCells 添加任意个单元格并设置内容
func (r *Row) AddCells(values ...interface{}) *Row {
	return r
}
