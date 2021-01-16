package template

const Styles = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:styles xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
          xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
          mc:Ignorable="w14">
    
    
</w:styles>`

const (
    StyleXmlStart = `<w:styles xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">`

    StyleDocDefaultStart = `<w:docDefaults>`
    StyleDocDefaultEnd   = `</w:docDefaults>`

    StyleDefaultParagraphStart = `<w:pPrDefault>`
    StyleDefaultParagraphEnd   = `</w:pPrDefault>`

    StyleDefaultRunStart = `<w:rPrDefault>`
    StyleDefaultRunEnd   = `</w:rPrDefault>`

    StyleStyleTag     = `w:style`
    StyleStyleNameTag = `w:name`
    StyleStyleVal     = `w:val`
    StyleStyleType    = `w:type`
    StyleStyleStyleId = `w:styleId`

    StyleXmlEnd = `</w:styles>`
)
