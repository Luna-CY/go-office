package template

const (
    TypesStart          = `<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">`
    TypesEnd            = `</Types>`
    DefaultXmlTag       = `<Default Extension="xml" ContentType="application/xml"/>`
    DefaultRelTag       = `<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>`
    OverrideTag         = `Override`
    OverridePartName    = `PartName`
    OverrideContentType = `ContentType`

    ContentTypeExtendProperties = `application/vnd.openxmlformats-officedocument.extended-properties+xml`
    ContentTypeCoreProperties   = `application/vnd.openxmlformats-package.core-properties+xml`
    ContentTypeWorkbook         = `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml`
    ContentTypeWorksheet        = `application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml`
)
