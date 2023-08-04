package messages

import (
	"encoding/xml"
	"fmt"
)

// UnmarshalXML is a custom unmarshaller that allows us to capture the xmlns attributes
func (v *Message) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if (attr.Name.Space == "" && attr.Name.Local == "xmlns") || (attr.Name.Space == "xmlns") {
			newAttr := xml.Attr{}
			newAttr.Value = attr.Value
			newAttr.Name = xml.Name{}
			if attr.Name.Space == "" {
				newAttr.Name.Local = attr.Name.Local
			} else {
				newAttr.Name.Local = fmt.Sprintf("%s:%s", attr.Name.Space, attr.Name.Local)
			}
			v.Xmlns = append(v.Xmlns, newAttr)
		}
	}

	// Go on with unmarshalling.
	type vv Message
	return d.DecodeElement((*vv)(v), &start)
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
