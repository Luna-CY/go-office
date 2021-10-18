package xlsx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

func NewDocPropsApp() *DocPropsApp {
	dpa := new(DocPropsApp)
	dpa.RootNameSpace = "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
	dpa.Application = "Golang Office Library"
	dpa.DocSecurity = 0
	dpa.AppVersion = "0.2.0"

	return dpa
}

type DocPropsApp struct {
	XMLName xml.Name `xml:"Properties"`

	RootNameSpace string `xml:"xmlns,attr"`

	Application string `xml:"Application"`
	DocSecurity int    `xml:"DocSecurity"`
	AppVersion  string `xml:"AppVersion"`
}

func (d *DocPropsApp) FilePath() string {
	return "/docProps/app.xml"
}

func (d *DocPropsApp) Marshal() ([]byte, error) {
	buffer := &bytes.Buffer{}

	ec := xml.NewEncoder(buffer)
	if err := ec.Encode(d); nil != err {
		return nil, errors.New(fmt.Sprintf("序列化XML结构失败: %v", err))
	}

	return buffer.Bytes(), nil
}
