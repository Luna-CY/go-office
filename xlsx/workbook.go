package xlsx

func NewWorkbook() *Workbook {
	return &Workbook{
		conformance: WorkbookConformanceStrict,
		sheets:      make([]*Sheet, 0),
	}
}

type Workbook struct {
	conformance WorkbookConformance // 工作簿的一致性等级，不可更改

	sheets []*Sheet // 所有工作表，按序输出
}

// NewSheet 新建数据表
func (s *Workbook) NewSheet() *Sheet {
	var sheet = new(Sheet)
	sheet.columns = make([]*Column, 0)
	sheet.data = make(map[int64]map[int64]*Cell)
	s.sheets = append(s.sheets, sheet)

	return sheet
}
