package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// CellMargin 单元格边距管理器
type CellMargin struct {
    // isSet 是否设置了单元格边距
    isSet bool

    // top 上边距
    top *int

    // right 右边距
    right *int

    // bottom 下边距
    bottom *int

    // left 左边距
    left *int
}

// SetMargin 设置全部边距
func (c *CellMargin) SetMargin(margin int) *CellMargin {
    c.isSet = true

    c.top = &margin
    c.right = &margin
    c.bottom = &margin
    c.left = &margin

    return c
}

// SetGroup 分组设置边距
func (c *CellMargin) SetGroup(topAndBottom int, LeftAndRight int) *CellMargin {
    c.isSet = true

    c.top = &topAndBottom
    c.bottom = &topAndBottom

    c.left = &LeftAndRight
    c.right = &LeftAndRight

    return c
}

// GetTop
func (c *CellMargin) GetTop() int {
    if nil == c.top {
        return 0
    }

    return *c.top
}

// SetTop 设置上边距
func (c *CellMargin) SetTop(top int) *CellMargin {
    c.isSet = true
    c.top = &top

    return c
}

// GetRight
func (c *CellMargin) GetRight() int {
    if nil == c.right {
        return 0
    }

    return *c.right
}

// SetRight 设置右边距
func (c *CellMargin) SetRight(right int) *CellMargin {
    c.isSet = true
    c.right = &right

    return c
}

// GetBottom
func (c *CellMargin) GetBottom() int {
    if nil == c.bottom {
        return 0
    }

    return *c.bottom
}

// SetBottom 设置下边距
func (c *CellMargin) SetBottom(bottom int) *CellMargin {
    c.isSet = true
    c.bottom = &bottom

    return c
}

// GetLeft
func (c *CellMargin) GetLeft() int {
    if nil == c.left {
        return 0
    }

    return *c.left
}

// SetLeft 设置左边距
func (c *CellMargin) SetLeft(left int) *CellMargin {
    c.isSet = true
    c.left = &left

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

    if nil != c.right {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginRightTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.right, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.bottom {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginBottomTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.bottom, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.left {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellMarginLeftTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *c.left, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    buffer.WriteString(template.TblPrCellMarginEnd)

    return buffer.Bytes(), nil
}
