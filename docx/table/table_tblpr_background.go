package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// Background 背景配置结构
type Background struct {
    // isSet 是否设置背景
    isSet bool

    // fill 16进制颜色值，指定背景色
    fill *string

    // color 16进制颜色值，指定前景色(文字颜色)
    // 可选值: auto
    color *string

    // val 前景色蒙版的值
    // 参考文档 http://officeopenxml.com/WPtableShading.php
    val *string
}

// GetBackgroundColor 获取背景色
func (b *Background) GetBackgroundColor() string {
    return *b.fill
}

// SetBackgroundColor 设置背景色，不包含#号
func (b *Background) SetBackgroundColor(color string) *Background {
    b.isSet = true
    b.fill = &color

    return b
}

// GetColor 获取前景色
func (b *Background) GetColor() string {
    return *b.color
}

// SetColor 设置前景色，不包含#号
func (b *Background) SetColor(color string) *Background {
    b.isSet = true
    b.color = &color

    return b
}

// SetVal
func (b *Background) SetVal(val string) *Background {
    b.isSet = true
    b.val = &val

    return b
}

func (b *Background) GetXmlBytes() ([]byte, error) {
    if !b.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TblPrBackgroundTag)

    if nil != b.fill {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrBackgroundFill, *b.fill))

        if nil == b.color {
            color := "auto"
            b.color = &color
        }

        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrBackgroundColor, *b.color))
    }

    if nil != b.val {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrVal, *b.val))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}
