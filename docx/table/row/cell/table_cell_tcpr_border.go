package cell

import (
	"bytes"
	"fmt"
	"github.com/Luna-CY/go-office/docx/template"
)

// BorderManager 边框管理器
// TODO: 未实现 tl2br(左上到右下斜线)以及tr2bl(左下到右上斜线)两个样式
type BorderManager struct {
	// isSet 是否设置了边框
	isSet bool

	// top 上边框
	top *Border

	// right 右边框
	right *Border

	// bottom 下边框
	bottom *Border

	// left 左边框
	left *Border

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
	b.right = border
	b.bottom = border
	b.left = border
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
	b.right = border

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
	b.left = border

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

	buffer.WriteString(template.TableCellTcPrBorderStart)

	if nil != b.top {
		body, err := b.getBorderBody(template.TableCellTcPrBorderTopTag, b.top)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != b.right {
		body, err := b.getBorderBody(template.TableCellTcPrBorderRightTag, b.right)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != b.bottom {
		body, err := b.getBorderBody(template.TableCellTcPrBorderBottomTag, b.bottom)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != b.left {
		body, err := b.getBorderBody(template.TableCellTcPrBorderLeftTag, b.left)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != b.insideH {
		body, err := b.getBorderBody(template.TableCellTcPrBorderInsideHTag, b.insideH)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != b.insideV {
		body, err := b.getBorderBody(template.TableCellTcPrBorderInsideVTag, b.insideV)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	buffer.WriteString(template.TableCellTcPrBorderEnd)

	return buffer.Bytes(), nil
}

func (b *BorderManager) getBorderBody(tag string, border *Border) ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteByte('<')
	buffer.WriteString(tag)

	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableCellTcPrBorderStyle, border.Style))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableCellTcPrBorderSize, border.Size))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableCellTcPrBorderColor, border.Color))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableCellTcPrBorderSpace, border.Space))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableCellTcPrBorderShadow, border.Shadow))

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
