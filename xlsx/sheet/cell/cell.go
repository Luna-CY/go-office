package cell

import (
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

func newBaseCell(row int64, col int64) *Cell {
	c := new(Cell)

	c.x = row
	c.y = col

	c.Reference = "" // TODO col做26进制转换

	return c
}

func NewCellWithTextForInline(row int64, col int64, text string) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeInlineStr
	c.Value = text

	return c
}

func NewCellWithText(row int64, col int64, text string) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeString
	c.Value = text

	return c
}

func NewCellWithBoolean(row int64, col int64, value bool) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeBoolean
	c.Value = fmt.Sprintf("%v", value)

	return c
}

func NewCellWithDate(row int64, col int64, date time.Time) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeDate
	c.Value = date.Format(time.RFC3339)

	return c
}

func NewCellWithNumber(row int64, col int64, number int64) *Cell {
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

	x int64
	y int64

	Formula *Formula
	Value   string `xml:"v"`
}
