package messages

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//SetParameterValuesResponse set param reponse
type SetParameterValuesResponse struct {
	ID           string
	Name         string
	Status       int
	ParameterKey string
}

type setParameterValuesResponseBodyStruct struct {
	Body setParameterValuesResponseStruct `xml:"cwmp:SetParameterValuesResponse"`
}

type setParameterValuesResponseStruct struct {
	Status       int
	ParameterKey string
}

//NewSetParameterValuesResponse create SetParameterValuesResponse object
func NewSetParameterValuesResponse() (m *SetParameterValuesResponse) {
	m = &SetParameterValuesResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetID get msg id
func (msg *SetParameterValuesResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg type
func (msg *SetParameterValuesResponse) GetName() string {
	return "SetParameterValuesResponse"
}

//CreateXML encode into xml
func (msg *SetParameterValuesResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id}
	setParamVal := setParameterValuesResponseStruct{
		Status:       msg.Status,
		ParameterKey: msg.ParameterKey,
	}
	env.Body = setParameterValuesResponseBodyStruct{setParamVal}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *SetParameterValuesResponse) Parse(doc *xmlx.Document) error {

	msg.ID = doc.SelectNode("*", "ID").GetValue()

	statusNode := doc.SelectNode("*", "Status")
	if statusNode != nil {
		var err error
		msg.Status, err = strconv.Atoi(statusNode.GetValue())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}

	paramsNode := doc.SelectNode("*", "ParameterKey")
	if paramsNode != nil {
		msg.ParameterKey = paramsNode.GetValue()
	}
	return nil
}
