package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// BorderManager 边框管理器
type BorderManager struct {
    // isSet 是否设置了边框
    isSet bool

    // top 上边框
    top *Border

    // end 右边框
    end *Border

    // bottom 下边框
    bottom *Border

    // start 左边框
    start *Border

    // insideH 表格内行间边框的样式
    insideH *Border

    // insideV 表格内列间边框的样式
    insideV *Border
}

// SetBorder 统一设置边框样式
func (b *BorderManager) SetBorder(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.top = border
    b.end = border
    b.bottom = border
    b.start = border
    b.insideH = border
    b.insideV = border

    return b
}

// SetTop 设置顶部边框
func (b *BorderManager) SetTop(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.top = border

    return b
}

// SetRight 设置右侧边框
func (b *BorderManager) SetRight(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.end = border

    return b
}

// SetBottom 设置底部边框
func (b *BorderManager) SetBottom(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.bottom = border

    return b
}

// SetLeft 设置左侧边框
func (b *BorderManager) SetLeft(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.start = border

    return b
}

// SetInsideHorizontal 设置表格内行间边框
func (b *BorderManager) SetInsideHorizontal(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.insideH = border

    return b
}

// SetInsideVertical 设置表格内列间边框
func (b *BorderManager) SetInsideVertical(style BorderStyle, color string, size uint8, space uint, shadow bool) *BorderManager {
    border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

    b.isSet = true
    b.insideV = border

    return b
}

func (b *BorderManager) GetXmlBytes() ([]byte, error) {
    if !b.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteString(template.TblPrBorderStart)

    if nil != b.top {
        body, err := b.getBorderBody(template.TblPrBorderTopTag, b.top)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != b.end {
        body, err := b.getBorderBody(template.TblPrBorderEndTag, b.end)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)

        body, err = b.getBorderBody(template.TblPrBorderRightTag, b.end)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != b.bottom {
        body, err := b.getBorderBody(template.TblPrBorderBottomTag, b.bottom)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != b.start {
        body, err := b.getBorderBody(template.TblPrBorderStartTag, b.start)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)

        body, err = b.getBorderBody(template.TblPrBorderLeftTag, b.start)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != b.insideH {
        body, err := b.getBorderBody(template.TblPrBorderInsideHTag, b.insideH)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != b.insideV {
        body, err := b.getBorderBody(template.TblPrBorderInsideVTag, b.insideV)
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.TblPrBorderEnd)

    return buffer.Bytes(), nil
}

func (b *BorderManager) getBorderBody(tag string, border *Border) ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(tag)

    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBorderStyle, border.Style))
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBorderSize, border.Size))
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBorderColor, border.Color))
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBorderSpace, border.Space))
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrBorderShadow, border.Shadow))

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

// Border 边框结构定义
type Border struct {
    // Color 颜色值，以十六进制字符串设置，不包含#号，例如FF0000
    Color string

    // Shadow 是否显示阴影
    Shadow bool

    // Space 间距偏移量，以磅（1/72英寸）为单位
    Space uint

    // Size 边框宽度
    // 最小值是2，最大值是96
    Size uint8

    // Style 边框样式
    // 段落边框只能是线条边框
    Style BorderStyle
}

type BorderStyle string

const (
    BorderStyleNil                    = BorderStyle("nil")
    BorderStyleNone                   = BorderStyle("none")
    BorderStyleSingle                 = BorderStyle("single")
    BorderStyleDashDotStroked         = BorderStyle("dashDotStroked")
    BorderStyleDashed                 = BorderStyle("dashed")
    BorderStyleDashSmallGap           = BorderStyle("dashSmallGap")
    BorderStyleDotDash                = BorderStyle("dotDash")
    BorderStyleDotDotDash             = BorderStyle("dotDotDash")
    BorderStyleDotted                 = BorderStyle("dotted")
    BorderStyleDouble                 = BorderStyle("double")
    BorderStyleDoubleWave             = BorderStyle("doubleWave")
    BorderStyleInset                  = BorderStyle("inset")
    BorderStyleOutset                 = BorderStyle("outset")
    BorderStyleThick                  = BorderStyle("thick")
    BorderStyleThickThinLargeGap      = BorderStyle("thickThinLargeGap")
    BorderStyleThickThinMediumGap     = BorderStyle("thickThinMediumGap")
    BorderStyleThickThinSmallGap      = BorderStyle("thickThinSmallGap")
    BorderStyleThinThickLargeGap      = BorderStyle("thinThickLargeGap")
    BorderStyleThinThickMediumGap     = BorderStyle("thinThickMediumGap")
    BorderStyleThinThickSmallGap      = BorderStyle("thinThickSmallGap")
    BorderStyleThinThickThinLargeGap  = BorderStyle("thinThickThinLargeGap")
    BorderStyleThinThickThinMediumGap = BorderStyle("thinThickThinMediumGap")
    BorderStyleThinThickThinSmallGap  = BorderStyle("thinThickThinSmallGap")
    BorderStyleThreeDEmboss           = BorderStyle("threeDEmboss")
    BorderStyleThreeDEngrave          = BorderStyle("threeDEngrave")
    BorderStyleTriple                 = BorderStyle("triple")
    BorderStyleWave                   = BorderStyle("wave")
)
