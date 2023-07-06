package messages

import (
	"encoding/xml"
	"fmt"
)

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01": "mr",
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02": "ne",
	"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01": "sr",
	"urn:iso:std:iso:20022:tech:xsd:admn.002.001.01": "rs",
	"urn:iso:std:iso:20022:tech:xsd:admn.003.001.01": "fr",
	"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01": "rf",
	"urn:iso:std:iso:20022:tech:xsd:admn.005.001.01": "er",
	"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01": "re",
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09": "tr",
	"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05": "ac",
	"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08": "rt",
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10": "ps",
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": "ct",
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07": "pr",
}

func (v *Message) Body() interface{} {
	if v.CreditTransfer != nil {
		return v.CreditTransfer
	}
	if v.MessageStatusReport != nil {
		return v.MessageStatusReport
	}
	if v.Acknowledgement != nil {
		return v.Acknowledgement
	}
	if v.ReturnOfFunds != nil {
		return v.ReturnOfFunds
	}
	if v.PaymentRequest != nil {
		return v.PaymentRequest
	}
	if v.ResponseReturnOfFunds != nil {
		return v.ResponseReturnOfFunds
	}
	if v.EchoRequest != nil {
		return v.EchoRequest
	}
	if v.EchoResponse != nil {
		return v.EchoResponse
	}
	if v.SignOffRequest != nil {
		return v.SignOffRequest
	}
	if v.SignOffResponse != nil {
		return v.SignOffResponse
	}
	if v.SignOnRequest != nil {
		return v.SignOnRequest
	}
	if v.SignOnResponse != nil {
		return v.SignOnResponse
	}
	if v.SystemNotificationEvent != nil {
		return v.SystemNotificationEvent
	}
	if v.MessageReject != nil {
		return v.MessageReject
	}
	return nil
}

func (v *Message) Validate() error {
	return nil
}

func newMessage() *Message {
	message := &Message{}
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns",
		},
		Value: "urn:tch",
	})
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:head",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01",
	})
	return message
}

func newMessageForNS(namespace string) *Message {
	message := newMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: fmt.Sprintf("xmlns:%s", NamespacePrefixMap[namespace]),
		},
		Value: namespace,
	})
	return message
}

func NewAdmi002Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admi.002.001.01")
}

func NewAdmi004Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admi.004.001.02")
}

func NewCamt029Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.029.001.09")
}

func NewCamt035Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.035.001.05")
}

func NewCamt056Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.056.001.08")
}

func NewPacs002Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10")
}

func NewPacs008Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08")
}

func NewPain013Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pain.013.001.07")
}

func NewAdmn001Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.001.001.01")
}

func NewAdmn002Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.002.001.01")
}

func NewAdmn003Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.003.001.01")
}

func NewAdmn004Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.004.001.01")
}

func NewAdmn005Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.005.001.01")
}

func NewAdmn006Message() *Message {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.006.001.01")
}
