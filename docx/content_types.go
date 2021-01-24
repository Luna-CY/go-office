package docx

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "sync"
)

// ContentTypes [Content_Types].xml
type ContentTypes struct {
    fm sync.RWMutex
    // files 将要生成的文件列表
    files []contentType
}

func (c *ContentTypes) AddContentType(path string, ct ContentTypeType) *ContentTypes {
    c.fm.Lock()
    defer c.fm.Unlock()

    c.files = append(c.files, contentType{path: path, typeType: ct})

    return c
}

func (c *ContentTypes) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.ContentTypesStart)

    buffer.WriteString(`<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>`)
    buffer.WriteString(`<Default Extension="xml" ContentType="application/xml"/>`)

    c.fm.RLock()
    for _, ct := range c.files {
        buffer.WriteString(fmt.Sprintf(`<%v %v="%v" %v="%v"/>`, template.ContentTypesOverrideTag, template.ContentTypesOverridePartName, ct.path, template.ContentTypesOverrideContentType, ct.typeType))
    }
    c.fm.RUnlock()

    buffer.WriteString(template.ContentTypesEnd)

    return buffer.Bytes(), nil
}

// contentType 内容类型定义
type contentType struct {

    // path 文件路径
    path string

    // typeType 类型
    typeType ContentTypeType
}

type ContentTypeType string

const (
    ContentTypeTypeMain   = ContentTypeType("application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml")
    ContentTypeTypeStyle  = ContentTypeType("application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml")
    ContentTypeTypeCore   = ContentTypeType("application/vnd.openxmlformats-package.core-properties+xml")
    ContentTypeTypeApp    = ContentTypeType("application/vnd.openxmlformats-officedocument.extended-properties+xml")
    ContentTypeTypeHeader = ContentTypeType("application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")
    ContentTypeTypeFooter = ContentTypeType("application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")
)
