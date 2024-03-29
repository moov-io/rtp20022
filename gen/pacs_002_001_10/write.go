// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 with prefix 'ps'
package pacs_002_001_10

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "ps:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "ps:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "ps:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "ps:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FIToFIPmtStsRpt, xml.StartElement{Name: xml.Name{Local: "ps:FIToFIPmtStsRpt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "ps:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "ps:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentStatusReportV10) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "ps:GrpHdr"}})
	e.EncodeElement(v.OrgnlGrpInfAndSts, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlGrpInfAndSts"}})
	e.EncodeElement(v.TxInfAndSts, xml.StartElement{Name: xml.Name{Local: "ps:TxInfAndSts"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentStatusReportV10TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "ps:GrpHdr"}})
	e.EncodeElement(v.OrgnlGrpInfAndSts, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlGrpInfAndSts"}})
	e.EncodeElement(v.TxInfAndSts, xml.StartElement{Name: xml.Name{Local: "ps:TxInfAndSts"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader91) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "ps:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "ps:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader91TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "ps:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "ps:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupHeader17) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlCreDtTm"}})
	e.EncodeElement(v.OrgnlNbOfTxs, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlNbOfTxs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupHeader17TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlCreDtTm"}})
	e.EncodeElement(v.OrgnlNbOfTxs, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlNbOfTxs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalTransactionReference28) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "ps:IntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "ps:IntrBkSttlmDt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlTxId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlUETR"}})
	e.EncodeElement(v.TxSts, xml.StartElement{Name: xml.Name{Local: "ps:TxSts"}})
	e.EncodeElement(v.StsRsnInf, xml.StartElement{Name: xml.Name{Local: "ps:StsRsnInf"}})
	e.EncodeElement(v.AccptncDtTm, xml.StartElement{Name: xml.Name{Local: "ps:AccptncDtTm"}})
	e.EncodeElement(v.ClrSysRef, xml.StartElement{Name: xml.Name{Local: "ps:ClrSysRef"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "ps:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "ps:InstdAgt"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction110TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlTxId"}})
	e.EncodeElement(v.OrgnlUETR, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlUETR"}})
	e.EncodeElement(v.TxSts, xml.StartElement{Name: xml.Name{Local: "ps:TxSts"}})
	e.EncodeElement(v.StsRsnInf, xml.StartElement{Name: xml.Name{Local: "ps:StsRsnInf"}})
	e.EncodeElement(v.AccptncDtTm, xml.StartElement{Name: xml.Name{Local: "ps:AccptncDtTm"}})
	e.EncodeElement(v.ClrSysRef, xml.StartElement{Name: xml.Name{Local: "ps:ClrSysRef"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "ps:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "ps:InstdAgt"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "ps:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReason6Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "ps:Cd"}})
	e.EncodeElement(v.Prtry, xml.StartElement{Name: xml.Name{Local: "ps:Prtry"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReason6ChoiceTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Cd, xml.StartElement{Name: xml.Name{Local: "ps:Cd"}})
	e.EncodeElement(v.Prtry, xml.StartElement{Name: xml.Name{Local: "ps:Prtry"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReasonInformation12) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "ps:Rsn"}})
	e.EncodeElement(v.AddtlInf, xml.StartElement{Name: xml.Name{Local: "ps:AddtlInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v StatusReasonInformation12TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Rsn, xml.StartElement{Name: xml.Name{Local: "ps:Rsn"}})
	e.EncodeElement(v.AddtlInf, xml.StartElement{Name: xml.Name{Local: "ps:AddtlInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return rtp.Amount(a).MarshalText()
}
