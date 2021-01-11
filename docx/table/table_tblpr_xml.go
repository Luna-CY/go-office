package table

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

func (t *TblPr) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.TblPrStart)

    if nil != t.horizontalAlignment {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrHorizontalAlignmentTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrVal, *t.horizontalAlignment))
        buffer.WriteString("/>")
    }

    if nil != t.layout {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrLayoutTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrType, *t.layout))
        buffer.WriteString("/>")
    }

    if nil != t.caption {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCaptionTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrVal, *t.caption))
        buffer.WriteString("/>")
    }

    if nil != t.width {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrWidthTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *t.width, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != t.cellSpacing {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrCellSpacingTag)
        // type为固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TblPrW, *t.cellSpacing, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != t.indentation {
        buffer.WriteByte('<')
        buffer.WriteString(template.TblPrIndentationTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrW, *t.indentation))
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrVal, *t.indentation))

        // 固定值dxa，点的二十分之一
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.TblPrType, "dxa"))
        buffer.WriteString("/>")
    }

    if nil != t.borderManager {
        body, err := t.borderManager.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != t.cellMargin {
        body, err := t.cellMargin.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.TblPrEnd)

    return buffer.Bytes(), nil
}
