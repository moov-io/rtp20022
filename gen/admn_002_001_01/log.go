package admn_002_001_01

import "github.com/moov-io/base/log"

func (d DocumentTCH) Context() map[string]log.Valuer {
	var rsnPrtry = ""
	if d.AdmnSignOnResp.SignOnResp.StsRsnInf != nil {
		rsnPrtry = string(*d.AdmnSignOnResp.SignOnResp.StsRsnInf.Rsn.Prtry)
	}
	return log.Fields{
		"status":        log.String(string(d.AdmnSignOnResp.SignOnResp.Sts)),
		"status_reason": log.String(rsnPrtry),
	}
}
