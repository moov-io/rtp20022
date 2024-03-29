// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 with prefix 'a2'
package acmt_022_001_02

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Othr,omitempty"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification5TCH struct {
	FinInstnId FinancialInstitutionIdentification8TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 FinInstnId"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MmbId"`
}

type DocumentTCH struct {
	XMLName   xml.Name
	IdModAdvc IdentificationModificationAdviceV02TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 IdModAdvc"`
}

type FinancialInstitutionIdentification8 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 ClrSysMmbId"`
}

type FinancialInstitutionIdentification8TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 ClrSysMmbId"`
}

type GenericAccountIdentification1 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Id"`
}

type IdentificationAssignment2 struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 CreDtTm"`
	Assgnr  Party12Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgnr"`
	Assgne  Party12Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgne"`
}

type IdentificationAssignment2TCH struct {
	MsgId   Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgId"`
	CreDtTm rtp.ISODateTime  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 CreDtTm"`
	Assgnr  Party12ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgnr"`
	Assgne  Party12ChoiceTCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgne"`
}

type IdentificationInformation2 struct {
	Acct AccountIdentification4Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Acct"`
	Agt  BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Agt"`
}

type IdentificationInformation2TCH struct {
	Acct AccountIdentification4Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Acct"`
	Agt  BranchAndFinancialInstitutionIdentification5TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Agt"`
}

type IdentificationModification2 struct {
	Id                Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Id"`
	OrgnlPtyAndAcctId *IdentificationInformation2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlPtyAndAcctId,omitempty"`
	UpdtdPtyAndAcctId IdentificationInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 UpdtdPtyAndAcctId"`
	AddtlInf          *Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 AddtlInf,omitempty"`
}

type IdentificationModification2TCH struct {
	Id                Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Id"`
	OrgnlPtyAndAcctId *IdentificationInformation2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlPtyAndAcctId,omitempty"`
	UpdtdPtyAndAcctId IdentificationInformation2TCH  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 UpdtdPtyAndAcctId"`
	AddtlInf          *Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 AddtlInf,omitempty"`
}

type IdentificationModificationAdviceV02 struct {
	Assgnmt    IdentificationAssignment2      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgnmt"`
	OrgnlTxRef OriginalTransactionReference18 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlTxRef"`
	Mod        IdentificationModification2    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Mod"`
}

type IdentificationModificationAdviceV02TCH struct {
	Assgnmt    IdentificationAssignment2TCH      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Assgnmt"`
	OrgnlTxRef OriginalTransactionReference18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlTxRef"`
	Mod        IdentificationModification2TCH    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Mod"`
}

type OriginalTransactionReference18 struct {
	MsgId   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgId"`
	MsgNmId OrigMsgName            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgNmId"`
	CreDtTm rtp.ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 CreDtTm"`
	OrgnlTx PaymentIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlTx"`
}

type OriginalTransactionReference18TCH struct {
	MsgId   Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgId"`
	MsgNmId OrigMsgName               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 MsgNmId"`
	CreDtTm rtp.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 CreDtTm"`
	OrgnlTx PaymentIdentification4TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 OrgnlTx"`
}

type Party12Choice struct {
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Agt,omitempty"`
}

type Party12ChoiceTCH struct {
	Agt *BranchAndFinancialInstitutionIdentification5TCH `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 Agt,omitempty"`
}

type PaymentIdentification4 struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 InstrId"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 EndToEndId"`
	TxId       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 TxId"`
}

type PaymentIdentification4TCH struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 InstrId"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 EndToEndId"`
	TxId       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 TxId"`
}

// XSD SimpleType declarations

type Max140Text string

type Max34Text string

type Max35Text string

type Max35TextTCH string

type OrigMsgName string

const OrigMsgNamePacs00800108 OrigMsgName = "pacs.008.001.08"
const OrigMsgNamePain01300107 OrigMsgName = "pain.013.001.07"
