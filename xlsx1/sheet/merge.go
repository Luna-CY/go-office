package sheet

import "encoding/xml"

type MergeCells struct {
	Name xml.Name `xml:"mergeCells"`

	MergeCells []*MergeCell
}

type MergeCell struct {
	Name xml.Name `xml:"mergeCell"`

	Reference string `xml:"ref,attr"`
}
