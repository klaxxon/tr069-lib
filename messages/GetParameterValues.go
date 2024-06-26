package messages

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//GetParameterValues get paramvalues
type GetParameterValues struct {
	ID             string
	Name           string
	NoMore         int
	ParameterNames []string
}

type getParameterValuesBodyStruct struct {
	Body getParameterValuesStruct `xml:"cwmp:GetParameterValues"`
}

type getParameterValuesStruct struct {
	Params parameterNamesStruct `xml:"ParameterNames"`
}

type parameterNamesStruct struct {
	Type       string   `xml:"SOAP-ENC:arrayType,attr"`
	ParamNames []string `xml:"string"`
}

//NewGetParameterValues create GetParameterValues object
func NewGetParameterValues() (m *GetParameterValues) {
	m = &GetParameterValues{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get type name
func (msg *GetParameterValues) GetName() string {
	return "GetParameterValues"
}

//GetID get tr069 msg id
func (msg *GetParameterValues) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *GetParameterValues) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	paramLen := strconv.Itoa(len(msg.ParameterNames))
	paramNames := parameterNamesStruct{
		Type: XsdString + "[" + paramLen + "]",
	}
	for _, v := range msg.ParameterNames {
		paramNames.ParamNames = append(paramNames.ParamNames, v)
	}
	getParam := getParameterValuesStruct{paramNames}
	env.Body = getParameterValuesBodyStruct{getParam}
	//output, err := xml.MarshalIndent(env, "  ", "    ")
	output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *GetParameterValues) Parse(doc *xmlx.Document) error {
	msg.ID = doc.SelectNode("*", "ID").GetValue()
	paramsNode := doc.SelectNode("*", "ParameterNames")
	if len(strings.TrimSpace(paramsNode.String())) > 0 {
		params := make([]string, 0)
		var name string
		for _, param := range paramsNode.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name = param.GetValue()
				params = append(params, name)
			}

		}
		msg.ParameterNames = params
	}
	return nil
}
