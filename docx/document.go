package docx

import (
    "github.com/Luna-CY/go-office/docx/paragraph"
    "github.com/Luna-CY/go-office/docx/section"
)

// Document 文档结构定义
// 该结构统一对.docx格式文档进行操作
// 生成的文档严格符合Office Open XML格式定义
type Document struct {
    // core core.xml
    core Core

    // app app.xml
    app App

    // style style.xml
    style StyleConfig

    // background 文档的背景配置
    background Background

    // paragraphs 文档的分段列表
    // 文档的全部内容都保存在这里
    paragraphs []*paragraph.Paragraph

    // section 文档的节属性配置
    section section.Section
}

// AddParagraph 添加一个段落
func (d *Document) AddParagraph() *paragraph.Paragraph {
    p := new(paragraph.Paragraph)
    d.paragraphs = append(d.paragraphs, p)

    return p
}

// GetSection 获取节属性配置结构指针
func (d *Document) GetSection() *section.Section {
    return &d.section
}

// GetBackground 获取背景配置结构指针
func (d *Document) GetBackground() *Background {
    return &d.background
}
