package docx

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/run"
    "github.com/Luna-CY/go-office/docx/template"
)

// StyleConfig 样式配置结构
type StyleConfig struct {
    // pPrDefault 段落的默认样式
    pPrDefault *paragraph.PPr

    // rPrDefault 文本的默认样式
    rPrDefault *run.RPr

    styleList []*Style
}

func (s *StyleConfig) GetDefaultParagraphProperties() *paragraph.PPr {
    if nil == s.pPrDefault {
        s.pPrDefault = new(paragraph.PPr)
    }

    return s.pPrDefault
}

func (s *StyleConfig) GetDefaultRunProperties() *run.RPr {
    if nil == s.rPrDefault {
        s.rPrDefault = new(run.RPr)
    }

    return s.rPrDefault
}

// AddStyle 添加一个样式结构
func (s *StyleConfig) AddStyle(styleId string, styleType StyleType, pPr *paragraph.PPr, rPr *run.RPr) {
    style := &Style{styleId: styleId, styleType: styleType, pPr: pPr, rPr: rPr}

    s.styleList = append(s.styleList, style)
}

func (s *StyleConfig) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.StyleXmlStart)

    // 输出全局默认样式docDefaults
    if nil != s.pPrDefault || nil != s.rPrDefault {
        buffer.WriteString(template.StyleDocDefaultStart)

        // 段落默认样式
        if nil != s.pPrDefault {
            buffer.WriteString(template.StyleDefaultParagraphStart)

            body, err := s.pPrDefault.GetXmlBytes()
            if nil != err {
                return nil, err
            }
            buffer.Write(body)

            buffer.WriteString(template.StyleDefaultParagraphEnd)
        }

        // 文本默认样式
        if nil != s.rPrDefault {
            buffer.WriteString(template.StyleDefaultRunStart)

            body, err := s.rPrDefault.GetXmlBytes()
            if nil != err {
                return nil, err
            }
            buffer.Write(body)

            buffer.WriteString(template.StyleDefaultRunEnd)
        }

        buffer.WriteString(template.StyleDocDefaultEnd)
    }

    for _, style := range s.styleList {
        body, err := style.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.StyleXmlEnd)

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

    // 文本样式属性
    rPr *run.RPr
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

//SetRPr 设置文本样式
func (s *Style) SetRPr(rPr *run.RPr) *Style {
    s.rPr = rPr

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

    if nil != s.rPr {
        body, err := s.rPr.GetXmlBytes()
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
    StyleTypeCharacter = StyleType("character")
)
