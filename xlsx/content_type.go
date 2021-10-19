package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"sync"
)

func NewContentType() *ContentType {
	ct := new(ContentType)
	ct.RootNamespace = "http://schemas.openxmlformats.org/package/2006/content-types"

	return ct
}

type ContentType struct {
	XMLName xml.Name `xml:"Types"`

	RootNamespace string `xml:"xmlns,attr"`

	dm       sync.Mutex
	Defaults []Default

	om        sync.Mutex
	Overrides []Override
}

func (c *ContentType) AddDefault(d Default) *ContentType {
	c.dm.Lock()
	defer c.dm.Unlock()

	c.Defaults = append(c.Defaults, d)

	return c
}

func (c *ContentType) AddOverride(o Override) *ContentType {
	c.om.Lock()
	defer c.om.Unlock()

	c.Overrides = append(c.Overrides, o)

	return c
}

func (c *ContentType) Filepath() string {
	return "[Content_Types].xml"
}

func (c *ContentType) Marshal() ([]byte, error) {
	c.dm.Lock()
	defer c.dm.Unlock()

	c.om.Lock()
	defer c.om.Unlock()

	buffer := &bytes.Buffer{}
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`)

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(c); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}

type Default struct {
	XMLName xml.Name `xml:"Default"`

	Extension string `xml:"Extension,attr"`

	ContentType string `xml:"ContentType,attr"`
}

// Override Override结构定义
type Override struct {
	XMLName xml.Name `xml:"Override"`

	// PartName 单元名称
	PartName string `xml:"PartName,attr"`

	// ContentType Type类型
	ContentType string `xml:"ContentType,attr"`
}
