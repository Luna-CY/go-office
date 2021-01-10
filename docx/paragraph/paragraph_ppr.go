package paragraph

import (
    "fmt"
    "github.com/Luna-CY/go-office/docx/section"
)

// PPr 段落的样式属性定义
type PPr struct {
    // horizontalAlignment 水平对齐方式
    horizontalAlignment *HorizontalAlignment

    // borderManager 边框管理器
    borderManager *BorderManager

    // keepLines 该段落是否尽可能的在一个页面上显示
    keepLines bool

    // keepNext 该段落与下一个段落是否尽可能的在一个页面上显示
    keepNext bool

    // identity 缩进配置结构
    identity *Identity

    // background 背景配置结构
    background *Background

    // spacing 段落间距配置
    spacing *Spacing

    // sect 段落内的章节属性配置
    sect *section.Section
}

// GetId 获取ID
func (p *PPr) GetId() string {
    return fmt.Sprintf("%p", p)
}

// GetBorder 获取边框管理器
func (p *PPr) GetBorder() *BorderManager {
    if nil == p.borderManager {
        p.borderManager = new(BorderManager)
        p.borderManager.isSet = false
    }

    return p.borderManager
}

// GetIdentity 获取缩进配置结构指针
func (p *PPr) GetIdentity() *Identity {
    if nil == p.identity {
        p.identity = new(Identity)
        p.identity.isSet = false
    }

    return p.identity
}

// GetBackground 获取背景配置结构指针
func (p *PPr) GetBackground() *Background {
    if nil == p.background {
        p.background = new(Background)
        p.background.isSet = false
    }

    return p.background
}

// GetSpacing 获取段落间距配置结构指针
func (p *PPr) GetSpacing() *Spacing {
    if nil == p.spacing {
        p.spacing = new(Spacing)
        p.spacing.isSet = false
    }

    return p.spacing
}

// GetSection 获取节属性配置
func (p *PPr) GetSection() *section.Section {
    if nil == p.sect {
        p.sect = new(section.Section)
    }

    return p.sect
}

// SetKeepLines
func (p *PPr) SetKeepLines(keepLines bool) *PPr {
    p.keepLines = keepLines

    return p
}

// SetKeepNext
func (p *PPr) SetKeepNext(keepNext bool) *PPr {
    p.keepNext = keepNext

    return p
}

// SetHorizontalAlignment 设置水平对齐方式
func (p *PPr) SetHorizontalAlignment(alignment HorizontalAlignment) *PPr {
    p.horizontalAlignment = &alignment

    return p
}

type HorizontalAlignment string

const (
    // HorizontalAlignmentStart 左对齐
    HorizontalAlignmentStart = HorizontalAlignment("start")

    // HorizontalAlignmentEnd 右对齐
    HorizontalAlignmentEnd = HorizontalAlignment("end")

    // HorizontalAlignmentCenter 居中对齐
    HorizontalAlignmentCenter = HorizontalAlignment("center")

    // HorizontalAlignmentBoth 左右对齐，不改变字符间距
    HorizontalAlignmentBoth = HorizontalAlignment("both")

    // HorizontalAlignmentDistribute 左右对齐，改变字符间距
    HorizontalAlignmentDistribute = HorizontalAlignment("distribute")
)
