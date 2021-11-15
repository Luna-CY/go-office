package sheet

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

const (
	RelationShipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	ContentType      = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
)

func NewSheet(id, name, rId string) *Sheet {
	s := new(Sheet)

	s.id = id
	s.relationshipId = rId
	s.name = name
	s.RootNamespace = "http://schemas.openxmlformats.org/spreadsheetml/2006/main"

	s.Views = new(Views)
	s.Views.AddView(View{TabSelected: 1, WorkbookViewId: 0})

	s.Format = new(FormatProperty)
	s.Data = new(Data)

	return s
}

// Sheet 工作表定义
type Sheet struct {
	XMLName xml.Name `xml:"worksheet"`

	RootNamespace string `xml:"xmlns,attr"`

	id             string
	name           string
	relationshipId string

	Views  *Views
	Format *FormatProperty
	Data   *Data
	Cols   *Cols
}

func (s *Sheet) Id() string {
	return s.id
}

func (s *Sheet) SetId(id string) {
	s.id = id
}

func (s *Sheet) RelationshipId() string {
	return s.relationshipId
}

func (s *Sheet) Name() string {
	return s.name
}

// SetName 设置数据表名
func (s *Sheet) SetName(name string) *Sheet {
	s.name = name

	return s
}

func (s *Sheet) Filepath() string {
	return fmt.Sprintf("xl/worksheets/sheet%v.xml", s.id)
}

func (s *Sheet) SetBaseColWidth(num int) *Sheet {
	s.Format.BaseColWidth = num

	return s
}

func (s *Sheet) SetDefaultColWidth(width int) *Sheet {
	s.Format.DefaultColWidth = width

	return s
}

func (s *Sheet) SetDefaultRowHeight(height int) *Sheet {
	s.Format.DefaultRowHeight = height

	return s
}

func (s *Sheet) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(s); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}
