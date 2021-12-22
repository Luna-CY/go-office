package sheet

import "encoding/xml"

type FormatProperty struct {
	XMLName xml.Name `xml:"sheetFormatPr"`

	BaseColWidth     int `xml:"baseColWidth,attr,omitempty"`
	DefaultColWidth  int `xml:"defaultColWidth,attr,omitempty"`
	DefaultRowHeight int `xml:"defaultRowHeight,attr,omitempty"`
}
