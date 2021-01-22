package cell

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// Margin 单元格边距管理器
type Margin struct {
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
func (c *Margin) SetMargin(margin int) *Margin {
    c.isSet = true

    c.top = &margin
    c.right = &margin
    c.bottom = &margin
    c.left = &margin

    return c
}

// SetGroup 分组设置边距
func (c *Margin) SetGroup(topAndBottom int, leftAndRight int) *Margin {
    c.isSet = true

    c.top = &topAndBottom
    c.bottom = &topAndBottom

    c.left = &leftAndRight
    c.right = &leftAndRight

    return c
}

// GetTop
func (c *Margin) GetTop() int {
    if nil == c.top {
        return 0
    }

    return *c.top
}

// SetTop 设置上边距
func (c *Margin) SetTop(top int) *Margin {
    c.isSet = true
    c.top = &top

    return c
}

// GetRight
func (c *Margin) GetRight() int {
    if nil == c.right {
        return 0
    }

    return *c.right
}

// SetRight 设置右边距
func (c *Margin) SetRight(right int) *Margin {
    c.isSet = true
    c.right = &right

    return c
}

// GetBottom
func (c *Margin) GetBottom() int {
    if nil == c.bottom {
        return 0
    }

    return *c.bottom
}

// SetBottom 设置下边距
func (c *Margin) SetBottom(bottom int) *Margin {
    c.isSet = true
    c.bottom = &bottom

    return c
}

// GetLeft
func (c *Margin) GetLeft() int {
    if nil == c.left {
        return 0
    }

    return *c.left
}

// SetLeft 设置左边距
func (c *Margin) SetLeft(left int) *Margin {
    c.isSet = true
    c.left = &left

    return c
}

func (c *Margin) GetXmlBytes() ([]byte, error) {
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
