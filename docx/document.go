package docx

import "github.com/Luna-CY/go-office/docx/section"

// Document 文档结构定义
// 该结构统一对.docx格式文档进行操作
// 生成的文档严格符合Office Open XML格式定义
type Document struct {

    // conformance 文档的模式设置
    // - strict 该文档严格遵循Office Open XML格式定义，此项默认值
    // - transitional 该文档遵循Office Open XML的过渡格式定义
    conformance string

    // core core.xml
    core Core

    // app app.xml
    app App

    // background 文档的背景配置
    background *Background

    // paragraphs 文档的分段列表
    // 文档的全部内容都保存在这里
    paragraphs []*Paragraph

    // section 文档的节属性配置
    section section.Section
}

// AddParagraph 添加一个段落
func (d *Document) AddParagraph() *Paragraph {
    p := new(Paragraph)
    d.paragraphs = append(d.paragraphs, p)

    return p
}

// GetSection 获取节属性配置结构指针
func (d *Document) GetSection() *section.Section {
    return &d.section
}
