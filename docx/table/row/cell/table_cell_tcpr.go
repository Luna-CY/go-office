package cell

// TcPr 单元格属性配置
type TcPr struct {
	// noWrap 设置为文本不换行
	noWrap bool

	// gridSpan 指定该单元格跨越的逻辑列数
	gridSpan *int

	// tcFitText 调整单元格内的文本以适应单元格宽度
	tcFitText bool

	// vMerge 设置单元格垂直合并
	vMerge *VMergeType

	// background 单元格背景配置
	background *Background

	// border 单元格边框配置
	border *BorderManager

	// margin 单元格的边距管理
	margin *Margin

	// width 单元格宽度
	// 仅支持以点的二十分之一为单位
	width *int
}

// GetMargin 获取边距配置
func (t *TcPr) GetMargin() *Margin {
	if nil == t.margin {
		t.margin = new(Margin)
		t.margin.isSet = false
	}

	return t.margin
}

// GetBackground 获取背景配置结构
func (t *TcPr) GetBackground() *Background {
	if nil == t.background {
		t.background = new(Background)
		t.background.isSet = false
	}

	return t.background
}

// GetBorder 获取边框管理器
func (t *TcPr) GetBorder() *BorderManager {
	if nil == t.border {
		t.border = new(BorderManager)
		t.border.isSet = false
	}

	return t.border
}

// SetNoWrap 设置文本不换行
func (t *TcPr) SetNoWrap(noWrap bool) *TcPr {
	t.noWrap = noWrap

	return t
}

// SetVMerge 设置纵向合并类型
func (t *TcPr) SetVMerge(merge VMergeType) *TcPr {
	t.vMerge = &merge

	return t
}

// SetTcFitText 设置文本宽度自适应
func (t *TcPr) SetTcFitText(fit bool) *TcPr {
	t.tcFitText = fit

	return t
}

// SetGridSpan 设置表格合并的列数
func (t *TcPr) SetGridSpan(span int) *TcPr {
	t.gridSpan = &span

	return t
}

// GetWidth 获取单元格宽度
func (t *TcPr) GetWidth() int {
	if nil == t.width {
		return 0
	}

	return *t.width
}

// SetWidth 设置单元格宽度
func (t *TcPr) SetWidth(width int) *TcPr {
	t.width = &width

	return t
}

type VMergeType string

const (
	// VMergeTypeRestart 重新开始一组垂直合并
	VMergeTypeRestart = VMergeType("restart")

	// VMergeTypeContinue 继续之前的合并
	VMergeTypeContinue = VMergeType("continue")
)
