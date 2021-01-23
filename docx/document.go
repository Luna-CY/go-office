package docx

import (
    "github.com/Luna-CY/go-office/docx/header"
    "github.com/Luna-CY/go-office/docx/section"
    "sync"
)

// TableDefaultWidth 表格默认宽度
const TableDefaultWidth = 8800

// Document 文档结构定义
// 该结构统一对.docx格式文档进行操作
type Document struct {
    // core core.xml
    core Core

    // app app.xml
    app App

    // styles styles.xml
    styles Styles

    // background 文档的背景配置
    background Background

    cm sync.RWMutex
    // contents 文档的内容列表
    contents []*DocumentContent

    hm sync.RWMutex
    // headers 表头列表
    headers []*header.Header

    // section 文档的节属性配置
    section section.Section

    // contentTypes [Content_Types].xml
    contentTypes ContentTypes

    // relationship document.xml.rels
    relationship Relationships
}

func (d *Document) GetProperties() *Styles {
    return &d.styles
}

func (d *Document) GetContents() []*DocumentContent {
    d.cm.RLock()
    defer d.cm.RUnlock()

    return d.contents
}

// GetSection 获取节属性配置结构指针
func (d *Document) GetSection() *section.Section {
    return &d.section
}

// GetBackground 获取背景配置结构指针
func (d *Document) GetBackground() *Background {
    return &d.background
}
