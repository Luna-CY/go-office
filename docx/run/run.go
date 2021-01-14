package run

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Run 内容的结构定义
type Run struct {
    rpr *RPr

    body bytes.Buffer
}

func (r *Run) GetProperties() *RPr {
    if nil == r.rpr {
        r.rpr = new(RPr)
    }

    return r.rpr
}

// AddText 添加一段文本内容
func (r *Run) AddText(text string) *Run {
    r.addText(text, false)

    return r
}

// AddTextSpace 添加一段文本内容并保留空白
func (r *Run) AddTextSpace(text string) *Run {
    r.addText(text, true)

    return r
}

// AddBreakLine 添加换行符
func (r *Run) AddBreakLine(breakLineType BreakLineType, breakLineClearType BreakLineClearType) {
    content := strings.Replace(template.BreakLine, "{{TYPE}}", string(breakLineType), 1)
    content = strings.Replace(content, "{{CLEAR}}", string(breakLineClearType), 1)

    r.body.WriteString(content)
}

// addText 添加文本内容
func (r *Run) addText(text string, space bool) {
    r.body.WriteByte('<')
    r.body.WriteString(template.RunTextTag)
    if space {
        r.body.WriteString(fmt.Sprintf(` %v="%v"`, template.RunTextSpace, "preserve"))
    }
    r.body.WriteByte('>')

    r.body.WriteString(text)

    r.body.Write([]byte{'<', '/'})
    r.body.WriteString(template.RunTextTag)
    r.body.WriteByte('>')
}