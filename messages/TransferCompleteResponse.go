package messages

import (
	"encoding/xml"
	"fmt"
	"time"

	xmlx "github.com/mattn/go-pkg-xmlx"
)

//TransferCompleteResponse transferComplete response
type TransferCompleteResponse struct {
	ID     string
	Name   string
	NoMore int
}

type transferCompleteResponseBodyStruct struct {
	Body transferCompleteResponseStruct `xml:"cwmp:TransferCompleteResponse"`
}

type transferCompleteResponseStruct struct {
}

//GetID get msg id
func (msg *TransferCompleteResponse) GetID() string {
	if len(msg.ID) < 1 {
		msg.ID = fmt.Sprintf("ID:intrnl.unset.id.%s%d.%d", msg.GetName(), time.Now().Unix(), time.Now().UnixNano())
	}
	return msg.ID
}

//GetName get msg name
func (msg *TransferCompleteResponse) GetName() string {
	return "TransferCompleteResponse"
}

//CreateXML encode into xml
func (msg *TransferCompleteResponse) CreateXML() ([]byte, error) {
	env := Envelope{}
	env.XmlnsEnv = "http://schemas.xmlsoap.org/soap/envelope/"
	env.XmlnsEnc = "http://schemas.xmlsoap.org/soap/encoding/"
	env.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	env.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.XmlnsCwmp = "urn:dslforum-org:cwmp-1-0"
	id := IDStruct{Attr: "1", Value: msg.GetID()}
	env.Header = HeaderStruct{ID: id, NoMore: msg.NoMore}
	transf := transferCompleteResponseStruct{}
	env.Body = transferCompleteResponseBodyStruct{transf}
	//output, err := xml.Marshal(env)
	output, err := xml.MarshalIndent(env, "  ", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

//Parse decode from xml
func (msg *TransferCompleteResponse) Parse(doc *xmlx.Document) error {
	//TODO
	return nil
}
