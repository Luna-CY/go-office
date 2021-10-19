package sheet

import "encoding/xml"

type Data struct {
	XMLName xml.Name `xml:"sheetData"`
}
