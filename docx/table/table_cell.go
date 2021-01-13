package table

import (
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/table/row/cell"
    "sync"
)

// Cell 表格单元格结构定义
type Cell struct {
    pr *cell.TcPr

    cm sync.RWMutex
    // contents 单元格的内容
    contents []*Content
}

// GetProperties 获取属性配置结构
func (c *Cell) GetProperties() *cell.TcPr {
    if nil == c.pr {
        c.pr = new(cell.TcPr)
    }

    return c.pr
}

// GetContents 获取所有内容列表
func (c *Cell) GetContents() []*Content {
    c.cm.RLock()
    defer c.cm.RUnlock()

    return c.contents
}

// AddParagraph 添加一个段落
func (c *Cell) AddParagraph() *paragraph.Paragraph {
    content := new(Content)
    content.ct = ContentTypeParagraph
    content.paragraph = new(paragraph.Paragraph)

    c.cm.Lock()
    c.contents = append(c.contents, content)
    c.cm.Unlock()

    return content.paragraph
}

// AddTable 添加一个表格
func (c *Cell) AddTable() *Table {
    content := new(Content)
    content.ct = ContentTypeTable
    content.table = new(Table)

    c.cm.Lock()
    c.contents = append(c.contents, content)
    c.cm.Unlock()

    return content.table
}

type Content struct {
    // 内容类型
    ct ContentType

    // 段落
    paragraph *paragraph.Paragraph

    // 表格
    table *Table
}

type ContentType string

const (
    ContentTypeParagraph = ContentType("paragraph")
    ContentTypeTable     = ContentType("table")
)
