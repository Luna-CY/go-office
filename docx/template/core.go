package template

const (
	CoreXmlStart        = `<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`
	CoreCreateTimeStart = `<dcterms:created xsi:type="dcterms:W3CDTF">`
	CoreCreateTimeEnd   = `</dcterms:created>`
	CoreCreateUserStart = `<dc:creator>`
	CoreCreateUserEnd   = `</dc:creator>`
	CoreModifyTimeStart = `<dcterms:modified xsi:type="dcterms:W3CDTF">`
	CoreModifyTimeEnd   = `</dcterms:modified>`
	CoreModifyUserStart = `<cp:lastModifiedBy>`
	CoreModifyUserEnd   = `</cp:lastModifiedBy>`
	CoreVersionStart    = `<cp:revision>`
	CoreVersionEnd      = `</cp:revision>`
	CoreXmlEnd          = `</cp:coreProperties>`
)
