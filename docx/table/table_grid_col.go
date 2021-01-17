package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

// GridCol 表格列定义
type GridCol struct {
    // w 宽度
    w *int
}

// GetWidth 获取宽度
func (g *GridCol) GetWidth() int {
    if nil == g.w {
        return 0
    }

    return *g.w
}

// SetWidth 设置宽度
func (g *GridCol) SetWidth(w int) *GridCol {
    g.w = &w

    return g
}

func (g *GridCol) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableGridColTag)

    if nil != g.w {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TableGridColWidth, *g.w))
    }

    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}
