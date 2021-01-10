package table

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (c *Cell) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableRowCellTag)
    buffer.WriteByte('>')

    for _, content := range c.contents {
        if CellContentTypeParagraph == content.ct {
            body, err := content.paragraph.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }

        if CellContentTypeTable == content.ct {
            body, err := content.table.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.TableRowCellTag)
    buffer.WriteByte('>')

    return buffer.Bytes(), nil
}
