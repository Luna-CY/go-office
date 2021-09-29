package section

import (
	"bytes"
	"fmt"
	"github.com/Luna-CY/go-office/docx/template"
)

// PageBorder 页面边框配置结构
type PageBorder struct {
	// isSet 是否设置了页面边框
	isSet bool

	// display 页面边框的显示设置
	display *PageBorderDisplay

	// offsetFrom 边框相对的元素
	offsetFrom *PageBorderOffsetFrom

	// zOrder 边距的显示排序设置
	zOrder *PageBorderZOrder

	// top 上边框配置
	top *Border

	// right 右边框配置
	right *Border

	// bottom 下边框配置
	bottom *Border

	// left 左边框配置
	left *Border
}

// GetDisplay 获取显示配置
func (p *PageBorder) GetDisplay() PageBorderDisplay {
	return *p.display
}

// SetDisplay 设置显示设置
func (p *PageBorder) SetDisplay(display PageBorderDisplay) *PageBorder {
	p.isSet = true
	p.display = &display

	return p
}

// GetOffsetFrom 获取偏移配置
func (p *PageBorder) GetOffsetFrom() PageBorderOffsetFrom {
	return *p.offsetFrom
}

// SetOffsetFrom 设置边框相对于哪个元素进行偏移
func (p *PageBorder) SetOffsetFrom(offsetFrom PageBorderOffsetFrom) *PageBorder {
	p.isSet = true
	p.offsetFrom = &offsetFrom

	return p
}

// GetZOrder 获取显示排序
func (p *PageBorder) GetZOrder() PageBorderZOrder {
	return *p.zOrder
}

// SetZOrder 设置显示排序
func (p *PageBorder) SetZOrder(zOrder PageBorderZOrder) *PageBorder {
	p.isSet = true
	p.zOrder = &zOrder

	return p
}

// SetBorder 统一设置边框
func (p *PageBorder) SetBorder(style PageBorderStyle, color string, size uint8, space uint, shadow bool) *PageBorder {
	border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

	p.isSet = true
	p.top = border
	p.right = border
	p.bottom = border
	p.left = border

	return p
}

func (p *PageBorder) GetTop() *Border {
	return p.top
}

// SetTop 设置上边框
func (p *PageBorder) SetTop(style PageBorderStyle, color string, size uint8, space uint, shadow bool) *PageBorder {
	border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

	p.isSet = true
	p.top = border

	return p
}

func (p *PageBorder) GetRight() *Border {
	return p.top
}

// SetRight 设置右边框
func (p *PageBorder) SetRight(style PageBorderStyle, color string, size uint8, space uint, shadow bool) *PageBorder {
	border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

	p.isSet = true
	p.top = border

	return p
}

func (p *PageBorder) GetBottom() *Border {
	return p.top
}

// SetBottom 设置下边框
func (p *PageBorder) SetBottom(style PageBorderStyle, color string, size uint8, space uint, shadow bool) *PageBorder {
	border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

	p.isSet = true
	p.top = border

	return p
}

func (p *PageBorder) GetLeft() *Border {
	return p.top
}

// SetLeft 设置左边框
func (p *PageBorder) SetLeft(style PageBorderStyle, color string, size uint8, space uint, shadow bool) *PageBorder {
	border := &Border{Style: style, Color: color, Size: size, Space: space, Shadow: shadow}

	p.isSet = true
	p.top = border

	return p
}

