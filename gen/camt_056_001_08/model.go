// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 with prefix 'rt'
package camt_056_001_08

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
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 FinInstnId"`
}

type CancellationReason33Choice struct {
	Cd    *ExternalCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cd,omitempty"`
	Prtry *Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Prtry,omitempty"`
}

type CancellationReason33ChoiceTCH struct {
	Cd    *ExternalCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cd,omitempty"`
	Prtry *Max35TextTCH3                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Prtry,omitempty"`
}

type Case5 struct {
	Id    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
	Cretr Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cretr"`
}

type Case5TCH struct {
	Id    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
	Cretr Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cretr"`
}

type CaseAssignment5 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
	Assgnr  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgnr"`
	Assgne  Party40Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgne"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CreDtTm"`
}

type CaseAssignment5TCH struct {
	Id      Max35TextTCH     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
	Assgnr  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgnr"`
	Assgne  Party40ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgne"`
	CreDtTm rtp.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CreDtTm"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 MmbId"`
}

type Contact4 struct {
	PhneNb *PhoneNumber `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PhneNb,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     rtp.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 BirthDt"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtryOfBirth"`
}

type DocumentTCH struct {
	XMLName         xml.Name
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV08TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 FIToFIPmtCxlReq"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 ClrSysMmbId"`
}

type FIToFIPaymentCancellationRequestV08 struct {
	Assgnmt CaseAssignment5         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgnmt"`
	Case    Case5                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Case"`
	Undrlyg UnderlyingTransaction23 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Undrlyg"`
}

type FIToFIPaymentCancellationRequestV08TCH struct {
	Assgnmt CaseAssignment5TCH         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgnmt"`
	Case    Case5TCH                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Case"`
	Undrlyg UnderlyingTransaction23TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Undrlyg"`
}

type GenericOrganisationIdentification1 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
}

type GenericOrganisationIdentification1TCH struct {
	Id Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id"`
}

type OrganisationIdentification29 struct {
	LEI  *LEIIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 LEI,omitempty"`
	Othr *GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Othr,omitempty"`
}

type OrganisationIdentification29TCH struct {
	Othr GenericOrganisationIdentification1TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Othr"`
}

type OrganisationIdentification29TCH2 struct {
	LEI LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 LEI"`
}

type OriginalGroupHeader15 struct {
	OrgnlMsgId   Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlMsgId"`
	OrgnlMsgNmId OrigMsgName      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlMsgNmId"`
	OrgnlCreDtTm *rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlCreDtTm,omitempty"`
}

type OriginalTransactionReference28 struct {
	Dbtr *Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Dbtr,omitempty"`
	Cdtr *Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cdtr,omitempty"`
}

type OriginalTransactionReference28TCH struct {
	Dbtr *Party40ChoiceTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Dbtr,omitempty"`
	Cdtr *Party40ChoiceTCH3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Cdtr,omitempty"`
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PrvtId,omitempty"`
}

type Party38ChoiceTCH struct {
	OrgId *OrganisationIdentification29TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgId,omitempty"`
}

type Party38ChoiceTCH2 struct {
	OrgId  *OrganisationIdentification29TCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgId,omitempty"`
	PrvtId *PersonIdentification13TCH        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PrvtId,omitempty"`
}

type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Agt,omitempty"`
}

type Party40ChoiceTCH struct {
	Agt *BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Agt,omitempty"`
}

type Party40ChoiceTCH2 struct {
	Pty *PartyIdentification135TCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Pty,omitempty"`
}

type Party40ChoiceTCH3 struct {
	Pty *PartyIdentification135TCH3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Pty,omitempty"`
}

type PartyIdentification135 struct {
	Nm       *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Nm,omitempty"`
	PstlAdr  *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PstlAdr,omitempty"`
	Id       *Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id,omitempty"`
	CtctDtls *Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtctDtls,omitempty"`
}

type PartyIdentification135TCH struct {
	Nm *Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Nm,omitempty"`
	Id *Party38ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id,omitempty"`
}

type PartyIdentification135TCH2 struct {
	Nm      Max140Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Nm"`
	PstlAdr *PostalAddress24TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PstlAdr,omitempty"`
	Id      *Party38ChoiceTCH2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id,omitempty"`
}

