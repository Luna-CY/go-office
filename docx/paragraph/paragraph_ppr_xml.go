package paragraph

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

func (p *PPr) GetStyleXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.ParagraphPPrStart)

    buffer.WriteByte('<')
    buffer.WriteString(template.ParagraphPPrStyleTag)
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrStyleVal, p.GetId()))
    buffer.WriteString("/>")

    buffer.WriteString(template.ParagraphPPrEnd)

    return buffer.Bytes(), nil
}

// GetInnerXmlBytes 获取内联样式
func (p *PPr) GetInnerXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    if nil != p.sect {
        body, err := p.sect.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if p.keepLines {
        buffer.WriteString(template.ParagraphPPrKeepLines)
    }

    if p.keepNext {
        buffer.WriteString(template.ParagraphPPrKeepNext)
    }

    return buffer.Bytes(), nil
}

// GetExtraXmlBytes 获取外部样式
func (p *PPr) GetExtraXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.ParagraphPPrStart)

    if nil != p.horizontalAlignment {
        buffer.WriteString(strings.Replace(template.ParagraphPPrHorizontalAlignment, "{{TYPE}}", string(*p.horizontalAlignment), 1))
    }

    if nil != p.identity {
        body, err := p.identity.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.spacing {
        body, err := p.spacing.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.borderManager {
        body, err := p.borderManager.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.background {
        body, err := p.background.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.ParagraphPPrEnd)

    return buffer.Bytes(), nil
}

func (p *PPr) GetDefaultXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.ParagraphPPrStart)

    if nil != p.horizontalAlignment {
        buffer.WriteString(strings.Replace(template.ParagraphPPrHorizontalAlignment, "{{TYPE}}", string(*p.horizontalAlignment), 1))
    }

    if nil != p.borderManager {
        body, err := p.borderManager.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.identity {
        body, err := p.identity.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.background {
        body, err := p.background.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.spacing {
        body, err := p.spacing.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.sect {
        body, err := p.sect.GetXmlBytes()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if p.keepLines {
        buffer.WriteString(template.ParagraphPPrKeepLines)
    }

    if p.keepNext {
        buffer.WriteString(template.ParagraphPPrKeepNext)
    }

    buffer.WriteString(template.ParagraphPPrEnd)

    return buffer.Bytes(), nil
}
