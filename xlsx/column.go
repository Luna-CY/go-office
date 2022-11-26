package xlsx

import "github.com/Luna-CY/go-office/common"

type Column struct {
	bestFit     *bool    // 是否自适应宽度，默认为true
	customWidth *bool    // 是否自定义宽度，该标志不允许手动设置
	hidden      *bool    // 设置是否隐藏此列
	min         *uint64  // 受此列样式影响的最小列序号
	max         *uint64  // 受此列样式影响的最大列序号
	style       *uint64  // 样式表索引序号
	width       *float64 // 列宽度
}

// SetAutoWidth 设置自动宽度
func (s *Column) SetAutoWidth(auto bool) {
	s.bestFit = &auto
	s.customWidth = common.NewPoint(false)
}

// SetHidden 设置是否隐藏
func (s *Column) SetHidden(hidden bool) {
	s.hidden = &hidden
}

// SetEnableColumns 设置此列样式应用的列范围
func (s *Column) SetEnableColumns(min uint64, max uint64) {
	s.min = &min
	s.max = &max
}

// SetWidth 设置列宽
func (s *Column) SetWidth(width float64) {
	s.width = &width
	s.customWidth = common.NewPoint(true)
}
