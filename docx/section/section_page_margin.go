package section

import (
	"bytes"
	"fmt"
	"github.com/Luna-CY/go-office/docx/template"
)

// PageMargin 页面边距配置结构
// 所有距离均以二十分之一点为单位
type PageMargin struct {
	// 是否设置了页边距
	isSet bool

	// Bottom 文本距离页面底部的距离，
	Bottom *int

	// Footer 页脚距离页面底部的距离
	Footer *int

	// Left 文本距离页面左侧的距离
	Left *int

	// Right 文本距离页面右侧的距离
	Right *int

	// Header 页眉到页面顶部的距离
	Header *int

	// Top 文本距离页面顶部的距离
	Top *int

	// Gutter 页面装订线的距离
	Gutter *int
}

// SetMargin 联合设置页边距
func (p *PageMargin) SetMargin(top, right, bottom, left int) *PageMargin {
	p.isSet = true

	p.Top = &top
	p.Right = &right
	p.Bottom = &bottom
	p.Left = &left

	return p
}

// SetMarginGroup 分组设置页边距
func (p *PageMargin) SetMarginGroup(tb, rl int) *PageMargin {
	p.isSet = true

	p.Top = &tb
	p.Bottom = &tb
	p.Right = &rl
	p.Left = &rl

	return p
}

// SetBottom 设置底部页边距
func (p *PageMargin) SetBottom(bottom int) *PageMargin {
	p.isSet = true
	p.Bottom = &bottom

	return p
}

// SetFooter 设置页脚的边距
func (p *PageMargin) SetFooter(footer int) *PageMargin {
	p.isSet = true
	p.Footer = &footer

	return p
}

// SetLeft 设置左侧页边距
func (p *PageMargin) SetLeft(left int) *PageMargin {
	p.isSet = true
	p.Left = &left

	return p
}

// SetRight 设置右侧页边距
func (p *PageMargin) SetRight(right int) *PageMargin {
	p.isSet = true
	p.Right = &right

	return p
}

// SetHeader 设置页眉的边距
func (p *PageMargin) SetHeader(header int) *PageMargin {
	p.isSet = true
	p.Header = &header

	return p
}

// SetTop 设置顶部页边距
func (p *PageMargin) SetTop(top int) *PageMargin {
	p.isSet = true
	p.Top = &top

	return p
}

// SetGutter 设置装订线边距
func (p *PageMargin) SetGutter(gutter int) *PageMargin {
	p.isSet = true
	p.Gutter = &gutter

	return p
}

func (p *PageMargin) GetXmlBytes() ([]byte, error) {
	if !p.isSet {
		return []byte{}, nil
	}

	buffer := new(bytes.Buffer)

	buffer.WriteByte('<')
	buffer.WriteString(template.SectionPageMarginTag)

	if nil != p.Header {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginHeader, *p.Header))
	}

	if nil != p.Footer {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginFooter, *p.Footer))
	}

	if nil != p.Bottom {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginBottom, *p.Bottom))
	}

	if nil != p.Top {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginTop, *p.Top))
	}

	if nil != p.Left {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginLeft, *p.Left))
	}

	if nil != p.Right {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginRight, *p.Right))
	}

	if nil != p.Gutter {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageMarginGutter, *p.Gutter))
	}

	buffer.WriteString("/>")

	return buffer.Bytes(), nil
}
