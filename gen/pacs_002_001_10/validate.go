package pacs_002_001_10

import "time"

func (d Document) Validate() error {
	return d.FIToFIPmtStsRpt.Validate()
}

func (p *FIToFIPaymentStatusReportV10) Validate() error {
	return nil // TODO(adam):
}

func (d Document) MsgId() string {
	return string(d.FIToFIPmtStsRpt.GrpHdr.MsgId)
}

func (d Document) OrgnlMsgId() string {
	return string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)
}

func (d Document) OrgnlInstrId() string {
	return string(d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlInstrId)
}

func (d Document) OrgnlTxId() string {
	return string(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId)
}

func (d Document) TxSts() string {
	return string(d.FIToFIPmtStsRpt.TxInfAndSts.TxSts)
}

func (d Document) StsRsn() string {
	return string(*d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn.Cd)
}

func (d Document) IntrBkSttlmAmt() float64 {
	return float64(d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Value)
}

func (d Document) IntrBkSttlmDt() time.Time {
	return time.Time(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt)
}
