package run

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

func (r *RPr) GetStyleXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.RunRPrStyleTag)
    buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunRPrVal, r.GetId()))
    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}

func (r *RPr) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.RunRPrStart)

    if r.bold {
        buffer.WriteString(template.RunRPrBold)
    }

    if r.italics {
        buffer.WriteString(template.RunRPrItalics)
    }

    if r.emboss {
        buffer.WriteString(template.RunRPrEmboss)

        if nil == r.color {
            buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.RunRPrColorTag, template.RunRPrVal, "FFF"))
        }
    }

    if r.imprint {
        buffer.WriteString(template.RunRPrImprint)
    }

    if r.shadow {
        buffer.WriteString(template.RunRPrShadow)
    }

    if r.vanish {
        buffer.WriteString(template.RunRPrVanish)
    }

    if nil != r.color {
        buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.RunRPrColorTag, template.RunRPrVal, *r.color))
    }

    if nil != r.underline {
        buffer.WriteByte('<')
        buffer.WriteString(template.RunRPrUnderlineTag)

        if nil != r.underline.color {
            buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunRPrUnderlineColor, *r.underline.color))
        }

        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunRPrVal, r.underline.lineType))
        buffer.WriteString("/>")
    }

    if nil != r.deleteLine {
        if DeleteLineStrike == *r.deleteLine {
            buffer.WriteString(template.RunRPrStrike)
        }

        if DeleteLineDoubleStrike == *r.deleteLine {
            buffer.WriteString(template.RunRPrDStrike)
        }
    }

    if nil != r.size {
        buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.RunRPrSizeTag, template.RunRPrVal, *r.size))
    }

    if nil != r.background {
        body, err := r.background.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.WriteString(template.RunRPrEnd)

    return buffer.Bytes(), nil
}
