package docx

import (
    "bytes"
    "fmt"
    "github.com/Luna-CY/go-office/docx/template"
    "sync"
)

// Relationships 文档关联定义
type Relationships struct {
    rm sync.RWMutex
    // relationships 关联定义
    relationships []Relationship
}

// AddRelationship 添加一个关联定义
func (r *Relationships) AddRelationship(target string, relationshipType RelationshipType) string {
    r.rm.Lock()
    defer r.rm.Unlock()

    relationship := Relationship{target: target, typeType: relationshipType}
    relationship.id = fmt.Sprintf("rId%d", len(r.relationships)+1)

    r.relationships = append(r.relationships, relationship)

    return relationship.id
}

func (r *Relationships) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.RelationshipXmlStart)

    r.rm.RLock()
    for _, rel := range r.relationships {
        buffer.WriteString(fmt.Sprintf(`<%v %v="%v" %v="%v" %v="%v"/>`, template.RelationshipTag, template.RelationshipId, rel.id, template.RelationshipType, rel.typeType, template.RelationshipTarget, rel.target))
    }
    r.rm.RUnlock()

    buffer.WriteString(template.RelationshipXmlEnd)

    return buffer.Bytes(), nil
}

// Relationship 关联定义
type Relationship struct {
    // id id
    id string

    // target 目标路径
    target string

    // typeType 类型
    typeType RelationshipType
}

type RelationshipType string

const (
    RelationshipTypeStyle     = RelationshipType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles")
    RelationshipTypeSetting   = RelationshipType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings")
    RelationshipTypeFontTable = RelationshipType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable")
    RelationshipTypeHeader    = RelationshipType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/header")
    RelationshipTypeFooter    = RelationshipType("http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer")
)
