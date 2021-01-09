package paragraph

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
)

type Identity struct {
    // isSet 是否设置缩进
    isSet bool

    // 左缩进
    left *int

    // 右缩进
    right *int

    // 删除的缩进，与 firstLine 是互斥的
    hanging *int

    // 首行的缩进，与 hanging 是互斥的，设置 hanging 属性是此属性无效
    firstLine *int
}

// SetLeft 设置左侧缩进
func (i *Identity) SetLeft(left int) *Identity {
    i.isSet = true
    i.left = &left

    return i
}

// SetRight 设置右侧缩进
func (i *Identity) SetRight(right int) *Identity {
    i.isSet = true
    i.right = &right

    return i
}

// SetHanging
func (i *Identity) SetHanging(hanging int) *Identity {
    i.isSet = true
    i.hanging = &hanging

    return i
}

// SetFirstLine 设置首行缩进
func (i *Identity) SetFirstLine(firstLine int) *Identity {
    i.isSet = true
    i.firstLine = &firstLine

    return i
}

func (i *Identity) GetBody() ([]byte, error) {
    if !i.isSet {
        return []byte{}, nil
    }

    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.ParagraphPPrIdentityTag)

    if nil != i.left {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityLeft, *i.left))
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityStart, *i.left))
    }

    if nil != i.right {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityRight, *i.right))
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityEnd, *i.right))
    }

    if nil != i.hanging {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityHanging, *i.hanging))
    }

    if nil != i.firstLine {
        buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.ParagraphPPrIdentityFirstLine, *i.firstLine))
    }
    buffer.WriteString("/>")

    return buffer.Bytes(), nil
}
