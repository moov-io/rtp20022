// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 with prefix 'ps'
package pacs_002_001_10

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
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 FinInstnId"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MmbId"`
}

type DocumentTCH struct {
	XMLName         xml.Name
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV10TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 FIToFIPmtStsRpt"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysMmbId"`
}

type FIToFIPaymentStatusReportV10 struct {
	GrpHdr            GroupHeader91         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupHeader17 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlGrpInfAndSts"`
	TxInfAndSts       PaymentTransaction110 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxInfAndSts"`
}

type FIToFIPaymentStatusReportV10TCH struct {
	GrpHdr            GroupHeader91TCH         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupHeader17TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlGrpInfAndSts"`
	TxInfAndSts       PaymentTransaction110TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxInfAndSts"`
}

type GroupHeader91 struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CreDtTm"`
}

type GroupHeader91TCH struct {
	MsgId   Max35TextTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 CreDtTm"`
}

type OriginalGroupHeader17 struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlCreDtTm"`
	OrgnlNbOfTxs Max1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlNbOfTxs"`
}

type OriginalGroupHeader17TCH struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlCreDtTm"`
	OrgnlNbOfTxs Max1NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlNbOfTxs"`
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt  *rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 IntrBkSttlmDt,omitempty"`
}

type PaymentTransaction110 struct {
	OrgnlInstrId Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlInstrId"`
	OrgnlTxId    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlTxId,omitempty"`
	OrgnlUETR    *UUIDv4Identifier                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlUETR,omitempty"`
	TxSts        ExternalPaymentTransactionStatus1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxSts"`
	StsRsnInf    *StatusReasonInformation12                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 StsRsnInf,omitempty"`
	AccptncDtTm  rtp.ISODateTime                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AccptncDtTm"`
	ClrSysRef    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysRef,omitempty"`
	InstgAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstdAgt"`
	OrgnlTxRef   *OriginalTransactionReference28              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlTxRef,omitempty"`
}

type PaymentTransaction110TCH struct {
	OrgnlInstrId Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlInstrId"`
	OrgnlTxId    *Max35Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlTxId,omitempty"`
	OrgnlUETR    *UUIDv4Identifier                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlUETR,omitempty"`
	TxSts        ExternalPaymentTransactionStatus1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 TxSts"`
	StsRsnInf    *StatusReasonInformation12TCH                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 StsRsnInf,omitempty"`
	AccptncDtTm  rtp.ISODateTime                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AccptncDtTm"`
	ClrSysRef    *Max35Text                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 ClrSysRef,omitempty"`
	InstgAgt     BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 InstdAgt"`
	OrgnlTxRef   *OriginalTransactionReference28                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 OrgnlTxRef,omitempty"`
}

type StatusReason6Choice struct {
	Cd    *ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
	Prtry *ProprietaryReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Prtry,omitempty"`
}

type StatusReason6ChoiceTCH struct {
	Cd    *ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Cd,omitempty"`
	Prtry *ProprietaryReasonCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Prtry,omitempty"`
}

type StatusReasonInformation12 struct {
	Rsn      *StatusReason6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Rsn,omitempty"`
	AddtlInf *Max105Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AddtlInf,omitempty"`
}

