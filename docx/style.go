package docx

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/template"
)

// StyleConfig 样式配置结构
type StyleConfig struct {
    styleList []*Style
}

// AddStyle 添加一个样式结构
func (s *StyleConfig) AddStyle(styleId string, styleType StyleType, pPr *paragraph.PPr) {
    style := &Style{styleId: styleId, styleType: styleType, pPr: pPr}

    s.styleList = append(s.styleList, style)
}

func (s *StyleConfig) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    for _, style := range s.styleList {
        body, err := style.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    return buffer.Bytes(), nil
}

// Style 样式结构
type Style struct {
    // styleType 样式类型
    styleType StyleType

    // styleId 样式ID
    styleId string

    // 段落样式属性
    pPr *paragraph.PPr
}

// SetStyleType 设置样式类型
func (s *Style) SetStyleType(styleType StyleType) *Style {
    s.styleType = styleType

    return s
}

// SetStyleId 设置样式ID
func (s *Style) SetStyleId(styleId string) *Style {
    s.styleId = styleId

    return s
}

// SetPPr 设置段落样式
func (s *Style) SetPPr(pPr *paragraph.PPr) *Style {
    s.pPr = pPr

    return s
}

func (s *Style) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.StyleStyleTag)

    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.StyleStyleType, s.styleType))
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.StyleStyleStyleId, s.styleId))
    buffer.WriteByte('>')

    if nil != s.pPr {
        body, err := s.pPr.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.StyleStyleTag)
    buffer.WriteByte('>')

    return buffer.Bytes(), nil
}

type StyleType string

const (
    StyleTypeParagraph = StyleType("paragraph")
)
