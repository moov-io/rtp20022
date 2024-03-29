// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 with prefix 'fi'
package camt_026_001_07

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
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH2 struct {
	FinInstnId FinancialInstitutionIdentification18TCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 FinInstnId"`
}

type Case5 struct {
	Id    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Cretr Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cretr"`
}

type Case5TCH struct {
	Id    Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Cretr Party40ChoiceTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cretr"`
}

type CaseAssignment5 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Assgnr  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnr"`
	Assgne  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgne"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CreDtTm"`
}

type CaseAssignment5TCH struct {
	Id      Max35TextTCH     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Id"`
	Assgnr  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnr"`
	Assgne  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgne"`
	CreDtTm rtp.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 CreDtTm"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MmbId"`
}

type ClearingSystemMemberIdentification2TCH2 struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MmbId"`
}

type DateAndDateTime2Choice struct {
	Dt   *rtp.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Dt,omitempty"`
	DtTm *rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 DtTm,omitempty"`
}

type DocumentTCH struct {
	XMLName    xml.Name
	UblToApply UnableToApplyV07TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 UblToApply"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH2 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ClrSysMmbId"`
}

type MissingOrIncorrectInformation3 struct {
	MssngInf   []*UnableToApplyMissing1   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MssngInf,omitempty"`
	IncrrctInf []*UnableToApplyIncorrect1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 IncrrctInf,omitempty"`
}

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Agt,omitempty"`
}

type Party40ChoiceTCH struct {
	Agt *BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Agt,omitempty"`
}

type Party40ChoiceTCH2 struct {
	Pty *PartyIdentification135TCH                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6TCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Agt,omitempty"`
}

type PartyIdentification135 struct {
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Nm"`
}

type PartyIdentification135TCH struct {
	Nm Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Nm"`
}

type UnableToApplyIncorrect1 struct {
	Cd              UnableToApplyIncorrectInformation4Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd"`
	AddtlIncrrctInf *Max140Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AddtlIncrrctInf,omitempty"`
}

type UnableToApplyJustification3Choice struct {
	MssngOrIncrrctInf *MissingOrIncorrectInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 MssngOrIncrrctInf,omitempty"`
}

type UnableToApplyMissing1 struct {
	Cd            UnableToApplyMissingInformation3Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Cd"`
	AddtlMssngInf *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 AddtlMssngInf,omitempty"`
}

type UnableToApplyV07 struct {
	Assgnmt CaseAssignment5                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnmt"`
	Case    Case5                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Case"`
	Undrlyg UnderlyingTransaction5Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Undrlyg"`
	Justfn  UnableToApplyJustification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Justfn"`
}

type UnableToApplyV07TCH struct {
	Assgnmt CaseAssignment5TCH                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnmt"`
	Case    Case5TCH                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Case"`
	Undrlyg UnderlyingTransaction5ChoiceTCH   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Undrlyg"`
	Justfn  UnableToApplyJustification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Justfn"`
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlCreDtTm"`
}

type UnderlyingGroupInformation1TCH struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgNameTCH  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlCreDtTm"`
}

type UnderlyingGroupInformation1TCH2 struct {
	OrgnlMsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgNameTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlMsgNmId"`
	OrgnlCreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlCreDtTm"`
}

type UnderlyingPaymentInstruction5 struct {
	OrgnlGrpInf     UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlPmtInfId   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlPmtInfId"`
	OrgnlEndToEndId *Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       *UUIDv4Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstdAmt"`
	ReqdExctnDt     DateAndDateTime2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ReqdExctnDt"`
}

type UnderlyingPaymentInstruction5TCH struct {
	OrgnlGrpInf     UnderlyingGroupInformation1TCH    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlPmtInfId   Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlPmtInfId"`
	OrgnlEndToEndId *Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       *UUIDv4Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstdAmt"`
	ReqdExctnDt     DateAndDateTime2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 ReqdExctnDt"`
}

type UnderlyingPaymentTransaction4 struct {
	OrgnlGrpInf         UnderlyingGroupInformation1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlInstrId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstrId"`
	OrgnlEndToEndId     *Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlTxId"`
	OrgnlUETR           *UUIDv4Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmDt"`
}

type UnderlyingPaymentTransaction4TCH struct {
	OrgnlGrpInf         UnderlyingGroupInformation1TCH2   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlGrpInf"`
	OrgnlInstrId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlInstrId"`
	OrgnlEndToEndId     *Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlTxId"`
	OrgnlUETR           *UUIDv4Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlUETR,omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 OrgnlIntrBkSttlmDt"`
}

type UnderlyingTransaction5Choice struct {
	Initn  *UnderlyingPaymentInstruction5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Initn,omitempty"`
	IntrBk *UnderlyingPaymentTransaction4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 IntrBk,omitempty"`
}

type UnderlyingTransaction5ChoiceTCH struct {
	Initn  *UnderlyingPaymentInstruction5TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Initn,omitempty"`
	IntrBk *UnderlyingPaymentTransaction4TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 IntrBk,omitempty"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType rtp.Amount

type ActiveOrHistoricCurrencyCode string

const ActiveOrHistoricCurrencyCodeUsd ActiveOrHistoricCurrencyCode = "USD"

type Max140Text string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type OrigMsgName string

const OrigMsgNamePacs00800106 OrigMsgName = "pacs.008.001.06"
const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"

type OrigMsgNameTCH string

const OrigMsgNameTCHPain01300107 OrigMsgNameTCH = "pain.013.001.07"

type OrigMsgNameTCH2 string

const OrigMsgNameTCH2Pacs00800106 OrigMsgNameTCH2 = "pacs.008.001.06"
const OrigMsgNameTCH2Pacs00800108 OrigMsgNameTCH2 = "pacs.008.001.08"

type UnableToApplyIncorrectInformation4Code string

const UnableToApplyIncorrectInformation4CodeIn01 UnableToApplyIncorrectInformation4Code = "IN01"
const UnableToApplyIncorrectInformation4CodeIn04 UnableToApplyIncorrectInformation4Code = "IN04"
const UnableToApplyIncorrectInformation4CodeIn06 UnableToApplyIncorrectInformation4Code = "IN06"
const UnableToApplyIncorrectInformation4CodeIn15 UnableToApplyIncorrectInformation4Code = "IN15"
const UnableToApplyIncorrectInformation4CodeIn19 UnableToApplyIncorrectInformation4Code = "IN19"
const UnableToApplyIncorrectInformation4CodeIn38 UnableToApplyIncorrectInformation4Code = "IN38"
const UnableToApplyIncorrectInformation4CodeIn39 UnableToApplyIncorrectInformation4Code = "IN39"
const UnableToApplyIncorrectInformation4CodeMm20 UnableToApplyIncorrectInformation4Code = "MM20"
const UnableToApplyIncorrectInformation4CodeMm21 UnableToApplyIncorrectInformation4Code = "MM21"
const UnableToApplyIncorrectInformation4CodeNarr UnableToApplyIncorrectInformation4Code = "NARR"

type UnableToApplyMissingInformation3Code string

const UnableToApplyMissingInformation3CodeMs01 UnableToApplyMissingInformation3Code = "MS01"
const UnableToApplyMissingInformation3CodeNarr UnableToApplyMissingInformation3Code = "NARR"

type UUIDv4Identifier string
