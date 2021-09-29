package footer

import (
	"fmt"
	"github.com/Luna-CY/go-office/docx/paragraph"
	"github.com/Luna-CY/go-office/docx/table"
	"sync"
)

// TableDefaultWidth 表格默认宽度
const TableDefaultWidth = 8800

// Footer 页脚定义
type Footer struct {
	// fileName 该header的文件名
	fileName string

	// rId 关联的ID
	rId string

	cm sync.RWMutex
	// contents 文档的内容列表
	contents []*DocumentContent
}

// SetRId 设置此页脚的rId
func (f *Footer) SetRId(rId string) *Footer {
	f.rId = rId

	return f
}

// GetRId 获取该页脚的rId
func (f *Footer) GetRId() string {
	return f.rId
}

// GetFileName 获取文件名称
func (f *Footer) GetFileName() string {
	if "" == f.fileName {
		return f.GenerateFileName()
	}

	return f.fileName
}

// GenerateFileName 生成文件名
func (f *Footer) GenerateFileName() string {
	f.cm.RLock()
	defer f.cm.RUnlock()

	f.fileName = fmt.Sprintf("footer%d.xml", len(f.contents)+1)

	return f.fileName
}

func (f *Footer) GetContents() []*DocumentContent {
	f.cm.RLock()
	defer f.cm.RUnlock()

	return f.contents
}

// AddParagraph 添加一个段落
func (f *Footer) AddParagraph() *paragraph.Paragraph {
	content := new(DocumentContent)
	content.ct = DocumentContentTypeParagraph

	content.paragraph = new(paragraph.Paragraph)
	f.cm.Lock()
	f.contents = append(f.contents, content)
	f.cm.Unlock()

	return content.paragraph
}

// AddTable 添加一个表格
func (f *Footer) AddTable() *table.Table {
	content := new(DocumentContent)
	content.ct = DocumentContentTypeTable

	content.table = new(table.Table)

	f.cm.Lock()
	f.contents = append(f.contents, content)
	f.cm.Unlock()

	return content.table
}

// AddTableWithColumns 添加一个拥有指定列数量的表格
func (f *Footer) AddTableWithColumns(columns int) *table.Table {
	content := new(DocumentContent)
	content.ct = DocumentContentTypeTable

	content.table = new(table.Table)

	for i := 0; i < columns; i++ {
		content.table.AddCol()
	}

	f.cm.Lock()
	f.contents = append(f.contents, content)
	f.cm.Unlock()

	return content.table
}

// AddTableWithColumnsAndAutoWidth 添加一个拥有指定列数量的表格，并且自动计算所有列的宽度
func (f *Footer) AddTableWithColumnsAndAutoWidth(columns int) *table.Table {
	content := new(DocumentContent)
	content.ct = DocumentContentTypeTable

	content.table = new(table.Table)

	for i := 0; i < columns; i++ {
		content.table.AddColWithWidth(TableDefaultWidth / columns)
	}

	f.cm.Lock()
	f.contents = append(f.contents, content)
	f.cm.Unlock()

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

func (d *DocumentContent) GetContentType() DocumentContentType {
	return d.ct
}

func (d *DocumentContent) GetParagraph() *paragraph.Paragraph {
	return d.paragraph
}

func (d *DocumentContent) GetTable() *table.Table {
	return d.table
}

type DocumentContentType string

const (
	DocumentContentTypeParagraph = DocumentContentType("paragraph")
	DocumentContentTypeTable     = DocumentContentType("table")
)
