package header

import (
	"encoding/xml"

	"github.com/moov-io/base/log"
)

func NewReader() Reader {
	return Reader{}
}

type Reader struct {
	XMLName xml.Name `xml:"AppHdr"`
	Text    string   `xml:",chardata"`
	Head    string   `xml:"head,attr"`
	Fr      struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
		FIId struct {
			Text       string `xml:",chardata"`
			FinInstnId struct {
				Text        string `xml:",chardata"`
				ClrSysMmbId struct {
					Text  string `xml:",chardata"`
					MmbId string `xml:"MmbId"`
				} `xml:"ClrSysMmbId"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"Fr"`
	To struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
		FIId struct {
			Text       string `xml:",chardata"`
			FinInstnId struct {
				Text        string `xml:",chardata"`
				ClrSysMmbId struct {
					Text  string `xml:",chardata"`
					MmbId string `xml:"MmbId"`
				} `xml:"ClrSysMmbId"`
			} `xml:"FinInstnId"`
		} `xml:"FIId"`
	} `xml:"To"`
	BizMsgIdr struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
	} `xml:"BizMsgIdr"`
	MsgDefIdr struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
	} `xml:"MsgDefIdr"`
	CreDt struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
	} `xml:"CreDt"`
	CpyDplct struct {
		Text string `xml:",chardata"`
		Head string `xml:"head,attr"`
	} `xml:"CpyDplct"`
}

func (r Reader) Context() map[string]log.Valuer {
	return log.Fields{
		"from_institution_id":   log.String(r.FromMemberID()),
		"to_institution_id":     log.String(r.ToMemberID()),
		"business_message_id":   log.String(r.BizMsgIdr.Text),
		"message_definition_id": log.String(r.MsgDefIdr.Text),
		"creation_date":         log.String(r.CreDt.Text),
		"duplicate":             log.String(r.CpyDplct.Text),
	}
}

func (r Reader) FromMemberID() string {
	return r.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId
}

func (r Reader) ToMemberID() string {
	return r.To.FIId.FinInstnId.ClrSysMmbId.MmbId
}
