package head_001_001_01

import (
	"encoding/xml"
)

type Sgntr struct {
	Signature *Signature
}

type Signature struct {
	XMLName        xml.Name       `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo     `xml:"SignedInfo"`
	SignatureValue SignatureValue `xml:"SignatureValue"`
	KeyInfo        KeyInfo        `xml:"KeyInfo"`
}

func NewSignature() *Sgntr {
	sgntr := &Sgntr{
		Signature: &Signature{
			XMLName: xml.Name{
				Space: "http://www.w3.org/2000/09/xmldsig#",
				Local: "Signature",
			},
			SignedInfo: SignedInfo{
				CanonicalizationMethod: CanonicalizationMethod{
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: SignatureMethod{
					Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
				},
				References: []Reference{
					{
						URI: "",
						Transforms: Transforms{
							Transforms: []Transform{
								{
									Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
								},
								{
									Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
								},
							},
						},
						DigestMethod: DigestMethod{
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
				},
			},
			SignatureValue: SignatureValue{},
			KeyInfo:        KeyInfo{},
		},
	}
	return sgntr
}

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	References             []Reference            `xml:"Reference"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type Reference struct {
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"Transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}

type Transforms struct {
	Transforms []Transform `xml:"Transform"`
}

type Transform struct {
	Algorithm           string               `xml:"Algorithm,attr"`
	InclusiveNamespaces *InclusiveNamespaces `xml:"InclusiveNamespaces"`
}

type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type InclusiveNamespaces struct {
	PrefixList string `xml:"PrefixList,attr"`
}

type SignatureValue struct {
	Data string `xml:",chardata"`
}

type KeyInfo struct {
	X509Data X509Data `xml:"X509Data"`
}

type X509Data struct {
	X509SubjectName  string           `xml:"X509SubjectName"`
	X509IssuerSerial X509IssuerSerial `xml:"X509IssuerSerial"`
}

type X509IssuerSerial struct {
	X509IssuerName   string `xml:"X509IssuerName"`
	X509SerialNumber string `xml:"X509SerialNumber"`
}
