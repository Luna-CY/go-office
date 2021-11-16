package sheet

import (
	"encoding/xml"
	"fmt"
	"github.com/Luna-CY/go-office/xlsx/sheet/cell"
	"sync"
)

type Data struct {
	XMLName xml.Name `xml:"sheetData"`

	dm   sync.RWMutex
	data map[uint64]map[uint64]*cell.Cell
}

func (d *Data) SetCell(row, col uint64, value interface{}) *Data {
	d.dm.Lock()
	defer d.dm.Unlock()

	r, ok := d.data[row]
	if !ok {
		d.data[row] = map[uint64]*cell.Cell{}
		r = d.data[row]
	}

	r[col] = cell.NewCellWithTextForInline(row, col, fmt.Sprintf("%v", value))

	return d
}

func (d *Data) Marshal() ([]byte, error) {
	return nil, nil
}