type StatusReasonInformation12TCH struct {
	Rsn      *StatusReason6ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 Rsn,omitempty"`
	AddtlInf *Max105Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10 AddtlInf,omitempty"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType rtp.Amount

type ActiveOrHistoricCurrencyCode string

const ActiveOrHistoricCurrencyCodeUsd ActiveOrHistoricCurrencyCode = "USD"

type ExternalPaymentTransactionStatus1Code string

const ExternalPaymentTransactionStatus1CodeActc ExternalPaymentTransactionStatus1Code = "ACTC"
const ExternalPaymentTransactionStatus1CodeAcwp ExternalPaymentTransactionStatus1Code = "ACWP"
const ExternalPaymentTransactionStatus1CodeRcvd ExternalPaymentTransactionStatus1Code = "RCVD"
const ExternalPaymentTransactionStatus1CodeRjct ExternalPaymentTransactionStatus1Code = "RJCT"

type ExternalStatusReason1Code string

const ExternalStatusReason1CodeAc02 ExternalStatusReason1Code = "AC02"
const ExternalStatusReason1CodeAc03 ExternalStatusReason1Code = "AC03"
const ExternalStatusReason1CodeAc04 ExternalStatusReason1Code = "AC04"
const ExternalStatusReason1CodeAc06 ExternalStatusReason1Code = "AC06"
const ExternalStatusReason1CodeAc11 ExternalStatusReason1Code = "AC11"
const ExternalStatusReason1CodeAc13 ExternalStatusReason1Code = "AC13"
const ExternalStatusReason1CodeAc14 ExternalStatusReason1Code = "AC14"
const ExternalStatusReason1CodeAg01 ExternalStatusReason1Code = "AG01"
const ExternalStatusReason1CodeAg03 ExternalStatusReason1Code = "AG03"
const ExternalStatusReason1CodeAgnt ExternalStatusReason1Code = "AGNT"
const ExternalStatusReason1CodeAm02 ExternalStatusReason1Code = "AM02"
const ExternalStatusReason1CodeAm04 ExternalStatusReason1Code = "AM04"
const ExternalStatusReason1CodeAm09 ExternalStatusReason1Code = "AM09"
const ExternalStatusReason1CodeAm12 ExternalStatusReason1Code = "AM12"
const ExternalStatusReason1CodeAm13 ExternalStatusReason1Code = "AM13"
const ExternalStatusReason1CodeAm14 ExternalStatusReason1Code = "AM14"
const ExternalStatusReason1CodeBe04 ExternalStatusReason1Code = "BE04"
const ExternalStatusReason1CodeBe06 ExternalStatusReason1Code = "BE06"
const ExternalStatusReason1CodeBe07 ExternalStatusReason1Code = "BE07"
const ExternalStatusReason1CodeBe10 ExternalStatusReason1Code = "BE10"
const ExternalStatusReason1CodeBe11 ExternalStatusReason1Code = "BE11"
const ExternalStatusReason1CodeBe16 ExternalStatusReason1Code = "BE16"
const ExternalStatusReason1CodeBe17 ExternalStatusReason1Code = "BE17"
const ExternalStatusReason1CodeDs0H ExternalStatusReason1Code = "DS0H"
const ExternalStatusReason1CodeDs24 ExternalStatusReason1Code = "DS24"
const ExternalStatusReason1CodeDt04 ExternalStatusReason1Code = "DT04"
const ExternalStatusReason1CodeDupl ExternalStatusReason1Code = "DUPL"
const ExternalStatusReason1CodeFf02 ExternalStatusReason1Code = "FF02"
const ExternalStatusReason1CodeFf08 ExternalStatusReason1Code = "FF08"
const ExternalStatusReason1CodeMd07 ExternalStatusReason1Code = "MD07"
const ExternalStatusReason1CodeNarr ExternalStatusReason1Code = "NARR"
const ExternalStatusReason1CodeRc01 ExternalStatusReason1Code = "RC01"
const ExternalStatusReason1CodeRc02 ExternalStatusReason1Code = "RC02"
const ExternalStatusReason1CodeRc03 ExternalStatusReason1Code = "RC03"
const ExternalStatusReason1CodeRc04 ExternalStatusReason1Code = "RC04"
const ExternalStatusReason1CodeSl03 ExternalStatusReason1Code = "SL03"
const ExternalStatusReason1CodeTk01 ExternalStatusReason1Code = "TK01"
const ExternalStatusReason1CodeTk02 ExternalStatusReason1Code = "TK02"
const ExternalStatusReason1CodeTk03 ExternalStatusReason1Code = "TK03"
const ExternalStatusReason1CodeTk04 ExternalStatusReason1Code = "TK04"
const ExternalStatusReason1CodeTk05 ExternalStatusReason1Code = "TK05"
const ExternalStatusReason1CodeTk06 ExternalStatusReason1Code = "TK06"
const ExternalStatusReason1CodeTk07 ExternalStatusReason1Code = "TK07"
const ExternalStatusReason1CodeTk08 ExternalStatusReason1Code = "TK08"
const ExternalStatusReason1CodeTk09 ExternalStatusReason1Code = "TK09"
const ExternalStatusReason1CodeTm01 ExternalStatusReason1Code = "TM01"

type Max105Text string

type Max1NumericText string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type OrigMsgName string

const OrigMsgNameAcmt02200102 OrigMsgName = "acmt.022.001.02"
const OrigMsgNameCamt02600107 OrigMsgName = "camt.026.001.07"
const OrigMsgNameCamt02800109 OrigMsgName = "camt.028.001.09"
const OrigMsgNameCamt02900109 OrigMsgName = "camt.029.001.09"
const OrigMsgNameCamt03500105 OrigMsgName = "camt.035.001.05"
const OrigMsgNameCamt05600108 OrigMsgName = "camt.056.001.08"
const OrigMsgNamePacs00200110 OrigMsgName = "pacs.002.001.10"
const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePacs00900108 OrigMsgName = "pacs.009.001.08"
const OrigMsgNamePacs02800103 OrigMsgName = "pacs.028.001.03"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"
const OrigMsgNamePain01400107 OrigMsgName = "pain.014.001.07"
const OrigMsgNameRemt00100104 OrigMsgName = "remt.001.001.04"

type ProprietaryReasonCode string

const ProprietaryReasonCode1100 ProprietaryReasonCode = "1100"
const ProprietaryReasonCode9909 ProprietaryReasonCode = "9909"
const ProprietaryReasonCode9910 ProprietaryReasonCode = "9910"
const ProprietaryReasonCode9912 ProprietaryReasonCode = "9912"
const ProprietaryReasonCode9914 ProprietaryReasonCode = "9914"
const ProprietaryReasonCode9934 ProprietaryReasonCode = "9934"
const ProprietaryReasonCode9946 ProprietaryReasonCode = "9946"
const ProprietaryReasonCode9947 ProprietaryReasonCode = "9947"
const ProprietaryReasonCode9948 ProprietaryReasonCode = "9948"
const ProprietaryReasonCode9952 ProprietaryReasonCode = "9952"
const ProprietaryReasonCode9953 ProprietaryReasonCode = "9953"
const ProprietaryReasonCode9954 ProprietaryReasonCode = "9954"
const ProprietaryReasonCode9956 ProprietaryReasonCode = "9956"
const ProprietaryReasonCode9957 ProprietaryReasonCode = "9957"
const ProprietaryReasonCode9964 ProprietaryReasonCode = "9964"
const ProprietaryReasonCodeNoat ProprietaryReasonCode = "NOAT"

type UUIDv4Identifier string