func (p *PageBorder) GetXmlBytes() ([]byte, error) {
	if !p.isSet {
		return []byte{}, nil
	}

	buffer := new(bytes.Buffer)

	buffer.WriteByte('<')
	buffer.WriteString(template.SectionPageBorderTag)

	if nil != p.display {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageBorderDisplay, *p.display))
	}

	if nil != p.offsetFrom {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageBorderOffsetFrom, *p.offsetFrom))
	}

	if nil != p.zOrder {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageBorderZOrder, *p.zOrder))
	}
	buffer.WriteByte('>')

	if nil != p.top {
		body, err := p.getBorderBody(template.SectionPageBorderTop, p.top)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != p.right {
		body, err := p.getBorderBody(template.SectionPageBorderRight, p.right)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != p.bottom {
		body, err := p.getBorderBody(template.SectionPageBorderBottom, p.bottom)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	if nil != p.left {
		body, err := p.getBorderBody(template.SectionPageBorderLeft, p.left)
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	buffer.Write([]byte{'<', '/'})
	buffer.WriteString(template.SectionPageBorderTag)
	buffer.WriteByte('>')

	return buffer.Bytes(), nil
}

func (p *PageBorder) getBorderBody(tag string, border *Border) ([]byte, error) {
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

type PageBorderDisplay string

const (
	PageBorderDisplayAll      = PageBorderDisplay("allPages")
	PageBorderDisplayFirst    = PageBorderDisplay("firstPage")
	PageBorderDisplayNotFirst = PageBorderDisplay("notFirstPage")
)

type PageBorderOffsetFrom string

const (
	PageBorderOffsetFromPage = PageBorderOffsetFrom("page")
	PageBorderOffsetFromText = PageBorderOffsetFrom("text")
)

type PageBorderZOrder string

const (
	PageBorderZOrderFront = PageBorderZOrder("front")
	PageBorderZOrderBack  = PageBorderZOrder("back")
)

type Border struct {
	// id 边框图像的id
	// TODO: 该属性未实现
	id string

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
	Style PageBorderStyle
}

type PageBorderStyle string

const (
	PageBorderStyleNil                    = PageBorderStyle("nil")
	PageBorderStyleNone                   = PageBorderStyle("none")
	PageBorderStyleSingle                 = PageBorderStyle("single")
	PageBorderStyleDashDotStroked         = PageBorderStyle("dashDotStroked")
	PageBorderStyleDashed                 = PageBorderStyle("dashed")
	PageBorderStyleDashSmallGap           = PageBorderStyle("dashSmallGap")
	PageBorderStyleDotDash                = PageBorderStyle("dotDash")
	PageBorderStyleDotDotDash             = PageBorderStyle("dotDotDash")
	PageBorderStyleDotted                 = PageBorderStyle("dotted")
	PageBorderStyleDouble                 = PageBorderStyle("double")
	PageBorderStyleDoubleWave             = PageBorderStyle("doubleWave")
	PageBorderStyleInset                  = PageBorderStyle("inset")
	PageBorderStyleOutset                 = PageBorderStyle("outset")
	PageBorderStyleThick                  = PageBorderStyle("thick")
	PageBorderStyleThickThinLargeGap      = PageBorderStyle("thickThinLargeGap")
	PageBorderStyleThickThinMediumGap     = PageBorderStyle("thickThinMediumGap")
	PageBorderStyleThickThinSmallGap      = PageBorderStyle("thickThinSmallGap")
	PageBorderStyleThinThickLargeGap      = PageBorderStyle("thinThickLargeGap")
	PageBorderStyleThinThickMediumGap     = PageBorderStyle("thinThickMediumGap")
	PageBorderStyleThinThickSmallGap      = PageBorderStyle("thinThickSmallGap")
	PageBorderStyleThinThickThinLargeGap  = PageBorderStyle("thinThickThinLargeGap")
	PageBorderStyleThinThickThinMediumGap = PageBorderStyle("thinThickThinMediumGap")
	PageBorderStyleThinThickThinSmallGap  = PageBorderStyle("thinThickThinSmallGap")
	PageBorderStyleThreeDEmboss           = PageBorderStyle("threeDEmboss")
	PageBorderStyleThreeDEngrave          = PageBorderStyle("threeDEngrave")
	PageBorderStyleTriple                 = PageBorderStyle("triple")
	PageBorderStyleWave                   = PageBorderStyle("wave")
)
