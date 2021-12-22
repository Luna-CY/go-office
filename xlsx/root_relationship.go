package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"sync"
)

func NewRootRelationship() *RootRelationship {
	rrs := new(RootRelationship)
	rrs.RootNamespace = "http://schemas.openxmlformats.org/package/2006/relationships"

	return rrs
}

type RootRelationship struct {
	XMLName xml.Name `xml:"Relationships"`

	RootNamespace string `xml:"xmlns,attr"`

	rm            sync.Mutex
	Relationships []Relationship

	nm      sync.Mutex
	nextRId int
}

func (r *RootRelationship) NextRId() string {
	r.nm.Lock()
	defer r.nm.Unlock()

	r.nextRId += 1

	return fmt.Sprintf("rId%v", r.nextRId)
}

func (r *RootRelationship) AddRelationship(rel Relationship) *RootRelationship {
	r.rm.Lock()
	defer r.rm.Unlock()

	r.Relationships = append(r.Relationships, rel)

	return r
}

func (r *RootRelationship) Filepath() string {
	return "_rels/.rels"
}

func (r *RootRelationship) Marshal() ([]byte, error) {
	r.rm.Lock()
	defer r.rm.Unlock()

	buffer := &bytes.Buffer{}
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(r); nil != err {
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
