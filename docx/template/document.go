package template

const DocumentStart = `<w:document xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" mc:Ignorable="w14 w15 wp14"><w:body>`
const DocumentEnd = `</w:body></w:document>`

const ParagraphStart = `<w:p>`
const ParagraphEnd = `</w:p>`

const ParagraphPPrStart = `<w:pPr>`
const ParagraphPPrEnd = `</w:pPr>`

const HorizontalAlignment = `<w:jc w:val="{{TYPE}}"/>`

const RunStart = `<w:r><w:t>`
const RunEnd = `</w:t></w:r>`

const BreakLine = `<w:br w:type="{{TYPE}}" w:clear="" />`
