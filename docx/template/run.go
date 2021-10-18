package template

const (
	RunStart = `<w:r>`
	RunEnd   = `</w:r>`

	RunTextTag   = `w:t`
	RunTextSpace = `xml:space`

	BreakLine = `<w:br w:type="{{TYPE}}" w:clear="{{CLEAR}}"/>`

	RunRPrStart = `<w:rPr>`
	RunRPrEnd   = `</w:rPr>`

	RunRPrStyleTag = `w:rStyle`

	RunRPrVal             = `w:val`
	RunRPrBold            = `<w:b w:val="true"/>`
	RunRPrItalics         = `<w:i w:val="true"/>`
	RunRPrColorTag        = `w:color`
	RunRPrDStrike         = `<w:dstrike w:val="true"/>`
	RunRPrStrike          = `<w:strike w:val="true"/>`
	RunRPrEmboss          = `<w:emboss w:val="true"/>`
	RunRPrImprint         = `<w:imprint w:val="true"/>`
	RunRPrShadow          = `<w:shadow w:val="true"/>`
	RunRPrSizeTag         = `w:sz`
	RunRPrUnderlineTag    = `w:u`
	RunRPrUnderlineColor  = `w:color`
	RunRPrVanish          = `<vanish/>`
	RunPPrBackgroundTag   = `w:shd`
	RunPPrBackgroundFill  = `w:fill`
	RunPPrBackgroundColor = `w:color`
	RunRPrHighlightTag    = `w:highlight`
)
