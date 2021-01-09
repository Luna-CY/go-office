package docx

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
    "strconv"
    "time"
)

// Core core.xml
type Core struct {
    // CreateUser 创建文档的用户名称
    CreateUser string

    // CreateTime 文档的创建时间
    CreateTime time.Time

    // LastModifyUser 最后修改的用户名
    LastModifyUser string

    // LastModifyTime 最后修改的时间
    LastModifyTime time.Time

    // Version 版本号，每修改一次+1
    Version int
}

// GetBody 获取core.xml文件的内容
func (c *Core) GetBody() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.CoreXmlStart)

    buffer.WriteString(template.CoreCreateUserStart)
    buffer.WriteString(c.CreateUser)
    buffer.WriteString(template.CoreCreateUserEnd)

    buffer.WriteString(template.CoreCreateTimeStart)
    buffer.WriteString(c.CreateTime.UTC().String())
    buffer.WriteString(template.CoreCreateTimeEnd)

    buffer.WriteString(template.CoreModifyUserStart)
    buffer.WriteString(c.LastModifyUser)
    buffer.WriteString(template.CoreModifyUserEnd)

    buffer.WriteString(template.CoreModifyTimeStart)
    buffer.WriteString(c.LastModifyTime.UTC().String())
    buffer.WriteString(template.CoreModifyTimeEnd)

    buffer.WriteString(template.CoreVersionStart)
    buffer.WriteString(strconv.Itoa(c.Version + 1))
    buffer.WriteString(template.CoreVersionEnd)

    buffer.WriteString(template.CoreXmlEnd)

    return buffer.Bytes(), nil
}
