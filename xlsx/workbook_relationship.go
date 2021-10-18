package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

func NewWorkbookRelationship() *WorkbookRelationship {
	wrs := new(WorkbookRelationship)
	wrs.RootNameSpace = "http://schemas.openxmlformats.org/package/2006/relationships"

	return wrs
}

type WorkbookRelationship struct {
	XMLName xml.Name `xml:"Relationships"`

	RootNameSpace string `xml:"xmlns,attr"`

	Relationships []Relationship
}

func (w *WorkbookRelationship) AddRelationship(rel Relationship) *WorkbookRelationship {
	w.Relationships = append(w.Relationships, rel)

	return w
}

func (w *WorkbookRelationship) FilePath() string {
	return "/xl/_rels/workbook.xml.rels"
}

func (w *WorkbookRelationship) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(w); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}
