package docx

import (
    "github.com/Luna-CY/go-office/docx/template"
    "strconv"
    "strings"
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
    builder := new(strings.Builder)

    builder.WriteString(template.Xml)
    builder.WriteString(template.CoreXmlStart)

    builder.WriteString(template.CoreCreateUserStart)
    builder.WriteString(c.CreateUser)
    builder.WriteString(template.CoreCreateUserEnd)

    builder.WriteString(template.CoreCreateTimeStart)
    builder.WriteString(c.CreateTime.UTC().String())
    builder.WriteString(template.CoreCreateTimeEnd)

    builder.WriteString(template.CoreModifyUserStart)
    builder.WriteString(c.LastModifyUser)
    builder.WriteString(template.CoreModifyUserEnd)

    builder.WriteString(template.CoreModifyTimeStart)
    builder.WriteString(c.LastModifyTime.UTC().String())
    builder.WriteString(template.CoreModifyTimeEnd)

    builder.WriteString(template.CoreVersionStart)
    builder.WriteString(strconv.Itoa(c.Version + 1))
    builder.WriteString(template.CoreVersionEnd)

    builder.WriteString(template.CoreXmlEnd)

    return []byte(builder.String()), nil
}
