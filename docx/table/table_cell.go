package table

import "github.com/Luna-CY/go-office/docx/paragraph"

// Cell 表格单元格结构定义
type Cell struct {
    // contents 单元格的内容
    contents []*CellContent
}

// AddParagraph 添加一个段落
func (c *Cell) AddParagraph() *paragraph.Paragraph {
    content := new(CellContent)
    content.ct = CellContentTypeParagraph
    content.paragraph = new(paragraph.Paragraph)

    c.contents = append(c.contents, content)

    return content.paragraph
}

// AddTable 添加一个表格
func (c *Cell) AddTable() *Table {
    content := new(CellContent)
    content.ct = CellContentTypeTable
    content.table = new(Table)

    c.contents = append(c.contents, content)

    return content.table
}

type CellContent struct {
    // 内容类型
    ct CellContentType

    // 段落
    paragraph *paragraph.Paragraph

    // 表格
    table *Table
}

type CellContentType string

const (
    CellContentTypeParagraph = CellContentType("paragraph")
    CellContentTypeTable     = CellContentType("table")
)
