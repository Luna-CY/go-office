package workbook

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"sync"
)

func NewBookRelationship() *BookRelationship {
	wrs := new(BookRelationship)
	wrs.RootNamespace = "http://schemas.openxmlformats.org/package/2006/relationships"

	return wrs
}

type BookRelationship struct {
	XMLName xml.Name `xml:"Relationships"`

	RootNamespace string `xml:"xmlns,attr"`

	rm            sync.Mutex
	Relationships []Relationship

	nm      sync.Mutex
	nextRId int
}

func (w *BookRelationship) NextRId() string {
	w.nm.Lock()
	defer w.nm.Unlock()

	w.nextRId += 1

	return fmt.Sprintf("rId%v", w.nextRId)
}

func (w *BookRelationship) AddRelationship(rel Relationship) *BookRelationship {
	w.rm.Lock()
	defer w.rm.Unlock()

	w.Relationships = append(w.Relationships, rel)

	return w
}

func (w *BookRelationship) Filepath() string {
	return "xl/_rels/workbook.xml.rels"
}

func (w *BookRelationship) Marshal() ([]byte, error) {
	w.rm.Lock()
	defer w.rm.Unlock()

	buffer := &bytes.Buffer{}
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(w); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}

// Relationship 关联结构定义
type Relationship struct {
	XMLName xml.Name `xml:"Relationship"`

	// Id 关联的id
	Id string `xml:"Id,attr"`

	// Type 关联类型
	Type string `xml:"Type,attr"`

	// Target 目标位置
	Target string `xml:"Target,attr"`
}
