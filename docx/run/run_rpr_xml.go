package run

import (
	"bytes"
	"fmt"
	"github.com/Luna-CY/go-office/docx/template"
)

// GetInnerXmlBytes 获取内联样式
func (r *RPr) GetInnerXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.RunRPrStart)

	if nil != r.background {
		buffer.WriteByte('<')
		buffer.WriteString(template.RunRPrStyleTag)
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.RunRPrVal, r.GetId()))
		buffer.WriteString("/>")
	}

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

	buffer.WriteString(template.RunRPrEnd)

	empty := fmt.Sprintf(`%v%v`, template.RunRPrStart, template.RunRPrEnd)
	if buffer.String() == empty {
		return []byte{}, nil
	}

	return buffer.Bytes(), nil
}

// GetExtraXmlBytes 获取外部样式
func (r *RPr) GetExtraXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.RunRPrStart)

	if nil != r.background {
		body, err := r.background.GetXmlBytes()
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	buffer.WriteString(template.RunRPrEnd)

	empty := fmt.Sprintf(`%v%v`, template.RunRPrStart, template.RunRPrEnd)
	if buffer.String() == empty {
		return []byte{}, nil
	}

	return buffer.Bytes(), nil
}

func (r *RPr) GetDefaultXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.RunRPrStart)
	buffer.WriteString(`<w:lang w:val="en-US" w:eastAsia="zh-CN" w:bidi="ar-SA"/>`)

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
