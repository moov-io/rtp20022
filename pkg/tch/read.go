package tch

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022"
	"github.com/moov-io/rtp20022/pkg/header"
)

type Reader struct {
	XMLName xml.Name   `xml:"Message"`
	Attrs   []xml.Attr `xml:",attr"`

	Header header.Reader `xml:"AppHdr"`
	Body   rtp20022.Document
}

func NewReader(head header.Reader, body rtp20022.Document) *Reader {
	return &Reader{
		Attrs: []xml.Attr{
			{
				Name:  xml.Name{Local: "xmlns"},
				Value: "urn:tch",
			},
		},
		Header: head,
		Body:   body,
	}
}

func (m *Reader) Validate() error {
	return m.Body.Validate()
}
