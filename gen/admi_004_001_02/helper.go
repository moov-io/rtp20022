package admi_004_001_02

import (
	"time"
)

func (s SystemEventNotificationV02TCH) Code() string {
	return string(s.EvtInf.EvtCd)
}

func (s SystemEventNotificationV02TCH) Parameters() []string {
	var out []string
	for i := range s.EvtInf.EvtParam {
		out = append(out, string(s.EvtInf.EvtParam[i]))
	}
	return out
}

func (s SystemEventNotificationV02TCH) Description() string {
	if s.EvtInf.EvtDesc != nil {
		return string(*s.EvtInf.EvtDesc)
	}
	return ""
}

func (s SystemEventNotificationV02TCH) Time() time.Time {
	return time.Time(s.EvtInf.EvtTm)
}