type PartyIdentification135TCH3 struct {
	Nm       Max140Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Nm"`
	PstlAdr  *PostalAddress24TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PstlAdr,omitempty"`
	Id       *Party38ChoiceTCH2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Id,omitempty"`
	CtctDtls *Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtctDtls,omitempty"`
}

type PaymentCancellationReason5 struct {
	Orgtr    *PartyIdentification135    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Orgtr,omitempty"`
	Rsn      CancellationReason33Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Rsn"`
	AddtlInf *Max105Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 AddtlInf,omitempty"`
}

type PaymentCancellationReason5TCH struct {
	Orgtr    *PartyIdentification135TCH    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Orgtr,omitempty"`
	Rsn      CancellationReason33ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Rsn"`
	AddtlInf *Max105Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 AddtlInf,omitempty"`
}

type PaymentTransaction106 struct {
	OrgnlInstrId        Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlInstrId"`
	OrgnlEndToEndId     *Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlTxId"`
	OrgnlUETR           *UUIDv4Identifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlUETR,omitempty"`
	OrgnlClrSysRef      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlClrSysRef"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  rtp.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlIntrBkSttlmDt"`
	CxlRsnInf           PaymentCancellationReason5        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CxlRsnInf"`
	OrgnlTxRef          *OriginalTransactionReference28   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlTxRef,omitempty"`
}

type PaymentTransaction106TCH struct {
	OrgnlInstrId        Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlInstrId"`
	OrgnlEndToEndId     *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlTxId"`
	OrgnlUETR           *UUIDv4Identifier                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlUETR,omitempty"`
	OrgnlClrSysRef      Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlClrSysRef"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  rtp.ISODate                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlIntrBkSttlmDt"`
	CxlRsnInf           PaymentCancellationReason5TCH      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CxlRsnInf"`
	OrgnlTxRef          *OriginalTransactionReference28TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlTxRef,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 DtAndPlcOfBirth"`
}

type PersonIdentification13TCH struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 DtAndPlcOfBirth"`
}

type PostalAddress24 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 StrtNm"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 BldgNb,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 AdrLine,omitempty"`
}

type PostalAddress24TCH struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 StrtNm"`
	BldgNb      *Max16Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 BldgNb,omitempty"`
	PstCd       Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 PstCd"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 TwnNm"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtrySubDvsn"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Ctry"`
	AdrLine     *Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 AdrLine,omitempty"`
}

type UnderlyingTransaction23 struct {
	OrgnlGrpInfAndCxl OriginalGroupHeader15 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlGrpInfAndCxl"`
	TxInf             PaymentTransaction106 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 TxInf"`
}

