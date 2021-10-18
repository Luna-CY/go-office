package xlsx

import (
	"fmt"
	"strconv"
	"sync"
)

// Workbook 工作簿xlsx
type Workbook struct {
	// metaFile 元数据结构定义
	metaFile WorkBookMetaFile

	sm sync.RWMutex
	// sheets 数据表组
	sheets []*Sheet
}

// NewSheet 新建一个数据表
func (w *Workbook) NewSheet() *Sheet {
	name := fmt.Sprintf("Sheet%d", len(w.sheets)+1)

	return w.NewSheetWithName(name)
}

// NewSheetWithName 新建一个工作表并设置表名称
func (w *Workbook) NewSheetWithName(name string) *Sheet {
	sheet := new(Sheet)

	sheet.meta.name = name

	w.sm.Lock()
	defer w.sm.Unlock()

	sheet.meta.id = strconv.Itoa(len(w.sheets) + 1)
	sheet.meta.relationshipId = fmt.Sprintf("rId%d", len(w.sheets)+1)
	w.sheets = append(w.sheets, sheet)

	return sheet
}
