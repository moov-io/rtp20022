// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 with prefix 's8'
package pacs_028_001_03

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 FinInstnId"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MmbId"`
}

type DocumentTCH struct {
	XMLName         xml.Name
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV03TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 FIToFIPmtStsReq"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 ClrSysMmbId"`
}

type FIToFIPaymentStatusRequestV03 struct {
	GrpHdr      GroupHeader91              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 GrpHdr"`
	OrgnlGrpInf OriginalGroupInformation27 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlGrpInf"`
	TxInf       PaymentTransaction113      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 TxInf"`
}

type FIToFIPaymentStatusRequestV03TCH struct {
	GrpHdr      GroupHeader91TCH              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 GrpHdr"`
	OrgnlGrpInf OriginalGroupInformation27TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlGrpInf"`
	TxInf       PaymentTransaction113TCH      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 TxInf"`
}

type GroupHeader91 struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 CreDtTm"`
}

type GroupHeader91TCH struct {
	MsgId   Max35TextTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 CreDtTm"`
}

type OriginalGroupInformation27 struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlCreDtTm"`
	OrgnlNbOfTxs Max1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlNbOfTxs"`
}

type OriginalGroupInformation27TCH struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlCreDtTm"`
	OrgnlNbOfTxs Max1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlNbOfTxs"`
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 IntrBkSttlmAmt"`
	IntrBkSttlmDt  rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 IntrBkSttlmDt"`
}

type OriginalTransactionReference28TCH struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 IntrBkSttlmAmt"`
	IntrBkSttlmDt  rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 IntrBkSttlmDt"`
}

type PaymentTransaction113 struct {
	OrgnlInstrId Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlInstrId"`
	OrgnlTxId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlTxId"`
	AccptncDtTm  rtp.ISODateTime                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 AccptncDtTm"`
	InstgAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstdAgt"`
	OrgnlTxRef   OriginalTransactionReference28               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlTxRef"`
}

type PaymentTransaction113TCH struct {
	OrgnlInstrId Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlInstrId"`
	OrgnlTxId    Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlTxId"`
	AccptncDtTm  rtp.ISODateTime                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 AccptncDtTm"`
	InstgAgt     BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 InstdAgt"`
	OrgnlTxRef   OriginalTransactionReference28TCH               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 OrgnlTxRef"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType rtp.Amount

type ActiveOrHistoricCurrencyCode string

const ActiveOrHistoricCurrencyCodeUsd ActiveOrHistoricCurrencyCode = "USD"

type Max1NumericText string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type OrigMsgName string

const OrigMsgNamePacs00800106 OrigMsgName = "pacs.008.001.06"
const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePacs00900108 OrigMsgName = "pacs.009.001.08"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"
