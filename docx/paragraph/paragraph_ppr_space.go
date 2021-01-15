package paragraph

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// Spacing 段落间距配置
type Spacing struct {
    // isSet 是否设置了段间距
    isSet bool

    // before 设置段落前的间距
    before *int

    // after 设置段落后的间距
    after *int

    // line 设置段落内行的间距
    line *int

    // lineRule 设置段落内行间距的规则
    lineRule *LineRuleType

    // beforeAutoSpacing 段落前间距是否由编辑器自行定义
    // 设置了该属性后将会忽略 before 属性
    beforeAutoSpacing bool

    // afterAutoSpacing 段落后间距是否由编辑器自行定义
    // 设置了该属性后将会忽略 after 属性
    afterAutoSpacing bool
}

// SetSpace 同时设置 before 与 after
func (s *Spacing) SetSpace(space int) *Spacing {
    s.isSet = true

    s.before = &space
    s.after = &space

    return s
}

// SetBefore 设置段落前间距
func (s *Spacing) SetBefore(before int) *Spacing {
    s.isSet = true
    s.before = &before

    return s
}

// SetAfter 设置段落后间距
func (s *Spacing) SetAfter(after int) *Spacing {
    s.isSet = true
    s.after = &after

    return s
}

// SetLine 设置段落内行间距
func (s *Spacing) SetLine(line int) *Spacing {
    s.isSet = true
    s.line = &line

    return s
}

// SetLineRule 设置段落内行间距规则
func (s *Spacing) SetLineRule(rule LineRuleType) *Spacing {
    s.isSet = true
    s.lineRule = &rule

    return s
}

// SetBeforeAutoSpacing 设置段落前间距自动
func (s *Spacing) SetBeforeAutoSpacing(auto bool) *Spacing {
    s.isSet = true
    s.beforeAutoSpacing = auto

    return s
}

// SetAfterAutoSpacing 设置段落后间距自动
func (s *Spacing) SetAfterAutoSpacing(auto bool) *Spacing {
    s.isSet = true
    s.afterAutoSpacing = auto

    return s
}

func (s *Spacing) GetXmlBytes() ([]byte, error) {
    if !s.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.ParagraphPPrSpacingTag)

    if nil != s.after {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingAfter, *s.after))
    }

    if nil != s.before {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingBefore, *s.before))
    }

    if nil != s.line {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingLine, *s.line))
    }

    if nil != s.lineRule {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingLineRule, *s.lineRule))
    }

    if s.beforeAutoSpacing {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingBeforeAutoSpacing, s.beforeAutoSpacing))
    }

    if s.afterAutoSpacing {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrSpacingAfterAutoSpacing, s.afterAutoSpacing))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

type LineRuleType string

const (
    LineRuleTypeAtLeast = LineRuleType("atLeast")
    LineRuleTypeExactly = LineRuleType("exactly")
    LineRuleTypeAuto    = LineRuleType("auto")
)
