package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

func NewContentType() *ContentType {
	ct := new(ContentType)
	ct.RootNameSpace = "http://schemas.openxmlformats.org/package/2006/content-types"

	return ct
}

type ContentType struct {
	XMLName xml.Name `xml:"Types"`

	RootNameSpace string `xml:"xmlns,attr"`

	Defaults  []Default
	Overrides []Override
}

func (c *ContentType) AddDefault(d Default) *ContentType {
	c.Defaults = append(c.Defaults, d)

	return c
}

func (c *ContentType) AddOverride(o Override) *ContentType {
	c.Overrides = append(c.Overrides, o)

	return c
}

func (c *ContentType) FilePath() string {
	return "/[Content_Types].xml"
}

func (c *ContentType) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}

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
