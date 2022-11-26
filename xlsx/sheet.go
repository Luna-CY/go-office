package xlsx

// Sheet 工作表
type Sheet struct {
	name string // 工作表名称

	columns []*Column                 // 自定义列格式对象
	data    map[int64]map[int64]*Cell // 数据矩阵

	property struct {
		customHeight       *bool    // 是否自定义行高
		defaultColumnWidth *float64 // 默认列宽
		defaultRowHeight   *float64 // 默认行高
		zeroHeight         *bool    // 是否默认隐藏所有行
	}
}

// NewColumn 添加一个新的自定义列格式对象，该对象用于配置某些列的默认样式
func (s *Sheet) NewColumn() *Column {
	var column = new(Column)
	s.columns = append(s.columns, column)

	return column
}

// SetDefaultColumnWidth 设置默认列宽
func (s *Sheet) SetDefaultColumnWidth(width float64) {
	s.property.defaultColumnWidth = &width
}

// SetDefaultRowHeight 设置默认行高
func (s *Sheet) SetDefaultRowHeight(height float64) {
	s.property.defaultRowHeight = &height
}

func (s *Sheet) HiddenAllRows(hidden bool) {
	s.property.zeroHeight = &hidden
}
