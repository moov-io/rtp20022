package message_status_report

import (
	"encoding/xml"
	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/pkg/msg_status"
	"strconv"
	"time"
)

type Params struct {
	MsgId                       string
	CreDtTm                     time.Time
	OrgnlMsgId                  string
	OrgnlMsgNmId                string
	OrgnlCreDtTm                time.Time
	OrgnlNbOfTxs                int
	OrgnlInstrId                string
	OrgnlTxId                   string
	TxSts                       msg_status.TransactionStatus
	StsRsnInfRsnCd              string
	AccptncDtTm                 time.Time
	ClrSysRef                   string
	InstgAgt                    string
	InstdAgt                    string
	OrgnlTxRefIntrBkSttlmAmt    int64 // in cents
	OrgnlTxRefIntrBkSttlmAmtCcy string
	OrgnlTxRefIntrBkSttlmDt     time.Time // YYYY-MM-DD
}

type Pacs002 struct {
	XMLName xml.Name `xml:"MessageStatusReport"`
	Text    string   `xml:",chardata"`

	FIToFIPmtStsRpt pacs_002_001_10.FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt"`
}

func (p *Pacs002) Validate() error {
	return p.FIToFIPmtStsRpt.Validate()
}

const Namespace = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"

func New(params Params) *Pacs002 {
	namespaceAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: Namespace,
	}
	responseFIToFIPaymentStatusReport := pacs_002_001_10.FIToFIPaymentStatusReportV10{
		Attr: []xml.Attr{namespaceAttr},
		GrpHdr: pacs_002_001_10.GroupHeader91{
			MsgId:   pacs_002_001_10.Max35Text(params.MsgId),
			CreDtTm: pacs_002_001_10.ISODateTime(params.CreDtTm),
		},
		OrgnlGrpInfAndSts: []pacs_002_001_10.OriginalGroupHeader17{
			{
				OrgnlMsgId:   pacs_002_001_10.Max35Text(params.OrgnlMsgId),
				OrgnlMsgNmId: pacs_002_001_10.Max35Text(params.OrgnlMsgNmId),
				OrgnlCreDtTm: pacs_002_001_10.ISODateTime(params.OrgnlCreDtTm),
				OrgnlNbOfTxs: pacs_002_001_10.Max15NumericText(strconv.Itoa(params.OrgnlNbOfTxs)),
			},
		},
		TxInfAndSts: []pacs_002_001_10.PaymentTransaction110{
			{
				OrgnlInstrId: pacs_002_001_10.Max35Text(params.OrgnlInstrId),
				OrgnlTxId:    pacs_002_001_10.Max35Text(params.OrgnlTxId),
				TxSts:        pacs_002_001_10.ExternalPaymentTransactionStatus1Code(params.TxSts),
				AccptncDtTm:  pacs_002_001_10.ISODateTime(params.AccptncDtTm),
				ClrSysRef:    pacs_002_001_10.Max35Text(params.ClrSysRef),
				InstgAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18{
						ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2{
							MmbId: pacs_002_001_10.Max35Text(params.InstgAgt),
						},
					},
				},
				InstdAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18{
						ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2{
							MmbId: pacs_002_001_10.Max35Text(params.InstdAgt),
						},
					},
				},
				OrgnlTxRef: pacs_002_001_10.OriginalTransactionReference28{
					IntrBkSttlmAmt: pacs_002_001_10.ActiveOrHistoricCurrencyAndAmount{
						Value: float64(params.OrgnlTxRefIntrBkSttlmAmt),
						Ccy:   pacs_002_001_10.ActiveOrHistoricCurrencyCode(params.OrgnlTxRefIntrBkSttlmAmtCcy),
					},
					IntrBkSttlmDt: pacs_002_001_10.ISODate(params.OrgnlTxRefIntrBkSttlmDt),
				},
			},
		},
	}

	return &Pacs002{
		FIToFIPmtStsRpt: responseFIToFIPaymentStatusReport,
	}
}
