package cell

import (
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/table"
)

// Cell 表格单元格结构定义
type Cell struct {
    // contents 单元格的内容
    contents []*Content
}

// AddParagraph 添加一个段落
func (c *Cell) AddParagraph() *paragraph.Paragraph {
    content := new(Content)
    content.ct = ContentTypeParagraph
    content.paragraph = new(paragraph.Paragraph)

    c.contents = append(c.contents, content)

    return content.paragraph
}

// AddTable 添加一个表格
func (c *Cell) AddTable() *table.Table {
    content := new(Content)
    content.ct = ContentTypeTable
    content.table = new(table.Table)

    c.contents = append(c.contents, content)

    return content.table
}

type Content struct {
    // 内容类型
    ct ContentType

    // 段落
    paragraph *paragraph.Paragraph

    // 表格
    table *table.Table
}

type ContentType string

const (
    ContentTypeParagraph = ContentType("paragraph")
    ContentTypeTable     = ContentType("table")
)
