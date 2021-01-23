package docx

import (
    "fmt"
    "github.com/Luna-CY/go-office/docx/header"
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/table"
)

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

    for i := 0; i < columns; i++ {
        content.table.AddCol()
    }

    d.cm.Lock()
    d.contents = append(d.contents, content)
    d.cm.Unlock()

    return content.table
}

// AddTableWithColumnsAndAutoWidth 添加一个拥有指定列数量的表格，并且自动计算所有列的宽度
func (d *Document) AddTableWithColumnsAndAutoWidth(columns int) *table.Table {
    content := new(DocumentContent)
    content.ct = DocumentContentTypeTable

    content.table = new(table.Table)

    for i := 0; i < columns; i++ {
        content.table.AddColWithWidth(TableDefaultWidth / columns)
    }

    d.cm.Lock()
    d.contents = append(d.contents, content)
    d.cm.Unlock()

    return content.table
}

// NewHeader 新建一个Header
func (d *Document) NewHeader() *header.Header {
    hdr := new(header.Header)
    fileName := hdr.GenerateFileName()

    d.hm.Lock()
    defer d.hm.Unlock()

    d.headers = append(d.headers, hdr)

    d.contentTypes.AddContentType(fmt.Sprintf("word/%v", fileName), ContentTypeTypeHeader)
    rId := d.relationship.AddRelationship(fileName, RelationshipTypeHeader)
    hdr.SetRId(rId)

    return hdr
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
