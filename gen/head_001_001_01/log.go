package head_001_001_01

import "github.com/moov-io/base/log"

func (r BusinessApplicationHeaderV01TCH) Context() map[string]log.Valuer {
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
