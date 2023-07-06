package head_001_001_01

import (
	"github.com/moov-io/base/log"
)

func (r BusinessApplicationHeaderV01) Context() map[string]log.Valuer {
	var creDt, _ = r.CreDt.MarshalText()
	return log.Fields{
		"from_institution_id":   log.String(r.FromMemberID()),
		"to_institution_id":     log.String(r.ToMemberID()),
		"business_message_id":   log.String(string(r.BizMsgIdr)),
		"message_definition_id": log.String(string(r.MsgDefIdr)),
		"creation_date":         log.String(string(creDt)),
		"duplicate":             log.String(r.CopyDuplicateCode()),
	}
}

func (r BusinessApplicationHeaderV01) CopyDuplicateCode() string {
	if r.CpyDplct != nil {
		return string(*r.CpyDplct)
	}
	return ""
}

func (r BusinessApplicationHeaderV01) FromMemberID() string {
	return string(r.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
}

func (r BusinessApplicationHeaderV01) ToMemberID() string {
	return string(r.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
}
