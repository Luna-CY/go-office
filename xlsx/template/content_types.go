package template

const (
    TypesStart = `<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">`
    TypesEnd = `</Types>`
    DefaultXmlTag = `<Default Extension="xml" ContentType="application/xml"/>`
    DefaultRelTag = `<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>`
    OverrideAppTag = `<Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/>`
    OverrideCoreTag = `<Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/>`
    OverrideWorkbookTag = `<Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/>`
)
