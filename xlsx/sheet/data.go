package sheet

import (
	"encoding/xml"
	"github.com/Luna-CY/go-office/xlsx/sheet/row"
	"strconv"
	"sync"
)

type Data struct {
	XMLName xml.Name `xml:"sheetData"`

	rm   sync.RWMutex
	Rows []*row.Row
}

func (d *Data) GetRowNextId() string {
	d.rm.RLock()
	defer d.rm.RUnlock()

	return strconv.Itoa(len(d.Rows) + 1)
}

func (d *Data) NewRow() *row.Row {
	r := new(row.Row)
	d.AddRow(r)

	return r
}

func (d *Data) AddRow(r *row.Row) *Data {
	d.rm.Lock()
	defer d.rm.Unlock()

	d.Rows = append(d.Rows, r)

	return d
}
