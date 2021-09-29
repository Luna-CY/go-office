package table

import (
	"bytes"
	"github.com/Luna-CY/go-office/docx/template"
)

func (t *Table) GetXmlBytes() ([]byte, error) {
	if 0 == len(t.GetCols()) || 0 == len(t.GetRows()) {
		return []byte{}, nil
	}

	buffer := new(bytes.Buffer)

	buffer.WriteString(template.TableTagStart)

	// w:tblPr
	body, err := t.GetProperties().GetInnerXmlBytes()
	if nil != err {
		return nil, err
	}
	buffer.Write(body)

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

	buffer.WriteString(template.TableTagEnd)

	return buffer.Bytes(), nil
}
