package docx

import (
    "archive/zip"
    "errors"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "os"
    "strings"
)

// Save 保存文件到路径
// path 为一个完整的包含文件后缀名的路径地址
func (d *Document) Save(path string) error {
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

    body, err := d.app.GetBody()
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

    body, err := d.core.GetBody()
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
    bodyBuilder := new(strings.Builder)

    for _, paragraph := range d.paragraphs {
        body, err := paragraph.GetBody()
        if nil != err {
            return err
        }

        bodyBuilder.Write(body)
    }

    ct, err := word.Create("word/document.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    documentBuilder := new(strings.Builder)

    documentBuilder.WriteString(template.Xml)
    documentBuilder.WriteString(template.DocumentStart)

    documentBuilder.WriteString(bodyBuilder.String())
    sectionBody, err := d.GetSection().GetBody()
    if nil != err {
        return err
    }
    documentBuilder.Write(sectionBody)

    documentBuilder.WriteString(template.DocumentEnd)

    body := documentBuilder.String()
    n, err := ct.Write([]byte(body))
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(body) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(body), n))
    }

    return nil
}

func (d *Document) addStylesXml(word *zip.Writer) error {
    ct, err := word.Create("word/styles.xml")
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    n, err := ct.Write([]byte(template.Styles))
    if nil != err {
        return errors.New(fmt.Sprintf("保存文档失败: %v", err))
    }

    if n != len(template.Styles) {
        return errors.New(fmt.Sprintf("保存文档失败: 应写长度 %v 实写长度 %v", len(template.Styles), n))
    }

    return nil
}
