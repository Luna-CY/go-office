package paragraph

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

    // mask 前景色蒙版的值
    // 参考文档 http://officeopenxml.com/WPshading.php
    mask *MaskValue
}

// GetBackgroundColor 获取背景色
func (b *Background) GetBackgroundColor() string {
    if nil == b.fill {
        return ""
    }

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
    if nil == b.color {
        return ""
    }

    return *b.color
}

// SetColor 设置前景色，不包含#号
func (b *Background) SetColor(color string) *Background {
    b.isSet = true
    b.color = &color

    return b
}

// SetMask
func (b *Background) SetMask(val MaskValue) *Background {
    b.isSet = true
    b.mask = &val

    return b
}

func (b *Background) GetXmlBytes() ([]byte, error) {
    if !b.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.ParagraphPPrBackgroundTag)

    if nil != b.fill {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBackgroundFill, *b.fill))

        if nil == b.color {
            color := "auto"
            b.color = &color
        }

        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBackgroundColor, *b.color))
    }

    if nil != b.mask {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBackgroundVal, *b.mask))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

type MaskValue string

const (
    MaskValueClear = MaskValue("clear")
    MaskValueNil   = MaskValue("nil")
)
