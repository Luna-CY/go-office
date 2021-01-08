package section

import (
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Section 节属性配置结构
type Section struct {
    // cols 分栏设置
    cols *Cols

    // lineNumber 行号设置
    lineNumber *LineNumber

    // pageMargin 页边距配置
    pageMargin *PageMargin
}

// GetCols 获取分栏配置结构指针
func (s *Section) GetCols() *Cols {
    if nil == s.cols {
        s.cols = new(Cols)
        s.cols.isSet = false
        s.cols.SetNum(1)

        // 从wps文档中解析得到的默认值
        s.cols.SetSpace(425)
    }

    return s.cols
}

// GetLineNumber 获取行号配置结构指针
func (s *Section) GetLineNumber() *LineNumber {
    if nil == s.lineNumber {
        s.lineNumber = new(LineNumber)
        s.lineNumber.isSet = false
    }

    return s.lineNumber
}

// GetPageMargin 获取页边距配置结构指针
func (s *Section) GetPageMargin() *PageMargin {
    if nil == s.pageMargin {
        s.pageMargin = new(PageMargin)
        s.pageMargin.isSet = true

        // 从wps文档中解析得到的默认值
        s.pageMargin.SetMarginGroup(1440, 1800)
        s.pageMargin.SetHeader(851)
        s.pageMargin.SetFooter(992)
        s.pageMargin.SetGutter(0)
    }

    return s.pageMargin
}

func (s *Section) GetBody() ([]byte, error) {
    builder := new(strings.Builder)

    builder.WriteString(template.SectionStart)

    if nil != s.cols {
        colsBody, err := s.cols.GetBody()
        if nil != err {
            return nil, err
        }

        builder.Write(colsBody)
    }

    if nil != s.lineNumber {
        lineNumberBody, err := s.lineNumber.GetBody()
        if nil != err {
            return nil, err
        }

        builder.Write(lineNumberBody)
    }

    if nil != s.pageMargin {
        pageMarginBody, err := s.pageMargin.GetBody()
        if nil != err {
            return nil, err
        }

        builder.Write(pageMarginBody)
    }

    builder.WriteString(template.SectionEnd)

    return []byte(builder.String()), nil
}
