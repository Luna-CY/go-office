package run

import "fmt"

// RPr 内容样式定义
type RPr struct {
    // bold 粗体
    bold bool

    // italics 斜体
    italics bool

    // emboss 设置浮雕样式
    // 设置这个属性时如果未设置 color 属性
    // 将会自动将 color 设置为 FFF 白色
    // 该属性与 imprint shadow 为互斥属性
    emboss bool

    // imprint 刻印样式
    // 该属性与 emboss shadow 为互斥属性
    imprint bool

    // shadow 阴影样式
    // 该属性与 imprint emboss 为互斥属性
    shadow bool

    // vanish 不显示内容，仅显示样式
    vanish bool

    // color 文本颜色
    color *string

    // deleteLine 删除线样式
    deleteLine *DeleteLine

    // size 字体大小
    size *int

    // underline 下划线配置
    underline *Underline

    // background 背景设置
    background *Background

    // highlight 高亮背景色值
    // 设置高亮后将会无视背景设置
    highlight *Highlight
}

func (r *RPr) GetId() string {
    return fmt.Sprintf("%p", r)
}

func (r *RPr) GetBackground() *Background {
    if nil == r.background {
        r.background = new(Background)
        r.background.isSet = false
    }

    return r.background
}

// SetBold 设置粗体
func (r *RPr) SetBold(bold bool) *RPr {
    r.bold = bold

    return r
}

// SetItalics 设置斜体
func (r *RPr) SetItalics(italics bool) *RPr {
    r.italics = italics

    return r
}

// SetEmboss 设置浮雕样式
func (r *RPr) SetEmboss(emboss bool) *RPr {
    r.emboss = emboss
    if emboss {
        r.imprint = false
        r.shadow = false

        if nil == r.color {
            color := "FFF"
            r.color = &color
        }
    }

    return r
}

// SetImprint 设置刻印样式
func (r *RPr) SetImprint(imprint bool) *RPr {
    r.imprint = imprint
    if imprint {
        r.emboss = false
        r.shadow = false
    }

    return r
}

// SetShadow 设置阴影
func (r *RPr) SetShadow(shadow bool) *RPr {
    r.shadow = shadow
    if shadow {
        r.emboss = false
        r.imprint = false
    }

    return r
}

// SetVanish 仅显示样式，不显示内容
func (r *RPr) SetVanish(vanish bool) *RPr {
    r.vanish = vanish

    return r
}

// SetColor 设置文本颜色
func (r *RPr) SetColor(color string) *RPr {
    r.color = &color

    return r
}

// SetDeleteLine 设置删除线
func (r *RPr) SetDeleteLine(deleteLine DeleteLine) *RPr {
    r.deleteLine = &deleteLine

    return r
}

// SetSize 设置字体大小
func (r *RPr) SetSize(size int) *RPr {
    r.size = &size

    return r
}

// SetUnderline 设置下划线样式
func (r *RPr) SetUnderline(lineType UnderlineType) *RPr {
    underline := &Underline{lineType: lineType}
    r.underline = underline

    return r
}

// SetUnderlineWithColor 设置下划线样式与颜色
func (r *RPr) SetUnderlineWithColor(lineType UnderlineType, color string) *RPr {
    underline := &Underline{lineType: lineType, color: &color}
    r.underline = underline

    return r
}

// SetHighlight 设置高亮
func (r *RPr) SetHighlight(highlight Highlight) *RPr {
    r.highlight = &highlight

    return r
}
