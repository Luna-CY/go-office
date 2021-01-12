package row

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

func (t *TrPr) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)
    buffer.WriteString(template.TableRowTrPrStart)

    if t.cantSplit {
        buffer.WriteString(template.TableRowTrPrCantSplitTag)
    }

    if t.hidden {
        buffer.WriteString(template.TableRowTrPrHiddenTag)
    }

    if t.header {
        buffer.WriteString(template.TableRowTrPrHeaderTag)
    }

    if nil != t.cellSpacing {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableRowTrPrCellSpacingTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"/>`, template.TableRowTrPrType, "dxa", template.TableRowTrPrVal, *t.cellSpacing))
    }

    if nil != t.horizontalAlignment {
        buffer.WriteByte('<')
        buffer.WriteString(template.TableRowTrPrHorizontalAlignmentTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v"/>`, template.TableRowTrPrVal, *t.horizontalAlignment))
    }

    if nil != t.height {
        rule := HeightRuleTypeExact
        if nil != t.heightRule {
            rule = *t.heightRule
        }

        buffer.WriteByte('<')
        buffer.WriteString(template.TableRowTrPrHeightTag)
        buffer.WriteString(fmt.Sprintf(` %v="%v" %v="%v"`, template.TableRowTrPrVal, *t.height, template.TableRowTrPrHeightRule, rule))
    }

    buffer.WriteString(template.TableRowTrPrEnd)

    // 空标签就不用输出了
    empty := fmt.Sprintf(`%v%v`, template.TableRowTrPrStart, template.TableRowTrPrEnd)
    if empty == buffer.String() {
        return []byte{}, nil
    }

    return buffer.Bytes(), nil
}
