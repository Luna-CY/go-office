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

            d.style.addParagraphStyle(p.GetProperties().GetId(), p.GetProperties())
            d.style.addRunStyle(p.GetRunProperties().GetId(), p.GetRunProperties())

            for _, r := range p.GetRuns() {
                d.style.addRunStyle(r.GetProperties().GetId(), r.GetProperties())
            }
        }

        if DocumentContentTypeTable == content.ct {
            t := content.table

            d.style.addTableStyle(t.GetProperties().GetId(), t.GetProperties())
        }
    }

    file, err := os.Create(path)
    if nil != err {
        return errors.New(fmt.Sprintf("创建文件失败: %v %v", path, err))
    }
    defer file.Close()

    word := zip.NewWriter(file)
    word.RegisterCompressor(zip.Store, nil)

    if err := d.addContentTypesXml(word); nil != err {
        return err
    }

    if err := d.addRelXml(word); nil != err {
        return err
    }

    if err := d.addAppXml(word); nil != err {
        return err
    }

    if err := d.addCoreXml(word); nil != err {
        return err
    }

    if err := d.addDocumentXml(word); nil != err {
        return err
    }

    if err := d.addWordRelXml(word); nil != err {
        return err
    }

    if err := d.addStylesXml(word); nil != err {
        return err
    }

    if err := word.Close(); nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    return nil
}

// addContentTypesXml 添加[Content_Types].xml文件
func (d *Document) addContentTypesXml(word *zip.Writer) error {
    ct, err := word.Create("[Content_Types].xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write([]byte(template.ContentTypes))
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(template.ContentTypes) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(template.ContentTypes), n))
    }

    return nil
}

func (d *Document) addRelXml(word *zip.Writer) error {
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

func (d *Document) addAppXml(word *zip.Writer) error {
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

func (d *Document) addCoreXml(word *zip.Writer) error {
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

func (d *Document) addDocumentXml(word *zip.Writer) error {
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

    sectionBody, err := d.GetSection().GetXmlBytes()
    if nil != err {
        return err
    }
    buffer.Write(sectionBody)

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

func (d *Document) addWordRelXml(word *zip.Writer) error {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.RelationshipXmlStart)
    buffer.WriteString(template.RelationshipStyle)
    buffer.WriteString(template.RelationshipXmlEnd)

    ct, err := word.Create("word/_rels/document.xml.rels")
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

func (d *Document) addStylesXml(word *zip.Writer) error {
    body, err := d.style.GetXmlBytes()
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
