package pacs_002_001_10

import (
	"cloud.google.com/go/civil"
)

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
	if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId != nil {
		return string(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId)
	}
	return ""
}

func (d DocumentTCH) TxSts() string {
	return string(d.FIToFIPmtStsRpt.TxInfAndSts.TxSts)
}

func (d DocumentTCH) StsRsn() string {
	if d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf != nil {
		if d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn != nil {
			if d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn.Cd != nil {
				return string(*d.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn.Cd)
			}
		}
	}
	return ""
}

func (d DocumentTCH) IntrBkSttlmAmt() float64 {
	if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef != nil {
		if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt != nil {
			return float64(d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmAmt.Value)
		}
	}
	return 0.0
}

func (d DocumentTCH) IntrBkSttlmDt() civil.Date {
	if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef != nil {
		if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt != nil {
			return civil.Date(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt)
		}
	}
	return civil.Date{}
}
