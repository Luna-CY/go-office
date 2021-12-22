package sheet

import (
	"encoding/xml"
	"fmt"
	"github.com/Luna-CY/go-office/xlsx/sheet/cell"
	"github.com/Luna-CY/go-office/xlsx/sheet/row"
	"sync"
)

func (s *Sheet) SetCellWithValue(ri, ci uint64, value interface{}) *Sheet {
	s.Data.SetCellWithValue(ri, ci, value)

	return s
}

func (s *Sheet) SetCellWithText(ri, ci uint64, text string) *Sheet {
	s.Data.SetCellWithText(ri, ci, text)

	return s
}

func (s *Sheet) SetCellWithNumber(ri, ci uint64, number interface{}) *Sheet {
	s.Data.SetCellWithNumber(ri, ci, number)

	return s
}

type Data struct {
	XMLName xml.Name `xml:"sheetData"`

	Rows []*row.Row `xml:"row"`

	dm   sync.RWMutex
	data map[uint64]map[uint64]*cell.Cell
}

// SetCellWithValue 设置单元格并将单元格内容设置为行内字符串
// 设置多行字符串请使用 SetCellWithText 方法
func (d *Data) SetCellWithValue(ri, ci uint64, value interface{}) *Data {
	d.dm.Lock()
	defer d.dm.Unlock()

	r, ok := d.data[ri]
	if !ok {
		d.data[ri] = map[uint64]*cell.Cell{}
		r = d.data[ri]
	}

	r[ci] = cell.NewCellWithTextForInline(ri, ci, fmt.Sprintf("%v", value))

	return d
}

// SetCellWithText 设置多行文本
// TODO 需要实现SharedStringTable
func (d *Data) SetCellWithText(ri, ci uint64, text string) *Data {
	d.dm.Lock()
	defer d.dm.Unlock()

	r, ok := d.data[ri]
	if !ok {
		d.data[ri] = map[uint64]*cell.Cell{}
		r = d.data[ri]
	}

	r[ci] = cell.NewCellWithText(ri, ci, text)

	return d
}

// SetCellWithNumber 设置单元格并将单元格内容设置为数字类型（整形与浮点型）
func (d *Data) SetCellWithNumber(ri, ci uint64, number interface{}) *Data {
	d.dm.Lock()
	defer d.dm.Unlock()

	r, ok := d.data[ri]
	if !ok {
		d.data[ri] = map[uint64]*cell.Cell{}
		r = d.data[ri]
	}

	r[ci] = cell.NewCellWithNumber(ri, ci, number)

	return d
}

func (d *Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	data := struct {
		Rows []*row.Row `xml:"row"`
	}{}

	for ri, cells := range d.data {
		r := row.NewRowWithIndex(ri)

		for _, c := range cells {
			r.AddCell(c)
		}

		data.Rows = append(data.Rows, r)
	}

	return e.EncodeElement(data, start)
}
