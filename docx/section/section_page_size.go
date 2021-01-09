package section

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// PageSize 页面大小配置结构
type PageSize struct {
    // isSet 是否设置了页面大小
    isSet bool

    // height 页面高度
    // 以点的二十分之一为单位
    height *uint

    // width 页面宽度
    // 以点的二十分之一为单位
    width *uint

    // 页面方向
    orient *PageSizeOrient
}

// SetSize 设置页面大小
func (p *PageSize) SetSize(width, height uint) *PageSize {
    p.isSet = true

    p.width = &width
    p.height = &height

    return p
}

// SetWidth 设置页面宽度
func (p *PageSize) SetWidth(width uint) *PageSize {
    p.isSet = true
    p.width = &width

    return p
}

// SetHeight 设置页面高度
func (p *PageSize) SetHeight(height uint) *PageSize {
    p.isSet = true
    p.height = &height

    return p
}

// SetOrient 设置页面方向
func (p *PageSize) SetOrient(orient PageSizeOrient) *PageSize {
    p.isSet = true
    p.orient = &orient

    return p
}

func (p *PageSize) GetBody() ([]byte, error) {
    if !p.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.SectionPageSizeTag)

    if nil != p.width {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageSizeWidth, *p.width))
    }

    if nil != p.height {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageSizeHeight, *p.height))
    }

    if nil != p.orient {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageSizeOrient, *p.orient))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

type PageSizeOrient string

const (
    // PageSizeOrientLandscape 横向
    PageSizeOrientLandscape = PageSizeOrient("landscape")

    // PageSizeOrientPortrait 纵向
    PageSizeOrientPortrait = PageSizeOrient("portrait")
)
