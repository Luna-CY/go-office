package run

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (r *Run) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.RunStart)

    if nil != r.rpr {
        buffer.WriteString(template.RunRPrStart)

        body, err := r.rpr.GetStyleXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)

        body, err = r.rpr.GetInnerXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)

        buffer.WriteString(template.RunRPrEnd)
    }

    buffer.Write(r.body.Bytes())
    buffer.WriteString(template.RunEnd)

    return buffer.Bytes(), nil
}
