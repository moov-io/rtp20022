package pacs_002_001_10

import "time"

func (d DocumentTCH) Validate() error {
	return d.FIToFIPmtStsRpt.Validate()
}

func (p *FIToFIPaymentStatusReportV10TCH) Validate() error {
	return nil // TODO(adam):
}

func (d DocumentTCH) MsgId() string {
	return string(d.FIToFIPmtStsRpt.GrpHdr.MsgId)
}

func (d DocumentTCH) OrgnlMsgId() string {
	return string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)
}

func (d DocumentTCH) OrgnlInstrId() string {
	return string(d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlInstrId)
}

func (d DocumentTCH) OrgnlTxId() string {
	return string(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId)
}

func (d DocumentTCH) TxSts() string {
	return string(d.FIToFIPmtStsRpt.TxInfAndSts.TxSts)
}

func (d DocumentTCH) StsRsn() string {
	return string(*d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn.Cd)
}

func (d DocumentTCH) IntrBkSttlmAmt() float64 {
	return float64(d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Value)
}

func (d DocumentTCH) IntrBkSttlmDt() time.Time {
	return time.Time(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt)
}
