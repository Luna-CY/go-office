package header

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (h *Header) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.HeaderStart)

    h.cm.RLock()
    for _, content := range h.contents {
        if DocumentContentTypeParagraph == content.ct {
            body, err := content.paragraph.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }

        if DocumentContentTypeTable == content.ct {
            body, err := content.table.GetXmlBytes()
            if nil != err {
                return nil, err
            }

            buffer.Write(body)
        }
    }
    h.cm.RUnlock()

    buffer.WriteString(template.HeaderEnd)

    return buffer.Bytes(), nil
}
