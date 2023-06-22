// Models for urn:tch
package messages

import (
	"encoding/xml"
	"fmt"

	"github.com/moov-io/rtp20022/gen/admi_002_001_01"
	"github.com/moov-io/rtp20022/gen/admi_004_001_02"
	"github.com/moov-io/rtp20022/gen/admn_001_001_01"
	"github.com/moov-io/rtp20022/gen/admn_002_001_01"
	"github.com/moov-io/rtp20022/gen/admn_003_001_01"
	"github.com/moov-io/rtp20022/gen/admn_004_001_01"
	"github.com/moov-io/rtp20022/gen/admn_005_001_01"
	"github.com/moov-io/rtp20022/gen/admn_006_001_01"
	"github.com/moov-io/rtp20022/gen/camt_029_001_09"
	"github.com/moov-io/rtp20022/gen/camt_035_001_05"
	"github.com/moov-io/rtp20022/gen/camt_056_001_08"
	"github.com/moov-io/rtp20022/gen/head_001_001_01"
	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/gen/pain_013_001_07"
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

// XSD ComplexType declarations

type HdrAndData struct {
	XMLName                 xml.Name                                     `xml:"Message"`
	Xmlns                   []xml.Attr                                   `xml:",attr"`
	AppHdr                  head_001_001_01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	CreditTransfer          *pacs_008_001_08.Document                    `xml:"CreditTransfer,omitempty"`
	MessageStatusReport     *pacs_002_001_10.Document                    `xml:"MessageStatusReport,omitempty"`
	Acknowledgement         *camt_035_001_05.Document                    `xml:"Acknowledgement,omitempty"`
	ReturnOfFunds           *camt_056_001_08.Document                    `xml:"ReturnOfFunds,omitempty"`
	PaymentRequest          *pain_013_001_07.Document                    `xml:"PaymentRequest,omitempty"`
	ResponseReturnOfFunds   *camt_029_001_09.Document                    `xml:"ResponseReturnOfFunds,omitempty"`
	EchoRequest             *admn_005_001_01.Document                    `xml:"EchoRequest,omitempty"`
	EchoResponse            *admn_006_001_01.Document                    `xml:"EchoResponse,omitempty"`
	SignOffRequest          *admn_003_001_01.Document                    `xml:"SignOffRequest,omitempty"`
	SignOffResponse         *admn_004_001_01.Document                    `xml:"SignOffResponse,omitempty"`
	SignOnRequest           *admn_001_001_01.Document                    `xml:"SignOnRequest,omitempty"`
	SignOnResponse          *admn_002_001_01.Document                    `xml:"SignOnResponse,omitempty"`
	SystemNotificationEvent *admi_004_001_02.Document                    `xml:"SystemNotificationEvent,omitempty"`
	MessageReject           *admi_002_001_01.Document                    `xml:"MessageReject,omitempty"`
}

func (v *HdrAndData) Body() interface{} {
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

// UnmarshalXML is a custom unmarshaller that allows us to capture the xmlns attributes
func (v *HdrAndData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	type vv HdrAndData
	return d.DecodeElement((*vv)(v), &start)
}

func (v *HdrAndData) Validate() error {
	return nil
}
