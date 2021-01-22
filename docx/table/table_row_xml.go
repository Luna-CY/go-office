package table

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (r *Row) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.TableRowTagStart)

    body, err := r.GetProperties().GetXmlBytes()
    if nil != err {
        return nil, err
    }

    buffer.Write(body)

    for _, cell := range r.GetCells() {
        body, err := cell.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.TableRowTagEnd)

    return buffer.Bytes(), nil
}
