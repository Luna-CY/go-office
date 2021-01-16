package paragraph

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

func (p *PPr) GetStyleXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.ParagraphPPrStyleTag)
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrStyleVal, p.GetId()))
    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

func (p *PPr) GetXmlBytes() ([]byte, error) {
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
