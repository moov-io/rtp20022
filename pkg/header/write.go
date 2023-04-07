package header

import (
	"encoding/xml"
	"time"
)

type Writer struct {
	XMLName xml.Name `xml:"AppHdr"`
	Text    string   `xml:",chardata"`
	Head    string   `xml:"xmlns:head,attr"`
	Fr      struct {
		Head string `xml:"xmlns:head,attr"`
		Text string `xml:",chardata"`
		FIId struct {
			Text       string `xml:",chardata"`
			FinInstnId struct {
				Text        string `xml:",chardata"`
				ClrSysMmbId struct {
					Text  string `xml:",chardata"`
					MmbId string `xml:"head:MmbId"`
				} `xml:"head:ClrSysMmbId"`
			} `xml:"head:FinInstnId"`
		} `xml:"head:FIId"`
	} `xml:"head:Fr"`
	To struct {
		Head string `xml:"xmlns:head,attr"`
		Text string `xml:",chardata"`
		FIId struct {
			Text       string `xml:",chardata"`
			FinInstnId struct {
				Text        string `xml:",chardata"`
				ClrSysMmbId struct {
					Text  string `xml:",chardata"`
					MmbId string `xml:"head:MmbId"`
				} `xml:"head:ClrSysMmbId"`
			} `xml:"head:FinInstnId"`
		} `xml:"head:FIId"`
	} `xml:"head:To"`
	BizMsgIdr struct {
		Head string `xml:"xmlns:head,attr"`
		Text string `xml:",chardata"`
	} `xml:"head:BizMsgIdr"`
	MsgDefIdr struct {
		Head string `xml:"xmlns:head,attr"`
		Text string `xml:",chardata"`
	} `xml:"head:MsgDefIdr"`
	CreDt struct {
		Head string `xml:"xmlns:head,attr"`
		Text string `xml:",chardata"`
	} `xml:"head:CreDt"`
	CpyDplct *CpyDplct `xml:"head:CpyDplct,omitempty"`
	Sgntr    *Sgntr    `xml:"head:Sgntr"`
}

type CpyDplct struct {
	Head string  `xml:"xmlns:head,attr"`
	Text *string `xml:",chardata"`
}

type WriteParams struct {
	FromMemberID      string
	ToMemberID        string
	BusinessID        string
	MessageDefinition string
	CreatedOn         time.Time
	Duplicate         bool
}

const (
	DefaultDateTimeFormat = "2006-01-02T15:04:05"

	Namespace = "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"
)

func NewWriter(params WriteParams) Writer {
	headAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns:head"},
		Value: Namespace,
	}

	w := Writer{}
	w.Head = headAttr.Value

	w.Fr.Head = headAttr.Value
	w.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId = params.FromMemberID
	w.To.Head = headAttr.Value
	w.To.FIId.FinInstnId.ClrSysMmbId.MmbId = params.ToMemberID

	w.BizMsgIdr.Head = headAttr.Value
	w.BizMsgIdr.Text = params.BusinessID

	w.MsgDefIdr.Head = headAttr.Value
	w.MsgDefIdr.Text = params.MessageDefinition

	w.CreDt.Head = headAttr.Value
	w.CreDt.Text = params.CreatedOn.Format(DefaultDateTimeFormat)

	if params.Duplicate {
		dupl := "DUPL"
		w.CpyDplct = &CpyDplct{
			Head: headAttr.Value,
			Text: &dupl,
		}
	}

	// Setup XML Signature
	w.Sgntr = &Sgntr{
		Signature: &Signature{
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
	return w
}
