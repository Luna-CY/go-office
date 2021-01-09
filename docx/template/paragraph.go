package template

const (
    ParagraphStart = `<w:p>`
    ParagraphEnd   = `</w:p>`

    ParagraphPPrStart               = `<w:pPr>`
    ParagraphPPrEnd                 = `</w:pPr>`
    ParagraphPPrHorizontalAlignment = `<w:jc w:val="{{TYPE}}"/>`
    ParagraphPPrBorderStart         = `<w:pBdr>`
    ParagraphPPrBorderEnd           = `</w:pBdr>`
    ParagraphPPrBorderTop           = `w:top`
    ParagraphPPrBorderRight         = `w:right`
    ParagraphPPrBorderBottom        = `w:bottom`
    ParagraphPPrBorderLeft          = `w:left`
    ParagraphPPrBorderStyle         = `w:val`
    ParagraphPPrBorderSize          = `w:sz`
    ParagraphPPrBorderSpace         = `w:space`
    ParagraphPPrBorderColor         = `w:color`
    ParagraphPPrBorderShadow        = `w:shadow`
    ParagraphPPrKeepLines           = `<w:keeplines/>`
    ParagraphPPrKeepNext            = `<w:keepNext/>`

    ParagraphPPrIdentityTag       = `w:ind`
    ParagraphPPrIdentityLeft      = `w:left`
    ParagraphPPrIdentityStart     = `w:start`
    ParagraphPPrIdentityRight     = `w:right`
    ParagraphPPrIdentityEnd       = `w:end`
    ParagraphPPrIdentityHanging   = `w:hanging`
    ParagraphPPrIdentityFirstLine = `w:firstLine`

    ParagraphPPrBackgroundTag   = `w:shd`
    ParagraphPPrBackgroundFill  = `w:fill`
    ParagraphPPrBackgroundColor = `w:color`
    ParagraphPPrBackgroundVal   = `w:val`
)
