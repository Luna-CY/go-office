package docx

import (
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

func (p *Paragraph) GetPPr() *PPr {
    if nil == p.ppr {
        p.ppr = new(PPr)
    }

    return p.ppr
}

func (p *Paragraph) GetBody() ([]byte, error) {
    runBuilder := new(strings.Builder)

    for _, run := range p.runs {
        body, err := run.GetBody()
        if nil != err {
            return nil, err
        }

        runBuilder.Write(body)
    }

    bodyBuilder := new(strings.Builder)
    bodyBuilder.WriteString(template.ParagraphStart)

    if nil != p.ppr {
        bodyBuilder.WriteString(template.ParagraphPPrStart)
        body, err := p.ppr.GetBody()
        if nil != err {
            return nil, err
        }

        bodyBuilder.WriteString(body)
        bodyBuilder.WriteString(template.ParagraphPPrEnd)
    }

    bodyBuilder.WriteString(runBuilder.String())
    bodyBuilder.WriteString(template.ParagraphEnd)

    return []byte(bodyBuilder.String()), nil
}

// PPr 段落的样式属性定义
type PPr struct {

    // ha 对齐方式
    ha *HorizontalAlignment

    topBorder     *Border
    bottomBorder  *Border
    leftBorder    *Border
    rightBorder   *Border
    betweenBorder *Border
}

func (p *PPr) GetBody() (string, error) {
    builder := new(strings.Builder)

    if nil != p.ha {
        builder.WriteString(strings.Replace(template.HorizontalAlignment, "{{TYPE}}", string(*p.ha), 1))
    }

    return builder.String(), nil
}

// SetHorizontalAlignment 设置水平对齐方式
func (p *PPr) SetHorizontalAlignment(alignment HorizontalAlignment) *PPr {
    p.ha = &alignment

    return p
}

// SetBorder 设置边框样式
func (p *PPr) SetBorder(top *Border, bottom *Border, left *Border, right *Border, between *Border) *PPr {
    p.topBorder = top
    p.bottomBorder = bottom
    p.leftBorder = left
    p.rightBorder = right
    p.betweenBorder = between

    return p
}

type HorizontalAlignment string

const (
    // HorizontalAlignmentStart 左对齐
    HorizontalAlignmentStart = HorizontalAlignment("start")

    // HorizontalAlignmentEnd 右对齐
    HorizontalAlignmentEnd = HorizontalAlignment("end")

    // HorizontalAlignmentCenter 居中对齐
    HorizontalAlignmentCenter = HorizontalAlignment("center")

    // HorizontalAlignmentBoth 左右对齐，不改变字符间距
    HorizontalAlignmentBoth = HorizontalAlignment("both")

    // HorizontalAlignmentDistribute 左右对齐，改变字符间距
    HorizontalAlignmentDistribute = HorizontalAlignment("distribute")
)

// Border 边框结构定义
type Border struct {
    // Color 颜色值，以十六进制字符串设置，不包含#号，例如FF0000
    Color string

    // Shadow 是否显示阴影
    Shadow bool

    // Space 间距偏移量，为1-72的值
    // 大于72的值将被设置为72
    Space uint8

    // Size 边框宽度
    // 最小值是2，最大值是96
    Size uint8

    // Style 边框样式
    // 段落边框只能是线条边框
    Style string
}
