package cell

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
func (c *CellMargin) SetGroup(topAndBottom int, leftAndRight int) *CellMargin {
    c.isSet = true

    c.top = &topAndBottom
    c.bottom = &topAndBottom

    c.left = &leftAndRight
    c.right = &leftAndRight

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

// GetRight
func (c *CellMargin) GetRight() int {
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

    buffer.WriteString(template.TableCellTcPrCellMarginStart)

    if nil != c.top {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrCellMarginTopTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TableCellTcPrW, *c.top, template.TableCellTcPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.right {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrCellMarginRightTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TableCellTcPrW, *c.right, template.TableCellTcPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.bottom {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrCellMarginBottomTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TableCellTcPrW, *c.bottom, template.TableCellTcPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != c.left {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrCellMarginLeftTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TableCellTcPrW, *c.left, template.TableCellTcPrType, "dxa"))
        buffer.WriteString("/>")
    }

    buffer.WriteString(template.TableCellTcPrCellMarginEnd)

    return buffer.Bytes(), nil
}
