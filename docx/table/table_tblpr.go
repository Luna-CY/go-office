package table

import "fmt"

// TblPr 表格属性配置
type TblPr struct {
    // horizontalAlignment
    // 表格本身相对于页面的对齐方式，不是表格内内容的对齐方式
    horizontalAlignment *HorizontalAlignmentType

    // borderManager
    // 表格的边框管理器
    borderManager *BorderManager

    // caption
    // 设置表的标题，注意此属性是Word 2008之后添加的
    // 在Word 2007及以前是不支持的
    caption *string

    // indentation
    // 设置表格相对于文档的左侧缩进距离
    // 仅支持以点的二十分之一为单位
    indentation *int

    // width
    // 设置表格的宽度，该宽度仅作为默认值，在与表格实际宽度冲突时以实际为准
    // 仅支持以点的二十分之一为单位
    width *int

    // cellMargin
    // 单元格边距管理器
    cellMargin *CellMargin

    // cellSpacing
    // 单元格间距，该项指定的是每个单元格与相邻单元格或表格边框的距离
    // 仅支持以点的二十分之一为单位
    cellSpacing *int

    // layout
    // 表格的布局模式
    layout *LayoutType
}

// GetId 获取ID
func (t *TblPr) GetId() string {
    return fmt.Sprintf("%p", t)
}

// GetBorder 获取边框管理器
func (t *TblPr) GetBorder() *BorderManager {
    if nil == t.borderManager {
        t.borderManager = new(BorderManager)
        t.borderManager.isSet = false
    }

    return t.borderManager
}

// GetCellMargin 获取单元格边距管理器
func (t *TblPr) GetCellMargin() *CellMargin {
    if nil == t.cellMargin {
        t.cellMargin = new(CellMargin)
        t.cellMargin.isSet = false
    }

    return t.cellMargin
}

// SetHorizontalAlignment 设置水平对齐
func (t *TblPr) SetHorizontalAlignment(horizontalAlignment HorizontalAlignmentType) *TblPr {
    t.horizontalAlignment = &horizontalAlignment

    return t
}

// SetCaption 设置标题
func (t *TblPr) SetCaption(caption string) *TblPr {
    t.caption = &caption

    return t
}

// SetIndentation 设置缩进
func (t *TblPr) SetIndentation(indentation int) *TblPr {
    t.indentation = &indentation

    return t
}

// GetWidth 获取表格宽度
func (t *TblPr) GetWidth() int {
    if nil == t.width {
        return 0
    }

    return *t.width
}

// SetWidth 设置宽度
func (t *TblPr) SetWidth(width int) *TblPr {
    t.width = &width

    return t
}

// SetWidthAuto 设为宽度自动
func (t *TblPr) SetWidthAuto() *TblPr {
    t.width = nil

    return t
}

// SetCellSpacing 设置单元格间距
func (t *TblPr) SetCellSpacing(spacing int) *TblPr {
    t.cellSpacing = &spacing

    return t
}

type HorizontalAlignmentType string

const (
    // HorizontalAlignmentStart
    // 表格以文本段的左侧沿线对齐
    HorizontalAlignmentStart = HorizontalAlignmentType("left")

    // HorizontalAlignmentEnd
    // 表格以文本段的右侧沿线对齐
    HorizontalAlignmentEnd = HorizontalAlignmentType("right")

    // HorizontalAlignmentCenter
    // 表格以文本段的居中位置显示
    HorizontalAlignmentCenter = HorizontalAlignmentType("center")
)

type LayoutType string

const (
    // LayoutTypeFixed
    // 表示表格的单元格布局固定，不会自动调整单元格宽高
    LayoutTypeFixed = LayoutType("fixed")

    // LayoutTypeAuto
    // 表示单元格宽高自动处理
    LayoutTypeAuto = LayoutType("autofit")
)
