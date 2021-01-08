package docx

import (
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Run 内容的结构定义
type Run struct {
    rpr  *RPr
    Text string
}

func (r *Run) SetText(text string) *Run {
    r.Text = text

    return r
}

func (r *Run) GetBody() ([]byte, error) {
    bodyBuilder := new(strings.Builder)
    bodyBuilder.WriteString(template.RunStart)
    bodyBuilder.WriteString(r.Text)
    bodyBuilder.WriteString(template.RunEnd)

    return []byte(bodyBuilder.String()), nil
}

// RPr 内容样式定义
type RPr struct{}
