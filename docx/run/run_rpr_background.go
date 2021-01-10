package run

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

type Background struct {
    // isSet 是否设置背景
    isSet bool

    // fill 16进制颜色值，指定背景色
    fill *string

    // color 16进制颜色值，指定前景色
    color *string

    // val 前景色蒙版的值
    // 参考文档 http://officeopenxml.com/WPtextShading.php
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

// GetFrontColor 获取前景色
func (b *Background) GetFrontColor() string {
    return *b.color
}

// SetFrontColor 设置前景色，不包含#号
func (b *Background) SetFrontColor(color string) *Background {
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
    buffer.WriteString(template.RunPPrBackgroundTag)

    if nil != b.fill {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunPPrBackgroundFill, *b.fill))
    }

    if nil != b.color {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunPPrBackgroundColor, *b.color))
    }

    if nil != b.val {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunRPrVal, *b.val))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}
