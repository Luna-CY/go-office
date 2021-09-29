package template

const (
	ParagraphStart  = `<w:p>`
	ParagraphEnd    = `</w:p>`
	ParagraphSingle = `<w:p/>`

	ParagraphPPrStart = `<w:pPr>`
	ParagraphPPrEnd   = `</w:pPr>`

	ParagraphPPrStyleTag = `w:pStyle`
	ParagraphPPrStyleVal = `w:val`

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

	ParagraphPPrIndentationTag       = `w:ind`
	ParagraphPPrIndentationLeft      = `w:left`
	ParagraphPPrIndentationStart     = `w:start`
	ParagraphPPrIndentationRight     = `w:right`
	ParagraphPPrIndentationEnd       = `w:end`
	ParagraphPPrIndentationHanging   = `w:hanging`
	ParagraphPPrIndentationFirstLine = `w:firstLine`

	ParagraphPPrBackgroundTag   = `w:shd`
	ParagraphPPrBackgroundFill  = `w:fill`
	ParagraphPPrBackgroundColor = `w:color`
	ParagraphPPrBackgroundVal   = `w:val`

	ParagraphPPrSpacingTag               = `w:spacing`
	ParagraphPPrSpacingAfter             = `w:after`
	ParagraphPPrSpacingBefore            = `w:before`
	ParagraphPPrSpacingLine              = `w:line`
	ParagraphPPrSpacingLineRule          = `w:lineRule`
	ParagraphPPrSpacingBeforeAutoSpacing = `w:beforeAutospacing`
	ParagraphPPrSpacingAfterAutoSpacing  = `w:afterAutospacing`
)
