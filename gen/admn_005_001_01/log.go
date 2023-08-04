package admn_005_001_01

import "github.com/moov-io/base/log"

func (d DocumentTCH) Context() map[string]log.Valuer {
	return log.Fields{
		"message_id":        log.String(string(d.AdmnEchoReq.GrpHdr.MsgId)),
		"instructing_agent": log.String(string(d.AdmnEchoReq.EchoTxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId)),
		"instructed_agent":  log.String(string(d.AdmnEchoReq.EchoTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId)),
	}
}
