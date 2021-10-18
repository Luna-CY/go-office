package template

const Rel = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"><Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/><Relationship Id="rId2" Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties" Target="docProps/core.xml"/><Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties" Target="docProps/app.xml"/></Relationships>`

const (
	RelationshipXmlStart = `<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">`
	RelationshipXmlEnd   = `</Relationships>`

	RelationshipTag    = `Relationship`
	RelationshipId     = `Id`
	RelationshipType   = `Type`
	RelationshipTarget = `Target`
)
