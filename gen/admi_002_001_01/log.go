package admi_002_001_01

import "github.com/moov-io/base/log"

func (d DocumentTCH) Context() map[string]log.Valuer {
	return log.Fields{
		"original_identification": log.String(d.Admi00200101.Reference()),
		"reject_reason":           log.String(d.Admi00200101.Reason()),
		"additional_data":         log.String(d.Admi00200101.AdditionalData()),
	}
}
