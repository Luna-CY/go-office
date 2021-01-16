package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (p *Paragraph) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)
    buffer.WriteString(template.ParagraphStart)

    if nil != p.ppr {
        body, err := p.ppr.GetStyleXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != p.rpr {
        buffer.WriteString(template.RunRPrStart)
        body, err := p.rpr.GetStyleXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
        buffer.WriteString(template.RunRPrEnd)
    }

    for _, r := range p.GetRuns() {
        body, err := r.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.ParagraphEnd)

    return buffer.Bytes(), nil
}
