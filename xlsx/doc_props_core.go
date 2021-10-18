package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

func NewDocPropsCore() *DocPropsCore {
	dpc := new(DocPropsCore)
	dpc.CpNameSpace = "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"
	dpc.DcNameSpace = "http://purl.org/dc/elements/1.1/"
	dpc.DctNameSpace = "http://purl.org/dc/terms/"
	dpc.XsiNameSpace = "http://www.w3.org/2001/XMLSchema-instance"

	return dpc
}

type DocPropsCore struct {
	XMLName xml.Name `xml:"cp:coreProperties"`

	CpNameSpace  string `xml:"xmlns:cp,attr"`
	DcNameSpace  string `xml:"xmlns:dc,attr"`
	DctNameSpace string `xml:"xmlns:dcterms,attr"`
	XsiNameSpace string `xml:"xmlns:xsi,attr"`

	Creator    string `xml:"dc:creator"`
	LastModify string `xml:"cp:lastModifiedBy"`
	Created    DocPropsCoreCreated
	Modified   DocPropsCoreModified
}

func (d *DocPropsCore) SetUsername(user string) *DocPropsCore {
	d.Creator = user
	d.LastModify = user

	return d
}

func (d *DocPropsCore) SetTime(t time.Time) *DocPropsCore {
	d.Created = DocPropsCoreCreated{Type: "dcterms:W3CDTF", Created: t.Format("2006-01-02T15:04:05Z")}
	d.Modified = DocPropsCoreModified{Type: "dcterms:W3CDTF", Created: t.Format("2006-01-02T15:04:05Z")}

	return d
}

func (d *DocPropsCore) FilePath() string {
	return "/docProps/core.xml"
}

func (d *DocPropsCore) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(d); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}

type DocPropsCoreCreated struct {
	XMLName xml.Name `xml:"dcterms:created"`

	Type string `xml:"xsi:type,attr"`

	Created string `xml:",innerxml"`
}

type DocPropsCoreModified struct {
	XMLName xml.Name `xml:"dcterms:modified"`

	Type string `xml:"xsi:type,attr"`

	Created string `xml:",innerxml"`
}
