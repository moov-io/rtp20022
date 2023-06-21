package test

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/gen/head_001_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
)

func NewMessage() *messages.HdrAndData {
	message := &messages.HdrAndData{}
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

func NewSignature() *head_001_001_01.Sgntr {
	sgntr := &head_001_001_01.Sgntr{
		Signature: &head_001_001_01.Signature{
			XMLName: xml.Name{
				Space: "http://www.w3.org/2000/09/xmldsig#",
				Local: "Signature",
			},
			SignedInfo: head_001_001_01.SignedInfo{
				CanonicalizationMethod: head_001_001_01.CanonicalizationMethod{
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: head_001_001_01.SignatureMethod{
					Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
				},
				References: []head_001_001_01.Reference{
					{
						URI: "",
						Transforms: head_001_001_01.Transforms{
							Transforms: []head_001_001_01.Transform{
								{
									Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
								},
								{
									Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
								},
							},
						},
						DigestMethod: head_001_001_01.DigestMethod{
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
				},
			},
			SignatureValue: head_001_001_01.SignatureValue{},
			KeyInfo:        head_001_001_01.KeyInfo{},
		},
	}
	return sgntr
}

func NewAdmi002Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:mr",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admi.002.001.01",
	})
	return message
}

func NewAdmi004Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:ne",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
	})
	return message
}

func NewCamt029Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:tr",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09",
	})
	return message
}

func NewCamt035Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:ac",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:camt.035.001.05",
	})
	return message
}

func NewCamt056Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:rt",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:camt.056.001.08",
	})
	return message
}

func NewPacs002Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:ps",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
	})
	return message
}

func NewPacs008Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:ct",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
	})
	return message
}

func NewPain013Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:pr",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07",
	})
	return message
}

func NewAdmn001Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:sr",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.001.001.01",
	})
	return message
}

func NewAdmn002Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:rs",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.002.001.01",
	})
	return message
}

func NewAdmn003Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:fr",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.003.001.01",
	})
	return message
}

func NewAdmn004Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:rf",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.004.001.01",
	})
	return message
}

func NewAdmn005Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:er",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.005.001.01",
	})
	return message
}

func NewAdmn006Message() *messages.HdrAndData {
	message := NewMessage()
	message.Xmlns = append(message.Xmlns, xml.Attr{
		Name: xml.Name{
			Local: "xmlns:re",
		},
		Value: "urn:iso:std:iso:20022:tech:xsd:admn.006.001.01",
	})
	return message
}
