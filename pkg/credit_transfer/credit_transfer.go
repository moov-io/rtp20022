package credit_transfer

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/pkg/dt"
)

type Params struct {
	MessageID                   string
	CreatedOn                   time.Time
	SettlementAmount            int       // in cents
	SettlementDate              time.Time // YYYY-MM-DD
	InstructionID               string
	EndToEndID                  string
	TxID                        string
	SenderType                  SenderType
	DebtorRoutingNumber         string
	CreditorRoutingNumber       string
	DebtorName                  string
	DebtorAddress               Address
	DebtorAccountIdentification string // 17 alphanumeric chars
	CreditorName                string
	CreditorAddress             Address
	CreditorAccountNumber       string // 17 alphanumeric chars, probably has to be real account number
}

type SenderType string

var (
	SenderBusiness SenderType = SenderType("BUSINESS")
	SenderConsumer SenderType = SenderType("CONSUMER")
)

type Address struct {
	Street     string
	PostalCode string
	City       string
	State      string
	Country    string
}

type Pacs008 struct {
	XMLName xml.Name `xml:"CreditTransfer"`
	Text    string   `xml:",chardata"`

	FIToFICstmrCdtTrf pacs_008_001_08.FIToFICustomerCreditTransferV08 `xml:"FIToFICstmrCdtTrf"`
}

func (p *Pacs008) Validate() error {
	return p.FIToFICstmrCdtTrf.Validate()
}

const Namespace = "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"

func New(params Params) *Pacs008 {
	namespaceAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: Namespace,
	}
	xfer := pacs_008_001_08.FIToFICustomerCreditTransferV08{
		Attr: []xml.Attr{namespaceAttr},
		GrpHdr: pacs_008_001_08.GroupHeader93{
			MsgId:   pacs_008_001_08.Max35Text(params.MessageID),
			CreDtTm: dt.ISODateTime(params.CreatedOn),
			NbOfTxs: pacs_008_001_08.Max15NumericText("1"),
			TtlIntrBkSttlmAmt: &pacs_008_001_08.ActiveCurrencyAndAmount{
				Value: float64(params.SettlementAmount) / 100.0,
				Ccy:   "USD",
			},
			IntrBkSttlmDt: ptr(dt.ISODate(params.SettlementDate)),
			SttlmInf: pacs_008_001_08.SettlementInstruction7{
				SttlmMtd: pacs_008_001_08.SettlementMethod1Code("CLRG"),
				ClrSys: &pacs_008_001_08.ClearingSystemIdentification3Choice{
					Cd: ptr(pacs_008_001_08.ExternalCashClearingSystem1Code("TCH")),
				},
			},
		},
		CdtTrfTxInf: []pacs_008_001_08.CreditTransferTransaction39{
			{
				PmtId: pacs_008_001_08.PaymentIdentification7{
					InstrId:    ptr(pacs_008_001_08.Max35Text(params.InstructionID)),
					EndToEndId: pacs_008_001_08.Max35Text(params.EndToEndID),
					TxId:       ptr(pacs_008_001_08.Max35Text(params.TxID)),
				},
				PmtTpInf: &pacs_008_001_08.PaymentTypeInformation28{
					SvcLvl: []*pacs_008_001_08.ServiceLevel8Choice{
						{
							Cd: ptr(pacs_008_001_08.ExternalServiceLevel1Code("SDVA")),
						},
					},
					LclInstrm: &pacs_008_001_08.LocalInstrument2Choice{
						Prtry: ptr(pacs_008_001_08.Max35Text("STANDARD")),
					},
					CtgyPurp: &pacs_008_001_08.CategoryPurpose1Choice{
						Prtry: ptr(pacs_008_001_08.Max35Text(string(params.SenderType))),
					},
				},
				IntrBkSttlmAmt: pacs_008_001_08.ActiveCurrencyAndAmount{
					Value: float64(params.SettlementAmount) / 100.0,
					Ccy:   pacs_008_001_08.ActiveCurrencyCode("USD"),
				},
				ChrgBr: pacs_008_001_08.ChargeBearerType1Code("SLEV"),
				InstgAgt: &pacs_008_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_008_001_08.ClearingSystemMemberIdentification2{
							MmbId: pacs_008_001_08.Max35Text(params.DebtorRoutingNumber),
						},
					},
				},
				InstdAgt: &pacs_008_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_008_001_08.ClearingSystemMemberIdentification2{
							MmbId: pacs_008_001_08.Max35Text(params.CreditorRoutingNumber),
						},
					},
				},
				Dbtr: pacs_008_001_08.PartyIdentification135{
					Nm: ptr(pacs_008_001_08.Max140Text(params.DebtorName)),
					PstlAdr: &pacs_008_001_08.PostalAddress24{
						StrtNm:      ptr(pacs_008_001_08.Max70Text(params.DebtorAddress.Street)),
						PstCd:       ptr(pacs_008_001_08.Max16Text(params.DebtorAddress.PostalCode)),
						TwnNm:       ptr(pacs_008_001_08.Max35Text(params.DebtorAddress.City)),
						CtrySubDvsn: ptr(pacs_008_001_08.Max35Text(params.DebtorAddress.State)),
						Ctry:        ptr(pacs_008_001_08.CountryCode(params.DebtorAddress.Country)),
					},
				},
				DbtrAcct: &pacs_008_001_08.CashAccount38{
					Id: pacs_008_001_08.AccountIdentification4Choice{
						Othr: &pacs_008_001_08.GenericAccountIdentification1{
							Id: pacs_008_001_08.Max34Text(params.DebtorAccountIdentification),
						},
					},
				},
				DbtrAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_008_001_08.ClearingSystemMemberIdentification2{
							MmbId: pacs_008_001_08.Max35Text(params.DebtorRoutingNumber),
						},
					},
				},
				CdtrAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_008_001_08.ClearingSystemMemberIdentification2{
							MmbId: pacs_008_001_08.Max35Text(params.CreditorRoutingNumber),
						},
					},
				},
				Cdtr: pacs_008_001_08.PartyIdentification135{
					Nm: ptr(pacs_008_001_08.Max140Text(params.CreditorName)),
					PstlAdr: &pacs_008_001_08.PostalAddress24{
						StrtNm:      ptr(pacs_008_001_08.Max70Text(params.CreditorAddress.Street)),
						PstCd:       ptr(pacs_008_001_08.Max16Text(params.CreditorAddress.PostalCode)),
						TwnNm:       ptr(pacs_008_001_08.Max35Text(params.CreditorAddress.City)),
						CtrySubDvsn: ptr(pacs_008_001_08.Max35Text(params.CreditorAddress.State)),
						Ctry:        ptr(pacs_008_001_08.CountryCode(params.CreditorAddress.Country)),
					},
				},
				CdtrAcct: &pacs_008_001_08.CashAccount38{
					Id: pacs_008_001_08.AccountIdentification4Choice{
						Othr: &pacs_008_001_08.GenericAccountIdentification1{
							Id: pacs_008_001_08.Max34Text(params.CreditorAccountNumber),
						},
					},
				},
			},
		},
	}
	return &Pacs008{
		FIToFICstmrCdtTrf: xfer,
	}
}

func ptr[T any](v T) *T {
	return &v
}
