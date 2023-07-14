package pacs_002_001_10

import (
	"time"

	"github.com/moov-io/base/log"
)

func (d DocumentTCH) Context() map[string]log.Valuer {
	return log.Fields{
		"original_message_identification":     log.String(string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)),
		"original_message_name":               log.String(string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId)),
		"original_instruction_identification": log.String(d.OrgnlInstrId()),
		"original_transaction_identification": log.String(d.OrgnlTxId()),
	}
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

func (d DocumentTCH) IntrBkSttlmDt() time.Time {
	if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef != nil {
		if d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt != nil {
			return time.Time(*d.FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxRef.IntrBkSttlmDt)
		}
	}
	return time.Time{}
}
