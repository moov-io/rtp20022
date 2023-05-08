package msg_status

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/pkg/dt"
)

type Writer struct {
	XMLName         xml.Name `xml:"MessageStatusReport"`
	Text            string   `xml:",chardata"`
	Head            string   `xml:"xmlns:ps,attr"`
	FIToFIPmtStsRpt struct {
		Head   string `xml:"xmlns:ps,attr"`
		Text   string `xml:",chardata"`
		GrpHdr struct {
			Head    string                    `xml:"xmlns:ps,attr"`
			Text    string                    `xml:",chardata"`
			MsgId   pacs_002_001_10.Max35Text `xml:"ps:MsgId"`
			CreDtTm dt.ISODateTime            `xml:"ps:CreDtTm"`
		} `xml:"ps:GrpHdr"`
		OrgnlGrpInfAndSts []OrgnlGrpInfAndSts `xml:"ps:OrgnlGrpInfAndSts,omitempty"`
		TxInfAndSts       []TxInfAndSts       `xml:"ps:TxInfAndSts,omitempty"`
	} `xml:"ps:FIToFIPmtStsRpt"`
}

type TxInfAndSts struct {
	Head         string                                                `xml:"xmlns:ps,attr"`
	Text         string                                                `xml:",chardata"`
	OrgnlInstrId pacs_002_001_10.Max35Text                             `xml:"ps:OrgnlInstrId"`
	OrgnlTxId    pacs_002_001_10.Max35Text                             `xml:"ps:OrgnlTxId"`
	TxSts        pacs_002_001_10.ExternalPaymentTransactionStatus1Code `xml:"ps:TxSts,omitempty"`
	StsRsnInf    []StsRsnInf                                           `xml:"ps:StsRsnInf,omitempty"`
	AccptncDtTm  dt.ISODateTime                                        `xml:"ps:AccptncDtTm"`
	ClrSysRef    string                                                `xml:"ps:ClrSysRef,omitempty"`
	InstgAgt     BranchAndFinancialInstitutionIdentification6          `xml:"ps:InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification6          `xml:"ps:InstdAgt"`
	OrgnlTxRef   *OriginalTransactionReference28                       `xml:"ps:OrgnlTxRef,omitempty"`
}

type StsRsnInf struct {
	Rsn StatusReason6Choice `xml:"ps:Rsn"`
}

type StatusReason6Choice struct {
	Cd pacs_002_001_10.ExternalStatusReason1Code `xml:"ps:Cd"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"ps:FinInstnId"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ps:ClrSysMmbId,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId pacs_002_001_10.Max35Text `xml:"ps:MmbId"`
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"ps:IntrBkSttlmAmt"`
	IntrBkSttlmDt  dt.ISODate                        `xml:"ps:IntrBkSttlmDt"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyCode string

type OrgnlGrpInfAndSts struct {
	Head         string                    `xml:"xmlns:ps,attr"`
	Text         string                    `xml:",chardata"`
	OrgnlMsgId   pacs_002_001_10.Max35Text `xml:"ps:OrgnlMsgId"`
	OrgnlMsgNmId pacs_002_001_10.Max35Text `xml:"ps:OrgnlMsgNmId"`
	OrgnlCreDtTm dt.ISODateTime            `xml:"ps:OrgnlCreDtTm"`
	OrgnlNbOfTxs int                       `xml:"ps:OrgnlNbOfTxs"`
}

func (r Writer) Validate() error {
	return nil // TODO(adam):
}

// WriteParams is a human-readable struct of the Writer fields
type WriteParams struct {
	MessageID                    string
	CreatedOn                    time.Time
	OriginalInstructionId        string
	OriginalTxId                 string
	InstructingAgentMemberId     string
	InstructedAgentMemberId      string
	OriginalMessageId            string
	OriginalMessageNameId        string
	OriginalCreatedOn            time.Time
	OriginalNumberOfTransactions int
	TxSts                        TransactionStatus
	ClrSysRef                    string
	StsRsnCd                     string
	AcceptanceDateTime           time.Time
	InterbankSettlementAmount    *float64
	InterbankSettlementAmountCcy string
	InterbankSettlementDate      *time.Time
}

const (
	Namespace = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"
)

func NewWriter(params WriteParams) Writer {
	headAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns:ps"},
		Value: Namespace,
	}

	w := Writer{}
	w.Head = headAttr.Value

	w.FIToFIPmtStsRpt.Head = headAttr.Value

	w.FIToFIPmtStsRpt.GrpHdr.Head = headAttr.Value
	w.FIToFIPmtStsRpt.GrpHdr.MsgId = pacs_002_001_10.Max35Text(params.MessageID)
	w.FIToFIPmtStsRpt.GrpHdr.CreDtTm = dt.ISODateTime(params.CreatedOn)

	var originalTransactionReference *OriginalTransactionReference28
	if params.InterbankSettlementAmount != nil && params.InterbankSettlementDate != nil {
		originalTransactionReference = &OriginalTransactionReference28{
			IntrBkSttlmAmt: ActiveOrHistoricCurrencyAndAmount{
				Value: *params.InterbankSettlementAmount,
				Ccy:   ActiveOrHistoricCurrencyCode(params.InterbankSettlementAmountCcy),
			},
			IntrBkSttlmDt: dt.ISODate(*params.InterbankSettlementDate),
		}
	}

	w.FIToFIPmtStsRpt.TxInfAndSts = append(w.FIToFIPmtStsRpt.TxInfAndSts, TxInfAndSts{
		Head:         headAttr.Value,
		OrgnlInstrId: pacs_002_001_10.Max35Text(params.OriginalInstructionId),
		OrgnlTxId:    pacs_002_001_10.Max35Text(params.OriginalTxId),
		TxSts:        pacs_002_001_10.ExternalPaymentTransactionStatus1Code(string(params.TxSts)),
		ClrSysRef:    params.ClrSysRef,
		AccptncDtTm:  dt.ISODateTime(params.AcceptanceDateTime),
		InstgAgt: BranchAndFinancialInstitutionIdentification6{
			FinancialInstitutionIdentification18{
				ClearingSystemMemberIdentification2{
					MmbId: pacs_002_001_10.Max35Text(params.InstructingAgentMemberId),
				},
			},
		},
		InstdAgt: BranchAndFinancialInstitutionIdentification6{
			FinancialInstitutionIdentification18{
				ClearingSystemMemberIdentification2{
					MmbId: pacs_002_001_10.Max35Text(params.InstructedAgentMemberId),
				},
			},
		},
		OrgnlTxRef: originalTransactionReference,
	})

	// Include status reason code if transaction was rejected
	if params.TxSts == TxStatusRejected {
		w.FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf = append(w.FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf, StsRsnInf{
			Rsn: StatusReason6Choice{
				Cd: pacs_002_001_10.ExternalStatusReason1Code(params.StsRsnCd),
			},
		})
	}

	w.FIToFIPmtStsRpt.OrgnlGrpInfAndSts = append(w.FIToFIPmtStsRpt.OrgnlGrpInfAndSts, OrgnlGrpInfAndSts{
		Head:         headAttr.Value,
		OrgnlMsgId:   pacs_002_001_10.Max35Text(params.OriginalMessageId),
		OrgnlMsgNmId: pacs_002_001_10.Max35Text(params.OriginalMessageNameId),
		OrgnlCreDtTm: dt.ISODateTime(params.OriginalCreatedOn),
		OrgnlNbOfTxs: params.OriginalNumberOfTransactions,
	})

	return w
}
