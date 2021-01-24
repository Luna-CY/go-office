package docx

import (
    "github.com/Luna-CY/go-office/docx/footer"
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
    // headers 页头列表
    headers []*header.Header

    // section 文档的节属性配置
    section section.Section

    // contentTypes [Content_Types].xml
    contentTypes ContentTypes

    // relationship document.xml.rels
    relationship Relationships

    uhm sync.RWMutex
    // headers 页头配置
    useHeaders map[HeaderType]*header.Header

    fm sync.RWMutex
    // footers 页脚列表
    footers []*footer.Footer

    ufm sync.RWMutex
    // useFooters 页脚配置
    useFooters map[FooterType]*footer.Footer
}

func (d *Document) GetProperties() *Styles {
    return &d.styles
}

func (d *Document) GetContents() []*DocumentContent {
    d.cm.RLock()
    defer d.cm.RUnlock()

    return d.contents
}

func (d *Document) GetHeaders() []*header.Header {
    d.hm.RLock()
    defer d.hm.RUnlock()

    return d.headers
}

func (d *Document) GetFooters() []*footer.Footer {
    d.fm.RLock()
    defer d.fm.RUnlock()

    return d.footers
}

// GetSection 获取节属性配置结构指针
func (d *Document) GetSection() *section.Section {
    return &d.section
}

// GetBackground 获取背景配置结构指针
func (d *Document) GetBackground() *Background {
    return &d.background
}

// UseHeader 使用指定的页头
func (d *Document) UseHeader(headerType HeaderType, hdr *header.Header) *Document {
    d.uhm.Lock()
    defer d.uhm.Unlock()

    if nil == d.useHeaders {
        d.useHeaders = make(map[HeaderType]*header.Header)
    }

    d.useHeaders[headerType] = hdr

    return d
}

// UseFooter 使用指定的页脚
func (d *Document) UseFooter(footerType FooterType, ftr *footer.Footer) *Document {
    d.ufm.Lock()
    defer d.ufm.Unlock()

    if nil == d.useFooters {
        d.useFooters = make(map[FooterType]*footer.Footer)
    }

    d.useFooters[footerType] = ftr

    return d
}
