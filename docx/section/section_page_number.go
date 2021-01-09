package section

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// PageNumber 页码配置结构
type PageNumber struct {
    // isSet 是否设置了页码
    isSet bool

    // chapSep 章节与页码的分隔符
    chapSep *PageNumberChapSep

    // chapStyle 指定从哪个标题查找章节编号
    chapStyle *int

    // fmt 页码的格式
    fmt PageNumberFmt

    // start 本节开始的页码
    // 如果不指定将从上一节的最高页码开始编号
    start *int
}

func (p *PageNumber) SetPageNumber(fmt PageNumberFmt) *PageNumber {
    p.isSet = true
    p.fmt = fmt

    return p
}

func (p *PageNumber) SetStart(start int) *PageNumber {
    p.isSet = true
    p.start = &start

    return p
}

func (p *PageNumber) SetChapStyle(chapStyle int) *PageNumber {
    p.isSet = true
    p.chapStyle = &chapStyle

    return p
}

func (p *PageNumber) SetChapSep(chapSep PageNumberChapSep) *PageNumber {
    p.isSet = true
    p.chapSep = &chapSep

    return p
}

func (p *PageNumber) GetXmlBytes() ([]byte, error) {
    if !p.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.SectionPageNumberTag)

    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageNumberFmt, p.fmt))

    if nil != p.start {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageNumberStart, *p.start))
    }

    if nil != p.chapStyle {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageNumberChapStyle, *p.chapStyle))
    }

    if nil != p.chapSep {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionPageNumberChapSep, *p.chapSep))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

type PageNumberChapSep string

const (
    PageNumberChapSepColon  = PageNumberChapSep("colon")
    PageNumberChapSepEmDash = PageNumberChapSep("emDash")
    PageNumberChapSepEndash = PageNumberChapSep("endash")
    PageNumberChapSepHyphen = PageNumberChapSep("hyphen")
    PageNumberChapSepPeriod = PageNumberChapSep("period")
)

type PageNumberFmt string

// 参考文档: http://officeopenxml.com/WPSectionPgNumType.php
const (
    // PageNumberFmtCardinalText the cardinal text of the run language. (In English, One, Two, Three, etc.)
    PageNumberFmtCardinalText = PageNumberFmt("cardinalText")

    // PageNumberFmtDecimal decimal numbering (1, 2, 3, 4, etc.)
    PageNumberFmtDecimal = PageNumberFmt("decimal")

    // PageNumberFmtDecimalEnclosedCircle decimal number enclosed in a circle
    PageNumberFmtDecimalEnclosedCircle = PageNumberFmt("decimalEnclosedCircle")

    // PageNumberFmtDecimalEnclosedFullstop decimal number followed by a period
    PageNumberFmtDecimalEnclosedFullstop = PageNumberFmt("decimalEnclosedFullstop")

    // PageNumberFmtDecimalEnclosedParen decimal number enclosed in parentheses
    PageNumberFmtDecimalEnclosedParen = PageNumberFmt("decimalEnclosedParen")

    // PageNumberFmtDecimalZero decimal number but with a zero added to numbers 1 through 9
    PageNumberFmtDecimalZero = PageNumberFmt("decimalZero")

    // PageNumberFmtLowerLetter based on the run language (e.g., a, b, c, etc.). Letters repeat for values greater than the size of the alphabet
    PageNumberFmtLowerLetter = PageNumberFmt("lowerLetter")

    // PageNumberFmtLowerRoman lowercase Roman numerals (i, ii, iii, iv, etc.)
    PageNumberFmtLowerRoman = PageNumberFmt("lowerRoman")

    // PageNumberFmtNone 不显示页码
    PageNumberFmtNone = PageNumberFmt("none")

    // PageNumberFmtOrdinalText ordinal text of the run language. (In English, First, Second, Third, etc.)
    PageNumberFmtOrdinalText = PageNumberFmt("ordinalText")

    // PageNumberFmtUpperLetter based on the run language (e.g., A, B, C, etc.). Letters repeat for values greater than the size of the alphabet
    PageNumberFmtUpperLetter = PageNumberFmt("upperLetter")

    // PageNumberFmtUpperRoman uppercase Roman numerals (I, II, III, IV, etc.)
    PageNumberFmtUpperRoman = PageNumberFmt("upperRoman")
)
