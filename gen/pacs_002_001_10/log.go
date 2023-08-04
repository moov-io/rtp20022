package pacs_002_001_10

import "github.com/moov-io/base/log"

func (d DocumentTCH) Context() map[string]log.Valuer {
	return log.Fields{
		"original_message_identification":     log.String(string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)),
		"original_message_name":               log.String(string(d.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId)),
		"original_instruction_identification": log.String(d.OrgnlInstrId()),
		"original_transaction_identification": log.String(d.OrgnlTxId()),
	}
}
