package xlsx

// New 新建一个工作簿
func New() (*Workbook, error) {
	workbook := new(Workbook)
	workbook.metaFile.contentType = NewContentType()
	workbook.metaFile.rootRelationship = NewRootRelationship()
	workbook.metaFile.docPropsApp = NewDocPropsApp()
	workbook.metaFile.docPropsCore = NewDocPropsCore()
	workbook.metaFile.WorkbookRelationship = NewWorkbookRelationship()

	return workbook, nil
}
