package xlsx

// Cell 单元格对象定义
type Cell struct {
	formula *Formula // 公式对象引用
}

func (s *Cell) SetFormula(formula *Formula) {
	s.formula = formula
}
