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
	for i := range d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts {
		return string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts[i].OrgnlMsgId)
	}
	return ""
}

func (d Document) OrgnlInstrId() string {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		return string(*d.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlInstrId)
	}
	return ""
}

func (d Document) OrgnlTxId() string {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		return string(*d.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxId)
	}
	return ""
}

func (d Document) TxSts() string {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		return string(*d.FIToFIPmtStsRpt.TxInfAndSts[i].TxSts)
	}
	return ""
}

func (d Document) StsRsn() string {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		for j := range d.FIToFIPmtStsRpt.TxInfAndSts[i].StsRsnInf {
			return string(*d.FIToFIPmtStsRpt.TxInfAndSts[i].StsRsnInf[j].Rsn.Cd)
		}
	}
	return ""
}

func (d Document) IntrBkSttlmAmt() float64 {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		return float64(d.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxRef.IntrBkSttlmAmt.Value)
	}
	return 0.0
}

func (d Document) IntrBkSttlmDt() time.Time {
	for i := range d.FIToFIPmtStsRpt.TxInfAndSts {
		return time.Time(*d.FIToFIPmtStsRpt.TxInfAndSts[i].OrgnlTxRef.IntrBkSttlmDt)
	}
	return time.Time{}
}
