package xlsx

// New 新建一个工作簿
func New() (*Workbook, error) {
	workbook := new(Workbook)

	return workbook, nil
}
