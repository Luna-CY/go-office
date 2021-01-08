package template

const (
    DocumentStart = `<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:conformance="strict"><w:body>`
    DocumentEnd   = `</w:body></w:document>`

    BackgroundTag        = `w:background`
    BackgroundColor      = `w:color`
    BackgroundThemeColor = `w:themeColor`
    BackgroundThemeShade = `w:themeShade`
    BackgroundThemeTint  = `w:themeTint`

    ParagraphStart = `<w:p>`
    ParagraphEnd   = `</w:p>`

    ParagraphPPrStart   = `<w:pPr>`
    ParagraphPPrEnd     = `</w:pPr>`
    HorizontalAlignment = `<w:jc w:val="{{TYPE}}"/>`

    RunStart = `<w:r><w:t>`
    RunEnd   = `</w:t></w:r>`

    BreakLine = `<w:br w:type="{{TYPE}}" w:clear="" />`
)
