package sheet

import (
	"encoding/xml"
	"sync"
)

type Views struct {
	XMLName xml.Name `xml:"sheetViews"`

	vm    sync.Mutex
	Views []View
}

func (s *Views) AddView(view View) *Views {
	s.vm.Lock()
	defer s.vm.Unlock()

	s.Views = append(s.Views, view)

	return s
}

type View struct {
	XMLName xml.Name `xml:"sheetView"`

	TabSelected    int `xml:"tabSelected,attr"`
	WorkbookViewId int `xml:"workbookViewId,attr"`
}
