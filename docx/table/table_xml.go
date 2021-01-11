package table

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (t *Table) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableTag)
    buffer.WriteByte('>')

    // w:tblPr
    if nil != t.tblPr {
        body, err := t.tblPr.GetXmlBytes()
        if nil != err {
            return nil, err
        }
        buffer.Write(body)
    }

    // w:tblGrid
    buffer.WriteByte('<')
    buffer.WriteString(template.TableGridTag)
    buffer.WriteByte('>')

    for _, grid := range t.gridCols {
        body, err := grid.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.TableGridTag)
    buffer.WriteByte('>')

    // w:tr
    for _, row := range t.rows {
        body, err := row.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.TableTag)
    buffer.WriteByte('>')

    return buffer.Bytes(), nil
}
