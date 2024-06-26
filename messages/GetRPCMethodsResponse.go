package messages

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//GetRPCMethodsResponse getRPCMethods reponse
type GetRPCMethodsResponse struct {
	ID      string
	Name    string
	Methods []string
}

type getRPCMethodsResponseBodyStruct struct {
	Body getRPCMethodsResponseStruct `xml:"cwmp:GetRPCMethodsResponse"`
}

type getRPCMethodsResponseStruct struct {
	MethodList methodListStruct `xml:"cwmp:MethodList"`
}

type methodListStruct struct {
	Type      string   `xml:"xsi:type,attr"`
	ArrayType string   `xml:"SOAP-ENC:arrayType,attr"`
	Methods   []string `xml:"string"`
}

//NewGetRPCMethodsResponse create GetRPCMethodsResponse object
func NewGetRPCMethodsResponse() (m *GetRPCMethodsResponse) {
	m = &GetRPCMethodsResponse{}
	m.ID = m.GetID()
	m.Name = m.GetName()
	return m
}

//GetName get msg type
func (msg *GetRPCMethodsResponse) GetName() string {
	return "GetRPCMethodsResponse"
}

//GetID get msg id
func (msg *GetRPCMethodsResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//CreateXML encode into xml
func (msg *GetRPCMethodsResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	id := IDStruct{"1", msg.GetID()}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	env.Header = HeaderStruct{ID: id}
	methodsLen := strconv.Itoa(len(msg.Methods))
	methodList := methodListStruct{
		Type:      SoapArray,
		ArrayType: XsdString + "[" + methodsLen + "]",
	}
	for _, v := range msg.Methods {
		methodList.Methods = append(methodList.Methods, v)
	}
	rpcMethods := getRPCMethodsResponseStruct{methodList}
	env.Body = getRPCMethodsResponseBodyStruct{rpcMethods}
	output, err := xml.MarshalIndent(env, "  ", "    ")
	//output, err := xml.Marshal(env)
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *GetRPCMethodsResponse) Parse(doc *xmlx.Document) error {

	msg.ID = doc.SelectNode("*", "ID").GetValue()

	methodList := doc.SelectNode("*", "MethodList")
	if len(strings.TrimSpace(methodList.String())) > 0 {
		var name string
		var methods []string
		for _, param := range methodList.Children {
			if len(strings.TrimSpace(param.String())) > 0 {
				name = param.GetValue()
				methods = append(methods, name)
			}
		}
		msg.Methods = methods
	}
	return nil
}
