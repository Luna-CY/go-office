package section

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// LineNumber 行号配置结构
type LineNumber struct {
    // isSet 是否设置行号
    isSet bool

    // CountBy 显示行号的步进增量
    CountBy uint

    // Distance 行号与文字的间距
    // 单位为点的二十分之一
    Distance uint

    // Start 起始行号
    Start uint

    // Restart 重新开始行号的方式
    Restart *LineNumberRestart
}

// SetCountBy 设置步进增量
func (l *LineNumber) SetCountBy(countBy uint) *LineNumber {
    l.isSet = true
    l.CountBy = countBy

    return l
}

// SetDistance 设置间距
func (l *LineNumber) SetDistance(distance uint) *LineNumber {
    l.isSet = true
    l.Distance = distance

    return l
}

// SetStart 设置起始行号
func (l *LineNumber) SetStart(start uint) *LineNumber {
    l.isSet = true
    l.Start = start

    return l
}

// SetRestart 设置重置方式
func (l *LineNumber) SetRestart(restart LineNumberRestart) *LineNumber {
    l.isSet = true
    l.Restart = &restart

    return l
}

func (l *LineNumber) GetXmlBytes() ([]byte, error) {
    if !l.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.SectionLineNumberTag)

    if 0 != l.Start {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionLineNumberStart, l.Start))
    }

    if 0 != l.CountBy {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionLineNumberCountBy, l.CountBy))
    }

    if 0 != l.Distance {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionLineNumberDistance, l.Distance))
    }

    if nil != l.Restart {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionLineNumberRestart, l.Restart))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

type LineNumberRestart string

const (
    // LineNumberRestartContinuous 连续的行号，不进行重置
    LineNumberRestartContinuous = LineNumberRestart("continuous")

    // LineNumberRestartNewPage 在新的一页重置
    LineNumberRestartNewPage = LineNumberRestart("newPage")

    // LineNumberRestartNewSection 在新的章节
    LineNumberRestartNewSection = LineNumberRestart("newSection")
)
