package messages

import (
	"encoding/xml"
	"fmt"
)

func newMessage() *HdrAndData {
	message := &HdrAndData{}
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

func newMessageForNS(namespace string) *HdrAndData {
	message := newMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: fmt.Sprintf("xmlns:%s", NamespacePrefixMap[namespace]),
		},
		Value: namespace,
	})
	return message
}

func NewAdmi002Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admi.002.001.01")
}

func NewAdmi004Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admi.004.001.02")
}

func NewCamt029Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.029.001.09")
}

func NewCamt035Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.035.001.05")
}

func NewCamt056Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:camt.056.001.08")
}

func NewPacs002Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10")
}

func NewPacs008Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08")
}

func NewPain013Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:pain.013.001.07")
}

func NewAdmn001Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.001.001.01")
}

func NewAdmn002Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.002.001.01")
}

func NewAdmn003Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.003.001.01")
}

func NewAdmn004Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.004.001.01")
}

func NewAdmn005Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.005.001.01")
}

func NewAdmn006Message() *HdrAndData {
	return newMessageForNS("urn:iso:std:iso:20022:tech:xsd:admn.006.001.01")
}
