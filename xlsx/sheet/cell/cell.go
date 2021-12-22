package cell

import (
	"bytes"
	"fmt"
)

const (
	STCellTypeNumber       = "n"
	STCellTypeInlineStr    = "inlineStr"
	STCellTypeSharedString = "s"
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
	c.IS = &inlineStr{T: text}

	return c
}

func NewCellWithText(row uint64, col uint64, text string) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeSharedString
	c.Value = &text

	return c
}

func NewCellWithNumber(row uint64, col uint64, number interface{}) *Cell {
	c := newBaseCell(row, col)
	c.Type = STCellTypeNumber
	v := fmt.Sprintf("%v", number)
	c.Value = &v

	return c
}

// Cell 单元格结构定义
type Cell struct {
	Reference string `xml:"r,attr"`
	Style     *uint  `xml:"s,attr"`
	Type      string `xml:"t,attr"`

	x uint64
	y uint64

	Formula *Formula
	Value   *string    `xml:"v"`
	IS      *inlineStr `xml:"is"`
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
