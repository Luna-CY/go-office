package sheet

import "encoding/xml"

type Cols struct {
	Name xml.Name `xml:"cols"`

	Cols []*Col
}

type Col struct {
	Name xml.Name `xml:"col"`

	CustomWidth bool    `xml:"customWidth,attr"`
	Hidden      bool    `xml:"hidden,attr"`
	Min         uint    `xml:"min,attr"`
	Max         uint    `xml:"max,attr"`
	Style       uint    `xml:"s,attr"`
	Width       float32 `xml:"width,attr"`
}
