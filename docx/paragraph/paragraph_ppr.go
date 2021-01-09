package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// PPr 段落的样式属性定义
type PPr struct {

    // horizontalAlignment 水平对齐方式
    horizontalAlignment *HorizontalAlignment

    // borderManager 边框管理器
    borderManager BorderManager
}

// GetBorderManager 获取边框管理器
func (p *PPr) GetBorderManager() *BorderManager {
    return &p.borderManager
}

func (p *PPr) GetBody() ([]byte, error) {
    buffer := new(bytes.Buffer)

    if nil != p.horizontalAlignment {
        buffer.WriteString(strings.Replace(template.ParagraphPPrHorizontalAlignment, "{{TYPE}}", string(*p.horizontalAlignment), 1))
    }

    borderBody, err := p.borderManager.GetBody()
    if nil != err {
        return nil, nil
    }
    buffer.Write(borderBody)

    return buffer.Bytes(), nil
}

// SetHorizontalAlignment 设置水平对齐方式
func (p *PPr) SetHorizontalAlignment(alignment HorizontalAlignment) *PPr {
    p.horizontalAlignment = &alignment

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
