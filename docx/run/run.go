package run

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Run 内容的结构定义
type Run struct {
    rpr  *RPr
    body bytes.Buffer
}

// AddText 添加一段文本内容
func (r *Run) AddText(text string) *Run {
    r.body.WriteString(template.RunTextStart)
    r.body.WriteString(text)
    r.body.WriteString(template.RunTextEnd)

    return r
}

// AddBreakLine 添加换行符
func (r *Run) AddBreakLine(breakLineType BreakLineType, breakLineClearType BreakLineClearType) {
    content := strings.Replace(template.BreakLine, "{{TYPE}}", string(breakLineType), 1)
    content = strings.Replace(content, "{{CLEAR}}", string(breakLineClearType), 1)

    r.body.WriteString(content)
}

func (r *Run) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.RunStart)
    buffer.Write(r.body.Bytes())
    buffer.WriteString(template.RunEnd)

    return buffer.Bytes(), nil
}

// RPr 内容样式定义
type RPr struct{}

type BreakLineType string

const (
    // BreakLineTypeDefault 下一行开始
    // 此类型为默认类型
    BreakLineTypeDefault = BreakLineType("textWrapping")

    // BreakLineTypePage 从下一页开始
    // 设置为此类型时，新的内容将从下一页开始
    BreakLineTypePage = BreakLineType("page")

    // BreakLineTypePage 从下一列开始
    BreakLineTypeColumn = BreakLineType("column")
)

type BreakLineClearType string

const (
    // BreakLineClearTypeDefault 不设置clear
    // 这是默认值
    BreakLineClearTypeDefault = BreakLineClearType("none")

    // BreakLineClearTypeLeft 从左侧开始
    BreakLineClearTypeLeft = BreakLineClearType("left")

    // BreakLineClearTypeRight 从右侧开始
    BreakLineClearTypeRight = BreakLineClearType("right")

    // BreakLineClearTypeAll
    BreakLineClearTypeAll = BreakLineClearType("all")
)
