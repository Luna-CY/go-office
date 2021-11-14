package cell

import "encoding/xml"

// Cell 单元格结构定义
type Cell struct {
	Name xml.Name `xml:"c"`

	// Id 单元格ID
	Id string

	// content 单元格内容
	content interface{}
}

// SetContent 设置单元格内容
func (c *Cell) SetContent(content interface{}) *Cell {
	c.content = content

	return c
}
