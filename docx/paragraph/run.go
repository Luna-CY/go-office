package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
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
    buffer := new(bytes.Buffer)
    buffer.WriteString(template.RunStart)
    buffer.WriteString(r.Text)
    buffer.WriteString(template.RunEnd)

    return buffer.Bytes(), nil
}

// RPr 内容样式定义
type RPr struct{}
