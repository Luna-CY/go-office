package template

const (
    DocumentStart = `<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:conformance="strict">`
    DocumentEnd   = `</w:document>`

    DocumentBodyStart = `<w:body>`
    DocumentBodyEnd   = `</w:body>`

    BackgroundTag        = `w:background`
    BackgroundColor      = `w:color`
    BackgroundThemeColor = `w:themeColor`
    BackgroundThemeShade = `w:themeShade`
    BackgroundThemeTint  = `w:themeTint`

    RunStart = `<w:r>`
    RunEnd   = `</w:r>`

    RunTextStart = `<w:t>`
    RunTextEnd = `</w:t>`

    BreakLine = `<w:br w:type="{{TYPE}}" w:clear="{{CLEAR}}"/>`
)
