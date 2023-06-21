package head_001_001_01

import (
	"github.com/moov-io/base/log"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func (r BusinessApplicationHeaderV01) Context() map[string]log.Valuer {
	var creDt, _ = (rtp.ISONormalisedDateTime)(r.CreDt).MarshalText()
	return log.Fields{
		"from_institution_id":   log.String(r.FromMemberID()),
		"to_institution_id":     log.String(r.ToMemberID()),
		"business_message_id":   log.String(string(r.BizMsgIdr)),
		"message_definition_id": log.String(string(r.MsgDefIdr)),
		"creation_date":         log.String(string(creDt)),
		"duplicate":             log.String(string(*r.CpyDplct)),
	}
}

func (r BusinessApplicationHeaderV01) FromMemberID() string {
	return string(r.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
}

func (r BusinessApplicationHeaderV01) ToMemberID() string {
	return string(r.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
}
