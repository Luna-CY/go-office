package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

type CellMargin struct {
    // isSet 是否设置了单元格边距
    isSet bool

    // top 上边距
    top *int

    // end 右边距
    end *int

    // bottom 下边距
    bottom *int

    // start 左边距
    start *int
}

// SetMargin 设置全部边距
func (c *CellMargin) SetMargin(margin int) *CellMargin {
    c.isSet = true

    c.top = &margin
    c.end = &margin
    c.bottom = &margin
    c.start = &margin

    return c
}

// SetGroup 分组设置边距
func (c *CellMargin) SetGroup(topAndBottom int, startAndEnd int) *CellMargin {
    c.isSet = true

    c.top = &topAndBottom
    c.bottom = &topAndBottom

    c.start = &startAndEnd
    c.end = &startAndEnd

    return c
}

// GetTop
func (c *CellMargin) GetTop() int {
    return *c.top
}

// SetTop 设置上边距
func (c *CellMargin) SetTop(top int) *CellMargin {
    c.isSet = true
    c.top = &top

    return c
}

// GetEnd
func (c *CellMargin) GetEnd() int {
    return *c.end
}

// SetEnd 设置右边距
func (c *CellMargin) SetEnd(end int) *CellMargin {
    c.isSet = true
    c.end = &end

    return c
}

// GetBottom
func (c *CellMargin) GetBottom() int {
    return *c.bottom
}

// SetBottom 设置下边距
func (c *CellMargin) SetBottom(bottom int) *CellMargin {
    c.isSet = true
    c.bottom = &bottom

    return c
}

func (c *CellMargin) GetStart() int {
    return *c.start
}

// SetStart 设置左边距
func (c *CellMargin) SetStart(start int) *CellMargin {
    c.isSet = true
    c.start = &start

    return c
}

func (c *CellMargin) GetXmlBytes() ([]byte, error) {
    if !c.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteString(template.TblPrCellMarginStart)

    if nil != c.top {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginTopTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.top, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.end {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginEndTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.end, template.TblPrType, "dxa"))
        buffer.WriteString("/>")

        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginRightTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.end, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.bottom {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginBottomTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.bottom, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.start {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginStartTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.start, template.TblPrType, "dxa"))
        buffer.WriteString("/>")

        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginLeftTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.start, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    buffer.WriteString(template.TblPrCellMarginEnd)

    return buffer.Bytes(), nil
}
