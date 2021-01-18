package table

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (t *Table) GetXmlBytes() ([]byte, error) {
    if nil == t.GetProperties().width {
        t.GetProperties().SetWidth(8520)
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableTag)
    buffer.WriteByte('>')

    // w:tblPr
    if nil != t.tblPr {
        buffer.WriteString(template.TblPrStart)

        body, err := t.tblPr.GetStyleXmlBytes()
        if nil != err {
            return nil, err
        }
        buffer.Write(body)

        body, err = t.tblPr.GetInnerXmlBytes()
        if nil != err {
            return nil, err
        }
        buffer.Write(body)

        buffer.WriteString(template.TblPrEnd)
    }

    // w:tblGrid
    buffer.WriteByte('<')
    buffer.WriteString(template.TableGridTag)
    buffer.WriteByte('>')

    for _, grid := range t.GetCols() {
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
