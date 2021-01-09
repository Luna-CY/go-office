package docx

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// themeTint 文档的背景配置
type Background struct {
    // isSet 是否设置了背景
    isSet bool

    // color 背景色，16进制颜色值
    color string

    // themeColor 主题颜色的id，其指向的id在主题中定义
    themeColor string

    // themeShade 主题阴影颜色，0-255的16进制值
    themeShade string

    // themeTint 主题色调值，0-255的16进制值
    themeTint string
}

// SetBackground 设置文档背景
func (b *Background) SetBackground(color, themeColor, themeShade, themeTint string) *Background {
    b.isSet = true

    b.color = color
    b.themeColor = themeColor
    b.themeShade = themeShade
    b.themeTint = themeTint

    return b
}

// SetColor 设置颜色
func (b *Background) SetColor(color string) *Background {
    b.isSet = true
    b.color = color

    return b
}

// SetThemeColor 设置主题颜色
func (b *Background) SetThemeColor(themeColor string) *Background {
    b.isSet = true
    b.themeColor = themeColor

    return b
}

// SetThemeShade 设置主题阴影
func (b *Background) SetThemeShade(themeShade string) *Background {
    b.isSet = true
    b.themeShade = themeShade

    return b
}

// SetThemeTint 设置主题的色调值
// 0-255的16进制字符串
func (b *Background) SetThemeTint(themeTint string) *Background {
    b.isSet = true
    b.themeTint = themeTint

    return b
}

func (b *Background) GetBody() ([]byte, error) {
    if !b.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.BackgroundTag)

    if "" != b.color {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.BackgroundColor, b.color))
    }

    if "" != b.themeColor {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.BackgroundThemeColor, b.themeColor))
    }

    if "" != b.themeShade {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.BackgroundThemeShade, b.themeShade))
    }

    if "" != b.themeTint {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.BackgroundThemeTint, b.themeTint))
    }

    buffer.WriteString(" />")

    return buffer.Bytes(), nil
}
