package table

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (c *Cell) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableCellTag)
    buffer.WriteByte('>')

    if nil != c.pr {
        body, err := c.pr.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    for _, content := range c.GetContents() {
        if ContentTypeParagraph == content.ct {
            body, err := content.paragraph.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }

        if ContentTypeTable == content.ct {
            body, err := content.table.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.TableCellTag)
    buffer.WriteByte('>')

    return buffer.Bytes(), nil
}
