package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Paragraph 段落结构
// 每个段落都包含自己的样式属性定义以及任意个 Run 结构
type Paragraph struct {

    // ppr 样式属性定义
    ppr *PPr

    // runs Run 结构列表
    runs []*Run
}

func (p *Paragraph) GetPPr() *PPr {
    if nil == p.ppr {
        p.ppr = new(PPr)
    }

    return p.ppr
}

func (p *Paragraph) GetBody() ([]byte, error) {
    runBuffer := new(bytes.Buffer)

    for _, run := range p.runs {
        body, err := run.GetBody()
        if nil != err {
            return nil, err
        }

        runBuffer.Write(body)
    }

    bodyBuffer := new(bytes.Buffer)
    bodyBuffer.WriteString(template.ParagraphStart)

    if nil != p.ppr {
        bodyBuffer.WriteString(template.ParagraphPPrStart)
        body, err := p.ppr.GetBody()
        if nil != err {
            return nil, err
        }

        bodyBuffer.Write(body)
        bodyBuffer.WriteString(template.ParagraphPPrEnd)
    }

    bodyBuffer.Write(runBuffer.Bytes())
    bodyBuffer.WriteString(template.ParagraphEnd)

    return bodyBuffer.Bytes(), nil
}

// AddRun 新增一个内容结构
func (p *Paragraph) AddRun() *Run {
    r := new(Run)
    p.runs = append(p.runs, r)

    return r
}

// AddBreakLine 添加换行符
func (p *Paragraph) AddBreakLine(breakLineType BreakLineType, breakLineClearType BreakLineClearType) {
    r := new(Run)
    p.runs = append(p.runs, r)

    r.Text = strings.Replace(template.BreakLine, "{{TYPE}}", string(breakLineType), 1)
    r.Text = strings.Replace(template.BreakLine, "{{TYPE}}", string(breakLineClearType), 1)
}

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
