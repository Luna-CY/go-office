package docx

import (
    "github.com/Luna-CY/go-office/docx/template"
    "strconv"
    "strings"
)

// App app.xml
type App struct {
    // Application 生产者应用的名称
    Application string

    // Security 加密级别
    Security int
}

func (a *App) GetBody() ([]byte, error) {
    builder := new(strings.Builder)

    builder.WriteString(template.Xml)
    builder.WriteString(template.AppXmlStart)

    builder.WriteString(template.AppApplicationStart)
    builder.WriteString(a.Application)
    builder.WriteString(template.AppApplicationEnd)

    builder.WriteString(template.AppSecurityStart)
    builder.WriteString(strconv.Itoa(a.Security))
    builder.WriteString(template.AppSecurityEnd)

    builder.WriteString(template.AppXmlEnd)

    return []byte(builder.String()), nil
}
