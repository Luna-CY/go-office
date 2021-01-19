package paragraph

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

func (p *Paragraph) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)
    buffer.WriteString(template.ParagraphStart)

    body, err := p.GetProperties().GetInnerXmlBytes()
    if nil != err {
        return nil, err
    }

    buffer.Write(body)

    body, err = p.GetRunProperties().GetInnerXmlBytes()
    if nil != err {
        return nil, err
    }

    buffer.Write(body)

    for _, r := range p.GetRuns() {
        body, err := r.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.ParagraphEnd)

    empty := fmt.Sprintf(`%v%v`, template.ParagraphStart, template.ParagraphEnd)
    if buffer.String() == empty {
        return []byte(template.ParagraphSingle), nil
    }

    return buffer.Bytes(), nil
}
