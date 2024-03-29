// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:tch
package messages

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/gen/acmt_022_001_02"
	"github.com/moov-io/rtp20022/gen/admi_002_001_01"
	"github.com/moov-io/rtp20022/gen/admi_004_001_02"
	"github.com/moov-io/rtp20022/gen/admn_001_001_01"
	"github.com/moov-io/rtp20022/gen/admn_002_001_01"
	"github.com/moov-io/rtp20022/gen/admn_003_001_01"
	"github.com/moov-io/rtp20022/gen/admn_004_001_01"
	"github.com/moov-io/rtp20022/gen/admn_005_001_01"
	"github.com/moov-io/rtp20022/gen/admn_006_001_01"
	"github.com/moov-io/rtp20022/gen/admn_007_001_01"
	"github.com/moov-io/rtp20022/gen/admn_008_001_01"
	"github.com/moov-io/rtp20022/gen/camt_026_001_07"
	"github.com/moov-io/rtp20022/gen/camt_028_001_09"
	"github.com/moov-io/rtp20022/gen/camt_029_001_09"
	"github.com/moov-io/rtp20022/gen/camt_035_001_05"
	"github.com/moov-io/rtp20022/gen/camt_056_001_08"
	"github.com/moov-io/rtp20022/gen/head_001_001_01"
	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/gen/pacs_009_001_08"
	"github.com/moov-io/rtp20022/gen/pacs_028_001_03"
	"github.com/moov-io/rtp20022/gen/pain_013_001_07"
	"github.com/moov-io/rtp20022/gen/pain_014_001_07"
	"github.com/moov-io/rtp20022/gen/remt_001_001_04"
)

// XSD Elements

type Message struct {
	XMLName                       xml.Name                                        `xml:"Message"`
	Xmlns                         []xml.Attr                                      `xml:",attr"`
	AppHdr                        head_001_001_01.BusinessApplicationHeaderV01TCH `xml:"urn:tch AppHdr"`
	CreditTransfer                *pacs_008_001_08.DocumentTCH                    `xml:"urn:tch CreditTransfer,omitempty"`
	MessageStatusReport           *pacs_002_001_10.DocumentTCH                    `xml:"urn:tch MessageStatusReport,omitempty"`
	FICreditTransfer              *pacs_009_001_08.DocumentTCH                    `xml:"urn:tch FICreditTransfer,omitempty"`
	Acknowledgement               *camt_035_001_05.DocumentTCH                    `xml:"urn:tch Acknowledgement,omitempty"`
	ResponseRequestForInformation *camt_028_001_09.DocumentTCH                    `xml:"urn:tch ResponseRequestForInformation,omitempty"`
	RequestForInformation         *camt_026_001_07.DocumentTCH                    `xml:"urn:tch RequestForInformation,omitempty"`
	ReturnOfFunds                 *camt_056_001_08.DocumentTCH                    `xml:"urn:tch ReturnOfFunds,omitempty"`
	PaymentRequest                *pain_013_001_07.DocumentTCH                    `xml:"urn:tch PaymentRequest,omitempty"`
	ResponsePaymentRequest        *pain_014_001_07.DocumentTCH                    `xml:"urn:tch ResponsePaymentRequest,omitempty"`
	ResponseReturnOfFunds         *camt_029_001_09.DocumentTCH                    `xml:"urn:tch ResponseReturnOfFunds,omitempty"`
	EchoRequest                   *admn_005_001_01.DocumentTCH                    `xml:"urn:tch EchoRequest,omitempty"`
	EchoResponse                  *admn_006_001_01.DocumentTCH                    `xml:"urn:tch EchoResponse,omitempty"`
	SignOffRequest                *admn_003_001_01.DocumentTCH                    `xml:"urn:tch SignOffRequest,omitempty"`
	SignOffResponse               *admn_004_001_01.DocumentTCH                    `xml:"urn:tch SignOffResponse,omitempty"`
	SignOnRequest                 *admn_001_001_01.DocumentTCH                    `xml:"urn:tch SignOnRequest,omitempty"`
	SignOnResponse                *admn_002_001_01.DocumentTCH                    `xml:"urn:tch SignOnResponse,omitempty"`
	StandaloneRemittance          *remt_001_001_04.DocumentTCH                    `xml:"urn:tch StandaloneRemittance,omitempty"`
	SystemNotificationEvent       *admi_004_001_02.DocumentTCH                    `xml:"urn:tch SystemNotificationEvent,omitempty"`
	MessageReject                 *admi_002_001_01.DocumentTCH                    `xml:"urn:tch MessageReject,omitempty"`
	TokenIdentification           *acmt_022_001_02.DocumentTCH                    `xml:"urn:tch TokenIdentification,omitempty"`
	ParticipantReport             *admn_007_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReport,omitempty"`
	ParticipantReportResponse     *admn_008_001_01.DocumentTCH                    `xml:"urn:tch ParticipantReportResponse,omitempty"`
	PaymentStatusRequest          *pacs_028_001_03.DocumentTCH                    `xml:"urn:tch PaymentStatusRequest,omitempty"`
}

var NamespacePrefixMap = map[string]string{
	"urn:iso:std:ma:20022:tech:xsd:admn.007.001.01":  "ut",
	"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01":  "tu",
	"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02": "a2",
	"urn:iso:std:iso:20022:tech:xsd:admi.002.001.01": "mr",
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02": "ne",
	"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01": "sr",
	"urn:iso:std:iso:20022:tech:xsd:admn.002.001.01": "rs",
	"urn:iso:std:iso:20022:tech:xsd:admn.003.001.01": "fr",
	"urn:iso:std:iso:20022:tech:xsd:admn.004.001.01": "rf",
	"urn:iso:std:iso:20022:tech:xsd:admn.005.001.01": "er",
	"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01": "re",
	"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07": "fi",
	"urn:iso:std:iso:20022:tech:xsd:camt.028.001.09": "if",
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09": "tr",
	"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05": "ac",
	"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08": "rt",
	"urn:iso:std:iso:20022:tech:xsd:head.001.001.01": "head",
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10": "ps",
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": "ct",
	"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08": "c9",
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03": "s8",
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07": "pr",
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.07": "rp",
	"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04": "ar",
}

func (v *Message) Body() interface{} {
	if v.CreditTransfer != nil {
		return v.CreditTransfer
	}
	if v.MessageStatusReport != nil {
		return v.MessageStatusReport
	}
	if v.FICreditTransfer != nil {
		return v.FICreditTransfer
	}
	if v.Acknowledgement != nil {
		return v.Acknowledgement
	}
	if v.ResponseRequestForInformation != nil {
		return v.ResponseRequestForInformation
	}
	if v.RequestForInformation != nil {
		return v.RequestForInformation
	}
	if v.ReturnOfFunds != nil {
		return v.ReturnOfFunds
	}
	if v.PaymentRequest != nil {
		return v.PaymentRequest
	}
	if v.ResponsePaymentRequest != nil {
		return v.ResponsePaymentRequest
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
	if v.StandaloneRemittance != nil {
		return v.StandaloneRemittance
	}
	if v.SystemNotificationEvent != nil {
		return v.SystemNotificationEvent
	}
	if v.MessageReject != nil {
		return v.MessageReject
	}
	if v.TokenIdentification != nil {
		return v.TokenIdentification
	}
	if v.ParticipantReport != nil {
		return v.ParticipantReport
	}
	if v.ParticipantReportResponse != nil {
		return v.ParticipantReportResponse
	}
	if v.PaymentStatusRequest != nil {
		return v.PaymentStatusRequest
	}
	return nil
}

// XSD ComplexType declarations
