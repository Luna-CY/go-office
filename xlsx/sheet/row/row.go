package row

import (
	"encoding/xml"
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
