package section

import (
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// Section 节属性配置结构
type Section struct {
    // Cols 分栏设置
    Cols Cols
}

// GetCols 获取分栏配置结构指针
func (s *Section) GetCols() *Cols {
    return &s.Cols
}

func (s *Section) GetBody() ([]byte, error) {
    builder := new(strings.Builder)

    builder.WriteString(template.SectionStart)

    colsBody, err := s.Cols.GetBody()
    if nil != err {
        return nil, err
    }
    builder.Write(colsBody)

    builder.WriteString(template.SectionEnd)

    return []byte(builder.String()), nil
}
