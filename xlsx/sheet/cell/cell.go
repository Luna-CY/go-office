package cell

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

const (
	STCellTypeBoolean   = "b"
	STCellTypeDate      = "d"
	STCellTypeNumber    = "n"
	STCellTypeInlineStr = "inlineStr"
	STCellTypeString    = "s"
)

func newBaseCell(row uint64, col uint64) *Cell {
	c := new(Cell)

	c.x = row
	c.y = col

	c.Reference = fmt.Sprintf("%v%v", ten2col(col), row)

	return c
}

func NewCellWithTextForInline(row uint64, col uint64, text string) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeInlineStr
	c.Value = text

	return c
}

func NewCellWithText(row uint64, col uint64, text string) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeString
	c.Value = text

	return c
}

func NewCellWithBoolean(row uint64, col uint64, value bool) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeBoolean
	c.Value = fmt.Sprintf("%v", value)

	return c
}

func NewCellWithDate(row uint64, col uint64, date time.Time) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeDate
	c.Value = date.Format(time.RFC3339)

	return c
}

func NewCellWithNumber(row uint64, col uint64, number int64) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeNumber
	c.Value = strconv.FormatInt(number, 10)

	return c
}

// Cell 单元格结构定义
type Cell struct {
	Name xml.Name `xml:"c"`

	Reference string `xml:"r,attr"`
	Style     uint   `xml:"s,attr"`
	Type      string `xml:"t,attr"`

	x uint64
	y uint64

	Formula *Formula
	Value   string `xml:"v"`
}

func ten2col(col uint64) string {
	builder := new(bytes.Buffer)

	for col > 0 {
		m := col % 26
		if 0 == m {
			m = 26
		}

		builder.WriteByte(byte(m + 64))
		col = (col - m) / 26
	}

	bs := builder.Bytes()
	for i, j := 0, builder.Len()-1; i < j; {
		bs[j], bs[i] = bs[i], bs[j]
		i++
		j--
	}

	return string(bs)
}
