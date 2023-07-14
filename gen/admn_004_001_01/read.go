package admn_004_001_01

import "github.com/moov-io/base/log"

func (d DocumentTCH) Context() map[string]log.Valuer {
	var rsnPrtry = ""
	if d.AdmnSignOffResp.SignOffResp.StsRsnInf != nil {
		rsnPrtry = string(*d.AdmnSignOffResp.SignOffResp.StsRsnInf.Rsn.Prtry)
	}
	return log.Fields{
		"status":        log.String(string(d.AdmnSignOffResp.SignOffResp.Sts)),
		"status_reason": log.String(rsnPrtry),
	}
}
