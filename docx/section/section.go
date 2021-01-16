package section

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// Section 节属性配置结构
// TODO: 未支持Header/Footer/VAlign/Type/TitlePg/PaperSrc/FormPort等配置
type Section struct {
    // cols 分栏设置
    cols *Cols

    // lineNumber 行号设置
    lineNumber *LineNumber

    // pageMargin 页边距配置
    pageMargin *PageMargin

    // PageBorder 页边框配置
    pageBorder *PageBorder

    // pageNumber 页码配置
    pageNumber *PageNumber

    // 页面尺寸配置
    pageSize *PageSize
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

// GetPageBorder 获取页边框配置结构指针
func (s *Section) GetPageBorder() *PageBorder {
    if nil == s.pageBorder {
        s.pageBorder = new(PageBorder)
        s.pageBorder.isSet = false
    }

    return s.pageBorder
}

// GetPageNumber 获取页码配置结构指针
func (s *Section) GetPageNumber() *PageNumber {
    if nil == s.pageNumber {
        s.pageNumber = new(PageNumber)
        s.pageNumber.isSet = false
    }

    return s.pageNumber
}

// GetPageSize 获取页面尺寸配置结构指针
func (s *Section) GetPageSize() *PageSize {
    if nil == s.pageSize {
        s.pageSize = new(PageSize)
        s.pageSize.isSet = false
    }

    return s.pageSize
}

func (s *Section) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.SectionStart)

    if nil != s.cols {
        body, err := s.cols.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != s.lineNumber {
        body, err := s.lineNumber.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != s.pageMargin {
        body, err := s.pageMargin.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != s.pageBorder {
        body, err := s.pageBorder.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != s.pageNumber {
        body, err := s.pageNumber.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != s.pageSize {
        body, err := s.pageSize.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(`<w:docGrid w:type="lines" w:linePitch="312"/>`)

    buffer.WriteString(template.SectionEnd)

    empty := fmt.Sprintf(`%v%v`, template.SectionStart, template.SectionEnd)
    if empty == buffer.String() {
        return []byte{}, nil
    }

    return buffer.Bytes(), nil
}
