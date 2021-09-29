package table

import (
	"bytes"
	"github.com/Luna-CY/go-office/docx/template"
)

func (c *Cell) GetXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.TableCellTagStart)

	body, err := c.GetProperties().GetXmlBytes()
	if nil != err {
		return nil, err
	}

	buffer.Write(body)

	// 如果是空的内容列表就添加一个默认的空段落
	if 0 == len(c.GetContents()) {
		c.AddParagraph()
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

	buffer.WriteString(template.TableCellTagEnd)

	return buffer.Bytes(), nil
}
