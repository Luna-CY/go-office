package header

import (
    "fmt"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/table"
    "sync"
)

// TableDefaultWidth 表格默认宽度
const TableDefaultWidth = 8800

// Header 页头定义
type Header struct {
    // fileName 该header的文件名
    fileName string

    // rId 关联的ID
    rId string

    cm sync.RWMutex
    // contents 文档的内容列表
    contents []*DocumentContent
}

// SetRId 设置此页头的rId
func (h *Header) SetRId(rId string) *Header {
    h.rId = rId

    return h
}

// GetRId 获取该页头的rId
func (h *Header) GetRId() string {
    return h.rId
}

// GetFileName 获取文件名称
func (h *Header) GetFileName() string {
    if "" == h.fileName {
        return h.GenerateFileName()
    }

    return h.fileName
}

// GenerateFileName 生成文件名
func (h *Header) GenerateFileName() string {
    h.cm.RLock()
    defer h.cm.RUnlock()

    h.fileName = fmt.Sprintf("header%d.xml", len(h.contents)+1)

    return h.fileName
}

// AddParagraph 添加一个段落
func (h *Header) AddParagraph() *paragraph.Paragraph {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeParagraph

    content.paragraph = new(paragraph.Paragraph)
    h.cm.Lock()
    h.contents = append(h.contents, content)
    h.cm.Unlock()

    return content.paragraph
}

// AddTable 添加一个表格
func (h *Header) AddTable() *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)

    h.cm.Lock()
    h.contents = append(h.contents, content)
    h.cm.Unlock()

    return content.table
}

// AddTableWithColumns 添加一个拥有指定列数量的表格
func (h *Header) AddTableWithColumns(columns int) *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)

    for i := 0; i < columns; i++ {
        content.table.AddCol()
    }

    h.cm.Lock()
    h.contents = append(h.contents, content)
    h.cm.Unlock()

    return content.table
}

// AddTableWithColumnsAndAutoWidth 添加一个拥有指定列数量的表格，并且自动计算所有列的宽度
func (h *Header) AddTableWithColumnsAndAutoWidth(columns int) *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)

    for i := 0; i < columns; i++ {
        content.table.AddColWithWidth(TableDefaultWidth / columns)
    }

    h.cm.Lock()
    h.contents = append(h.contents, content)
    h.cm.Unlock()

    return content.table
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
