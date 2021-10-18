package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

func NewRootRelationship() *RootRelationship {
	rrs := new(RootRelationship)
	rrs.RootNameSpace = "http://schemas.openxmlformats.org/package/2006/relationships"

	return rrs
}

type RootRelationship struct {
	XMLName xml.Name `xml:"Relationships"`

	RootNameSpace string `xml:"xmlns,attr"`

	Relationships []Relationship
}

func (r *RootRelationship) AddRelationship(rel Relationship) *RootRelationship {
	r.Relationships = append(r.Relationships, rel)

	return r
}

func (r *RootRelationship) FilePath() string {
	return "/_rels/.rels"
}

func (r *RootRelationship) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(r); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}
