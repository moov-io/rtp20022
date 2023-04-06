package msg_status

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
)

func NewReader() Reader {
	return Reader{}
}

type Reader struct {
	XMLName xml.Name `xml:"MessageStatusReport"`
	Text    string   `xml:",chardata"`

	FIToFIPmtStsRpt pacs_002_001_10.FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt"`
}

func (p *Reader) Validate() error {
	return p.FIToFIPmtStsRpt.Validate()
}

func (p Reader) MsgId() string {
	return string(p.FIToFIPmtStsRpt.GrpHdr.MsgId)
}

func (p Reader) OrgnlMsgId() string {
	for i := range p.FIToFIPmtStsRpt.OrgnlGrpInfAndSts {
		return string(p.FIToFIPmtStsRpt.OrgnlGrpInfAndSts[i].OrgnlMsgId)
	}
	return ""
}

func (p Reader) OrgnlInstrId() string {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		return string(p.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlInstrId)
	}
	return ""
}

func (p Reader) OrgnlTxId() string {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		return string(p.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxId)
	}
	return ""
}

func (p Reader) TxSts() string {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		return string(p.FIToFIPmtStsRpt.TxInfAndSts[i].TxSts)
	}
	return ""
}

func (p Reader) StsRsn() string {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		for j := range p.FIToFIPmtStsRpt.TxInfAndSts[i].StsRsnInf {
			return string(p.FIToFIPmtStsRpt.TxInfAndSts[i].StsRsnInf[j].Rsn.Cd)
		}
	}
	return ""
}

func (p Reader) IntrBkSttlmAmt() float64 {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		return p.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxRef.IntrBkSttlmAmt.Value
	}
	return 0.0
}

func (p Reader) IntrBkSttlmDt() time.Time {
	for i := range p.FIToFIPmtStsRpt.TxInfAndSts {
		return time.Time(p.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxRef.IntrBkSttlmDt)
	}
	return time.Time{}
}
