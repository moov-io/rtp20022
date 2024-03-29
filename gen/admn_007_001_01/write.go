// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:ma:20022:tech:xsd:admn.007.001.01 with prefix 'ut'
package admn_007_001_01

import (
	"encoding/xml"
)

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification4ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "ut:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "ut:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatabaseReportInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RptCd, xml.StartElement{Name: xml.Name{Local: "ut:RptCd"}})
	e.EncodeElement(v.InstrId, xml.StartElement{Name: xml.Name{Local: "ut:InstrId"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "ut:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "ut:InstdAgt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatabaseReportInformationTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.RptCd, xml.StartElement{Name: xml.Name{Local: "ut:RptCd"}})
	e.EncodeElement(v.InstrId, xml.StartElement{Name: xml.Name{Local: "ut:InstrId"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "ut:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "ut:InstdAgt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatabaseReportRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "ut:GrpHdr"}})
	e.EncodeElement(v.DBRptInf, xml.StartElement{Name: xml.Name{Local: "ut:DBRptInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DatabaseReportRequestTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "ut:GrpHdr"}})
	e.EncodeElement(v.DBRptInf, xml.StartElement{Name: xml.Name{Local: "ut:DBRptInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.DBRptReq, xml.StartElement{Name: xml.Name{Local: "ut:DBRptReq"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification7ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "ut:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "ut:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "ut:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdrTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "ut:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "ut:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}
