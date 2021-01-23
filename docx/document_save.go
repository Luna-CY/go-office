package docx

import (
    "archive/zip"
    "bytes"
    "errors"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "os"
)

// Save 保存文件到路径
// path 为一个完整的包含文件后缀名的路径地址
func (d *Document) Save(path string) error {
    // 在最后增加一个空的段落
    d.AddParagraph()

    for _, content := range d.GetContents() {
        if DocumentContentTypeParagraph == content.ct {
            p := content.paragraph

            d.styles.addParagraphStyle(p.GetProperties().GetId(), p.GetProperties())
            d.styles.addRunStyle(p.GetRunProperties().GetId(), p.GetRunProperties())

            for _, r := range p.GetRuns() {
                d.styles.addRunStyle(r.GetProperties().GetId(), r.GetProperties())
            }
        }

        if DocumentContentTypeTable == content.ct {
            t := content.table

            d.styles.addTableStyle(t.GetProperties().GetId(), t.GetProperties())
        }
    }

    d.contentTypes.AddContentType("/docProps/app.xml", ContentTypeTypeApp)
    d.contentTypes.AddContentType("/docProps/core.xml", ContentTypeTypeCore)
    d.contentTypes.AddContentType("/word/document.xml", ContentTypeTypeMain)
    d.contentTypes.AddContentType("/word/styles.xml", ContentTypeTypeStyle)

    d.relationship.AddRelationship("styles.xml", RelationshipTypeStyle)

    return d.save(path)
}

// save 保存文件
func (d *Document) save(path string) error {
    file, err := os.Create(path)
    if nil != err {
        return errors.New(fmt.Sprintf("创建文件失败: %v %v", path, err))
    }
    defer file.Close()

    word := zip.NewWriter(file)
    word.RegisterCompressor(zip.Store, nil)

    if err := d.saveContentTypesXml(word); nil != err {
        return err
    }

    if err := d.saveRelXml(word); nil != err {
        return err
    }

    if err := d.saveAppXml(word); nil != err {
        return err
    }

    if err := d.saveCoreXml(word); nil != err {
        return err
    }

    if err := d.saveDocumentXml(word); nil != err {
        return err
    }

    if err := d.saveWordRelXml(word); nil != err {
        return err
    }

    if err := d.saveStylesXml(word); nil != err {
        return err
    }

    if err := d.saveHeaders(word); nil != err {
        return err
    }

    if err := word.Close(); nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    return nil
}

// saveContentTypesXml 添加[Content_Types].xml文件
func (d *Document) saveContentTypesXml(word *zip.Writer) error {
    ct, err := word.Create("[Content_Types].xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    body, err := d.contentTypes.GetXmlBytes()
    if nil != err {
        return err
    }

    n, err := ct.Write(body)
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

// saveRelXml 保存全局关系表
func (d *Document) saveRelXml(word *zip.Writer) error {
    ct, err := word.Create("_rels/.rels")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write([]byte(template.Rel))
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(template.Rel) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(template.Rel), n))
    }

    return nil
}

// saveAppXml 保存app.xml
func (d *Document) saveAppXml(word *zip.Writer) error {
    ct, err := word.Create("docProps/app.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    body, err := d.app.GetXmlBytes()
    if nil != err {
        return err
    }

    n, err := ct.Write(body)
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

// saveCoreXml 保存core.xml
func (d *Document) saveCoreXml(word *zip.Writer) error {
    ct, err := word.Create("docProps/core.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    body, err := d.core.GetXmlBytes()
    if nil != err {
        return err
    }

    n, err := ct.Write(body)
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

// saveDocumentXml 保存主文档
func (d *Document) saveDocumentXml(word *zip.Writer) error {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.DocumentStart)

    body, err := d.background.GetXmlBytes()
    if nil != err {
        return err
    }
    buffer.Write(body)

    buffer.WriteString(template.DocumentBodyStart)

    for _, content := range d.GetContents() {
        if DocumentContentTypeParagraph == content.ct {
            body, err := content.paragraph.GetXmlBytes()
            if nil != err {
                return err
            }

            buffer.Write(body)
        }

        if DocumentContentTypeTable == content.ct {
            body, err := content.table.GetXmlBytes()
            if nil != err {
                return err
            }

            buffer.Write(body)
        }
    }

    buffer.WriteString(template.SectionStart)

    sectionBody, err := d.GetSection().GetXmlBytes()
    if nil != err {
        return err
    }
    buffer.Write(sectionBody)

    if nil != d.useHeaders {
        for typeType, hdr := range d.useHeaders {
            buffer.WriteString(fmt.Sprintf(`<%v %v="%v" %v="%v"/>`, template.SectionHeaderTag, template.SectionHeaderId, hdr.GetRId(), template.SectionHeaderType, typeType))
        }
    }

    buffer.WriteString(template.SectionEnd)

    buffer.WriteString(template.DocumentBodyEnd)
    buffer.WriteString(template.DocumentEnd)

    ct, err := word.Create("word/document.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write(buffer.Bytes())
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != buffer.Len() {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", buffer.Len(), n))
    }

    return nil
}

// saveWordRelXml 保存关联定义
func (d *Document) saveWordRelXml(word *zip.Writer) error {
    body, err := d.relationship.GetXmlBytes()
    if nil != err {
        return err
    }

    ct, err := word.Create("word/_rels/document.xml.rels")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write(body)
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

// saveStylesXml 保存样式表
func (d *Document) saveStylesXml(word *zip.Writer) error {
    body, err := d.styles.GetXmlBytes()
    if nil != err {
        return err
    }

    ct, err := word.Create("word/styles.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write(body)
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

// saveHeaders 保存页头
func (d *Document) saveHeaders(word *zip.Writer) error {
    for _, hdr := range d.headers {
        body, err := hdr.GetXmlBytes()
        if nil != err {
            return err
        }

        ct, err := word.Create(fmt.Sprintf("word/%v", hdr.GetFileName()))
        if nil != err {
            return errors.New(fmt.Sprintf("保存文档失败: %v", err))
        }

        n, err := ct.Write(body)
        if nil != err {
            return errors.New(fmt.Sprintf("保存文档失败: %v", err))
        }

        if n != len(body) {
            return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
        }
    }

    return nil
}
