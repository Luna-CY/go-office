package xlsx

import "encoding/xml"

// WorkBookMetaFile 工作簿元信息文件结构
type WorkBookMetaFile struct {
	// contentType [Content_Types].xml
	contentType *ContentType

	// rootRelationship /_rels/.rels
	rootRelationship *RootRelationship

	// docPropsApp /docProps/app.xml
	docPropsApp *DocPropsApp

	// docPropsCore /docProps/core.xml
	docPropsCore *DocPropsCore

	// WorkbookRelationship /xl/_rels/workbook.xml.rels
	WorkbookRelationship *WorkbookRelationship
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
