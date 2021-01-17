package docx

import (
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/section"
    "github.com/Luna-CY/go-office/docx/table"
    "sync"
)

// TableDefaultWidth 表格默认宽度
const TableDefaultWidth = DocumentDefaultWidth - 3000

// Document 文档结构定义
// 该结构统一对.docx格式文档进行操作
type Document struct {
    // core core.xml
    core Core

    // app app.xml
    app App

    // style style.xml
    style StyleConfig

    // background 文档的背景配置
    background Background

    cm sync.RWMutex
    // contents 文档的内容列表
    contents []*DocumentContent

    // section 文档的节属性配置
    section section.Section
}

func (d *Document) GetProperties() *StyleConfig {
    return &d.style
}

func (d *Document) GetContents() []*DocumentContent {
    d.cm.RLock()
    defer d.cm.RUnlock()

    return d.contents
}

// AddParagraph 添加一个段落
func (d *Document) AddParagraph() *paragraph.Paragraph {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeParagraph

    content.paragraph = new(paragraph.Paragraph)
    d.cm.Lock()
    d.contents = append(d.contents, content)
    d.cm.Unlock()

    return content.paragraph
}

// AddTable 添加一个表格
func (d *Document) AddTable() *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)
    // 设置默认宽度
    content.table.GetProperties().SetWidth(TableDefaultWidth)

    d.cm.Lock()
    d.contents = append(d.contents, content)
    d.cm.Unlock()

    return content.table
}

// AddTableWithColumns 添加一个拥有指定列数量的表格
func (d *Document) AddTableWithColumns(columns int) *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)
    // 设置默认宽度
    content.table.GetProperties().SetWidth(TableDefaultWidth)

    for i := 0; i < columns; i++ {
        content.table.AddColWithWidth(TableDefaultWidth / columns)
    }

    d.cm.Lock()
    d.contents = append(d.contents, content)
    d.cm.Unlock()

    return content.table
}

// GetSection 获取节属性配置结构指针
func (d *Document) GetSection() *section.Section {
    return &d.section
}

// GetBackground 获取背景配置结构指针
func (d *Document) GetBackground() *Background {
    return &d.background
}

// DocumentContent 文档内容
type DocumentContent struct {
    // ct 内容类型
    ct DocumentContentType

    // paragraph 段落
    paragraph *paragraph.Paragraph

    // table 表格
    table *table.Table
}

type DocumentContentType string

const (
    DocumentContentTypeParagraph = DocumentContentType("paragraph")
    DocumentContentTypeTable     = DocumentContentType("table")
)
