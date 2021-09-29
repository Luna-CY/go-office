package xlsx

// SheetMeta 数据表元数据结构
type SheetMeta struct {
	// contentTypes /[Content_Types].xml
	contentTypes []ContentType

	// relationships /_rels/.rels
	relationships []Relationship

	// wbRelationships /xl/_rels/workbook.xml.rels
	wbRelationships []Relationship
}

// ContentType ContentType结构定义
type ContentType struct {
	// PartName 单元名称
	PartName string

	// ContentType Type类型
	ContentType string
}

// Relationship 关联结构定义
type Relationship struct {
	// Id 关联的id
	Id string

	// Type 关联类型
	Type string

	// Target 目标位置
	Target string
}
