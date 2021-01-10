package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (p *Paragraph) GetXmlBytes() ([]byte, error) {
    runBuffer := new(bytes.Buffer)

    for _, r := range p.runs {
        body, err := r.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        runBuffer.Write(body)
    }

    bodyBuffer := new(bytes.Buffer)
    bodyBuffer.WriteString(template.ParagraphStart)

    if nil != p.ppr {
        bodyBuffer.WriteString(template.ParagraphPPrStart)
        body, err := p.ppr.GetStyleXmlBytes()
        if nil != err {
            return nil, err
        }

        bodyBuffer.Write(body)
        bodyBuffer.WriteString(template.ParagraphPPrEnd)
    }

    bodyBuffer.Write(runBuffer.Bytes())
    bodyBuffer.WriteString(template.ParagraphEnd)

    return bodyBuffer.Bytes(), nil
}