type UnderlyingTransaction23TCH struct {
	OrgnlGrpInfAndCxl OriginalGroupHeader15    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 OrgnlGrpInfAndCxl"`
	TxInf             PaymentTransaction106TCH `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 TxInf"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType rtp.Amount

type ActiveOrHistoricCurrencyCode string

const ActiveOrHistoricCurrencyCodeUsd ActiveOrHistoricCurrencyCode = "USD"

type CountryCode string

const CountryCodeAd CountryCode = "AD"
const CountryCodeAe CountryCode = "AE"
const CountryCodeAf CountryCode = "AF"
const CountryCodeAg CountryCode = "AG"
const CountryCodeAi CountryCode = "AI"
const CountryCodeAl CountryCode = "AL"
const CountryCodeAm CountryCode = "AM"
const CountryCodeAo CountryCode = "AO"
const CountryCodeAq CountryCode = "AQ"
const CountryCodeAr CountryCode = "AR"
const CountryCodeAs CountryCode = "AS"
const CountryCodeAt CountryCode = "AT"
const CountryCodeAu CountryCode = "AU"
const CountryCodeAw CountryCode = "AW"
const CountryCodeAx CountryCode = "AX"
const CountryCodeAz CountryCode = "AZ"
const CountryCodeBa CountryCode = "BA"
const CountryCodeBb CountryCode = "BB"
const CountryCodeBd CountryCode = "BD"
const CountryCodeBe CountryCode = "BE"
const CountryCodeBf CountryCode = "BF"
const CountryCodeBg CountryCode = "BG"
const CountryCodeBh CountryCode = "BH"
const CountryCodeBi CountryCode = "BI"
const CountryCodeBj CountryCode = "BJ"
const CountryCodeBl CountryCode = "BL"
const CountryCodeBm CountryCode = "BM"
const CountryCodeBn CountryCode = "BN"
const CountryCodeBo CountryCode = "BO"
const CountryCodeBq CountryCode = "BQ"
const CountryCodeBr CountryCode = "BR"
const CountryCodeBs CountryCode = "BS"
const CountryCodeBt CountryCode = "BT"
const CountryCodeBv CountryCode = "BV"
const CountryCodeBw CountryCode = "BW"
const CountryCodeBy CountryCode = "BY"
const CountryCodeBz CountryCode = "BZ"
const CountryCodeCa CountryCode = "CA"
const CountryCodeCc CountryCode = "CC"
const CountryCodeCd CountryCode = "CD"
const CountryCodeCf CountryCode = "CF"
const CountryCodeCg CountryCode = "CG"
const CountryCodeCh CountryCode = "CH"
const CountryCodeCi CountryCode = "CI"
const CountryCodeCk CountryCode = "CK"
const CountryCodeCl CountryCode = "CL"
const CountryCodeCm CountryCode = "CM"
const CountryCodeCn CountryCode = "CN"
const CountryCodeCo CountryCode = "CO"
const CountryCodeCr CountryCode = "CR"
const CountryCodeCu CountryCode = "CU"
const CountryCodeCv CountryCode = "CV"
const CountryCodeCw CountryCode = "CW"
const CountryCodeCx CountryCode = "CX"
const CountryCodeCy CountryCode = "CY"
const CountryCodeCz CountryCode = "CZ"
const CountryCodeDe CountryCode = "DE"
const CountryCodeDj CountryCode = "DJ"
const CountryCodeDk CountryCode = "DK"
const CountryCodeDm CountryCode = "DM"
const CountryCodeDo CountryCode = "DO"
const CountryCodeDz CountryCode = "DZ"
const CountryCodeEc CountryCode = "EC"
const CountryCodeEe CountryCode = "EE"
const CountryCodeEg CountryCode = "EG"
const CountryCodeEh CountryCode = "EH"
const CountryCodeEr CountryCode = "ER"
const CountryCodeEs CountryCode = "ES"
const CountryCodeEt CountryCode = "ET"
const CountryCodeFi CountryCode = "FI"
const CountryCodeFj CountryCode = "FJ"
const CountryCodeFk CountryCode = "FK"
const CountryCodeFm CountryCode = "FM"
const CountryCodeFo CountryCode = "FO"
const CountryCodeFr CountryCode = "FR"
const CountryCodeGa CountryCode = "GA"
const CountryCodeGb CountryCode = "GB"
const CountryCodeGd CountryCode = "GD"
const CountryCodeGe CountryCode = "GE"
const CountryCodeGf CountryCode = "GF"
const CountryCodeGg CountryCode = "GG"
const CountryCodeGh CountryCode = "GH"
const CountryCodeGi CountryCode = "GI"
const CountryCodeGl CountryCode = "GL"
const CountryCodeGm CountryCode = "GM"
const CountryCodeGn CountryCode = "GN"
const CountryCodeGp CountryCode = "GP"
const CountryCodeGq CountryCode = "GQ"
const CountryCodeGr CountryCode = "GR"
const CountryCodeGs CountryCode = "GS"
const CountryCodeGt CountryCode = "GT"
const CountryCodeGu CountryCode = "GU"
const CountryCodeGw CountryCode = "GW"
const CountryCodeGy CountryCode = "GY"
const CountryCodeHk CountryCode = "HK"
const CountryCodeHm CountryCode = "HM"
const CountryCodeHn CountryCode = "HN"
const CountryCodeHr CountryCode = "HR"
const CountryCodeHt CountryCode = "HT"
const CountryCodeHu CountryCode = "HU"
const CountryCodeId CountryCode = "ID"
const CountryCodeIe CountryCode = "IE"
const CountryCodeIl CountryCode = "IL"
const CountryCodeIm CountryCode = "IM"
const CountryCodeIn CountryCode = "IN"
const CountryCodeIo CountryCode = "IO"
const CountryCodeIq CountryCode = "IQ"
const CountryCodeIr CountryCode = "IR"
const CountryCodeIs CountryCode = "IS"
const CountryCodeIt CountryCode = "IT"
const CountryCodeJe CountryCode = "JE"
const CountryCodeJm CountryCode = "JM"
const CountryCodeJo CountryCode = "JO"
const CountryCodeJp CountryCode = "JP"
const CountryCodeKe CountryCode = "KE"
const CountryCodeKg CountryCode = "KG"
const CountryCodeKh CountryCode = "KH"
const CountryCodeKi CountryCode = "KI"
const CountryCodeKm CountryCode = "KM"
const CountryCodeKn CountryCode = "KN"
const CountryCodeKp CountryCode = "KP"
const CountryCodeKr CountryCode = "KR"
const CountryCodeKw CountryCode = "KW"
const CountryCodeKy CountryCode = "KY"
const CountryCodeKz CountryCode = "KZ"
const CountryCodeLa CountryCode = "LA"
const CountryCodeLb CountryCode = "LB"
const CountryCodeLc CountryCode = "LC"
const CountryCodeLi CountryCode = "LI"
const CountryCodeLk CountryCode = "LK"
const CountryCodeLr CountryCode = "LR"
const CountryCodeLs CountryCode = "LS"
const CountryCodeLt CountryCode = "LT"
const CountryCodeLu CountryCode = "LU"
const CountryCodeLv CountryCode = "LV"
const CountryCodeLy CountryCode = "LY"
const CountryCodeMa CountryCode = "MA"
const CountryCodeMc CountryCode = "MC"
const CountryCodeMd CountryCode = "MD"
const CountryCodeMe CountryCode = "ME"
const CountryCodeMf CountryCode = "MF"
const CountryCodeMg CountryCode = "MG"
const CountryCodeMh CountryCode = "MH"
const CountryCodeMk CountryCode = "MK"
const CountryCodeMl CountryCode = "ML"
const CountryCodeMm CountryCode = "MM"
const CountryCodeMn CountryCode = "MN"
const CountryCodeMo CountryCode = "MO"
const CountryCodeMp CountryCode = "MP"
const CountryCodeMq CountryCode = "MQ"
const CountryCodeMr CountryCode = "MR"
const CountryCodeMs CountryCode = "MS"
const CountryCodeMt CountryCode = "MT"
const CountryCodeMu CountryCode = "MU"
const CountryCodeMv CountryCode = "MV"
const CountryCodeMw CountryCode = "MW"
const CountryCodeMx CountryCode = "MX"
const CountryCodeMy CountryCode = "MY"
const CountryCodeMz CountryCode = "MZ"
const CountryCodeNa CountryCode = "NA"
const CountryCodeNc CountryCode = "NC"
const CountryCodeNe CountryCode = "NE"
const CountryCodeNf CountryCode = "NF"
const CountryCodeNg CountryCode = "NG"
const CountryCodeNi CountryCode = "NI"
const CountryCodeNl CountryCode = "NL"
const CountryCodeNo CountryCode = "NO"
const CountryCodeNp CountryCode = "NP"
const CountryCodeNr CountryCode = "NR"
const CountryCodeNu CountryCode = "NU"
const CountryCodeNz CountryCode = "NZ"
const CountryCodeOm CountryCode = "OM"
const CountryCodePa CountryCode = "PA"
const CountryCodePe CountryCode = "PE"
const CountryCodePf CountryCode = "PF"
const CountryCodePg CountryCode = "PG"
const CountryCodePh CountryCode = "PH"
const CountryCodePk CountryCode = "PK"
const CountryCodePl CountryCode = "PL"
const CountryCodePm CountryCode = "PM"
const CountryCodePn CountryCode = "PN"
const CountryCodePr CountryCode = "PR"
const CountryCodePs CountryCode = "PS"
const CountryCodePt CountryCode = "PT"
const CountryCodePw CountryCode = "PW"
const CountryCodePy CountryCode = "PY"
const CountryCodeQa CountryCode = "QA"
const CountryCodeRe CountryCode = "RE"
const CountryCodeRo CountryCode = "RO"
const CountryCodeRs CountryCode = "RS"
const CountryCodeRu CountryCode = "RU"
const CountryCodeRw CountryCode = "RW"
const CountryCodeSa CountryCode = "SA"
const CountryCodeSb CountryCode = "SB"
const CountryCodeSc CountryCode = "SC"
const CountryCodeSd CountryCode = "SD"
const CountryCodeSe CountryCode = "SE"
const CountryCodeSg CountryCode = "SG"
const CountryCodeSh CountryCode = "SH"
const CountryCodeSi CountryCode = "SI"
const CountryCodeSj CountryCode = "SJ"
const CountryCodeSk CountryCode = "SK"
const CountryCodeSl CountryCode = "SL"
const CountryCodeSm CountryCode = "SM"
const CountryCodeSn CountryCode = "SN"
const CountryCodeSo CountryCode = "SO"
const CountryCodeSr CountryCode = "SR"
const CountryCodeSs CountryCode = "SS"
const CountryCodeSt CountryCode = "ST"
const CountryCodeSv CountryCode = "SV"
const CountryCodeSx CountryCode = "SX"
const CountryCodeSy CountryCode = "SY"
const CountryCodeSz CountryCode = "SZ"
const CountryCodeTc CountryCode = "TC"
const CountryCodeTd CountryCode = "TD"
const CountryCodeTf CountryCode = "TF"
const CountryCodeTg CountryCode = "TG"
const CountryCodeTh CountryCode = "TH"
const CountryCodeTj CountryCode = "TJ"
const CountryCodeTk CountryCode = "TK"
const CountryCodeTl CountryCode = "TL"
const CountryCodeTm CountryCode = "TM"
const CountryCodeTn CountryCode = "TN"
const CountryCodeTo CountryCode = "TO"
const CountryCodeTr CountryCode = "TR"
const CountryCodeTt CountryCode = "TT"
const CountryCodeTv CountryCode = "TV"
const CountryCodeTw CountryCode = "TW"
const CountryCodeTz CountryCode = "TZ"
const CountryCodeUa CountryCode = "UA"
const CountryCodeUg CountryCode = "UG"
const CountryCodeUm CountryCode = "UM"
const CountryCodeUs CountryCode = "US"
const CountryCodeUy CountryCode = "UY"
const CountryCodeUz CountryCode = "UZ"
const CountryCodeVa CountryCode = "VA"
const CountryCodeVc CountryCode = "VC"
const CountryCodeVe CountryCode = "VE"
const CountryCodeVg CountryCode = "VG"
const CountryCodeVi CountryCode = "VI"
const CountryCodeVn CountryCode = "VN"
const CountryCodeVu CountryCode = "VU"
const CountryCodeWf CountryCode = "WF"
const CountryCodeWs CountryCode = "WS"
const CountryCodeYe CountryCode = "YE"
const CountryCodeYt CountryCode = "YT"
const CountryCodeZa CountryCode = "ZA"
const CountryCodeZm CountryCode = "ZM"
const CountryCodeZw CountryCode = "ZW"

type ExternalCancellationReason1Code string

const ExternalCancellationReason1CodeAc03 ExternalCancellationReason1Code = "AC03"
const ExternalCancellationReason1CodeAm09 ExternalCancellationReason1Code = "AM09"
const ExternalCancellationReason1CodeCust ExternalCancellationReason1Code = "CUST"
const ExternalCancellationReason1CodeDs24 ExternalCancellationReason1Code = "DS24"
const ExternalCancellationReason1CodeDupl ExternalCancellationReason1Code = "DUPL"
const ExternalCancellationReason1CodeFrad ExternalCancellationReason1Code = "FRAD"
const ExternalCancellationReason1CodeFrtr ExternalCancellationReason1Code = "FRTR"
const ExternalCancellationReason1CodeTech ExternalCancellationReason1Code = "TECH"
const ExternalCancellationReason1CodeUpay ExternalCancellationReason1Code = "UPAY"

type LEIIdentifier string

type Max105Text string

type Max140Text string

type Max16Text string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type Max35TextTCH3 string

const Max35TextTCH3Wiam Max35TextTCH3 = "WIAM"
const Max35TextTCH3Wict Max35TextTCH3 = "WICT"
const Max35TextTCH3Widp Max35TextTCH3 = "WIDP"
const Max35TextTCH3Wifd Max35TextTCH3 = "WIFD"
const Max35TextTCH3Wift Max35TextTCH3 = "WIFT"
const Max35TextTCH3With Max35TextTCH3 = "WITH"

type Max70Text string

type OrigMsgName string

const OrigMsgNamePacs00800106 OrigMsgName = "pacs.008.001.06"
const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePacs00900108 OrigMsgName = "pacs.009.001.08"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"

type PhoneNumber string

type UUIDv4Identifier string
