package cell

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

func (t *TcPr) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)
    buffer.WriteString(template.TableCellTcPrStart)

    if t.noWrap {
        buffer.WriteString(template.TableCellTcPrNoWrapTag)
    }

    if t.tcFitText {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrTcFitTextTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"/>`, template.TableCellTcPrVal, t.tcFitText))
    }

    if nil != t.gridSpan {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrGridSpanTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"/>`, template.TableCellTcPrVal, *t.gridSpan))
    }

    if nil != t.vMerge {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrVMergeTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"/>`, template.TableCellTcPrVal, *t.vMerge))
    }

    if nil != t.background {
        body, err := t.background.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != t.border {
        body, err := t.border.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != t.margin {
        body, err := t.margin.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    if nil != t.width {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableCellTcPrCellWidth)
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"/>`, template.TableCellTcPrVal, *t.width, template.TableCellTcPrType, "dxa"))
    }

    buffer.WriteString(template.TableCellTcPrEnd)

    empty := fmt.Sprintf(`%v%v`, template.TableCellTcPrStart, template.TableCellTcPrEnd)
    if empty == buffer.String() {
        return []byte{}, nil
    }

    return buffer.Bytes(), nil
}
