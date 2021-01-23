package template

const (
    DocumentStart = `<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">`
    DocumentEnd   = `</w:document>`

    DocumentBodyStart = `<w:body>`
    DocumentBodyEnd   = `</w:body>`

    BackgroundTag        = `w:background`
    BackgroundColor      = `w:color`
    BackgroundThemeColor = `w:themeColor`
    BackgroundThemeShade = `w:themeShade`
    BackgroundThemeTint  = `w:themeTint`
)
