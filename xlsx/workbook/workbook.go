package workbook

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/Luna-CY/go-office/xlsx/sheet"
	"strconv"
	"sync"
)

const (
	RelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	ContentType      = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
)

func NewWorkbook() *Workbook {
	wb := new(Workbook)
	wb.RootNamespace = "http://schemas.openxmlformats.org/spreadsheetml/2006/main"
	wb.RelationshipNamespace = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"

	wb.Views = new(BookViews)
	wb.Sheets = new(Sheets)

	wb.AddView(BookView{XWindow: 100, YWindow: 100, WindowWidth: 1000, WindowHeight: 1000})

	return wb
}

type Workbook struct {
	XMLName xml.Name `xml:"workbook"`

	RootNamespace         string `xml:"xmlns,attr"`
	RelationshipNamespace string `xml:"xmlns:r,attr"`

	Views  *BookViews
	Sheets *Sheets
}

func (w *Workbook) SheetNum() int {
	return w.Sheets.SheetNum()
}

func (w *Workbook) AddView(view BookView) *Workbook {
	w.Views.AddView(view)

	return w
}

func (w *Workbook) NewSheet(rId string) *sheet.Sheet {
	name := fmt.Sprintf("Sheet%d", w.Sheets.SheetNum()+1)

	return w.NewSheetWithName(name, rId)
}

func (w *Workbook) NewSheetWithName(name, rId string) *sheet.Sheet {
	sh := sheet.NewSheet(strconv.Itoa(w.Sheets.SheetNum()+1), name, rId)
	sh.SetDefaultRowHeight(24)
	sh.SetBaseColWidth(10)

	w.Sheets.AddSheet(SheetTag{Name: sh.Name(), SheetId: sh.Id(), RId: sh.RelationshipId()}, sh)

	return sh
}

func (w *Workbook) Filepath() string {
	return "xl/workbook.xml"
}

func (w *Workbook) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(w); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}

type BookViews struct {
	XMLName xml.Name `xml:"bookViews"`

	vm    sync.Mutex
	Views []BookView
}

func (b *BookViews) AddView(view BookView) *BookViews {
	b.vm.Lock()
	defer b.vm.Unlock()

	b.Views = append(b.Views, view)

	return b
}

type BookView struct {
	XMLName xml.Name `xml:"workbookView"`

	XWindow      int `xml:"xWindow,attr"`
	YWindow      int `xml:"yWindow,attr"`
	WindowWidth  int `xml:"windowWidth,attr"`
	WindowHeight int `xml:"windowHeight,attr"`
}

type Sheets struct {
	XMLName xml.Name `xml:"sheets"`

	sm     sync.Mutex
	Sheets []SheetTag

	sheets []*sheet.Sheet
}

func (s *Sheets) SheetNum() int {
	s.sm.Lock()
	defer s.sm.Unlock()

	return len(s.Sheets)
}

func (s *Sheets) AddSheet(sheetTag SheetTag, sheet *sheet.Sheet) *Sheets {
	s.sm.Lock()
	defer s.sm.Unlock()

	s.Sheets = append(s.Sheets, sheetTag)
	s.sheets = append(s.sheets, sheet)

	return s
}

func (s *Sheets) GetSheets() []*sheet.Sheet {
	s.sm.Lock()
	defer s.sm.Unlock()

	return s.sheets
}

type SheetTag struct {
	XMLName xml.Name `xml:"sheet"`

	Name    string `xml:"name,attr"`
	SheetId string `xml:"sheetId,attr"`
	RId     string `xml:"r:id,attr"`
}
