package footer

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (f *Footer) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.FooterStart)

    f.cm.RLock()
    for _, content := range f.contents {
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
    f.cm.RUnlock()

    buffer.WriteString(template.FooterEnd)

    return buffer.Bytes(), nil
}
