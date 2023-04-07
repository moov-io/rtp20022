package tch

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022"
	"github.com/moov-io/rtp20022/pkg/header"
)

type Writer struct {
	XMLName xml.Name   `xml:"Message"`
	Attrs   []xml.Attr `xml:",attr"`

	Header header.Writer `xml:"AppHdr"`
	Body   rtp20022.Document
}

func NewWriter(head header.Writer, body rtp20022.Document) *Writer {
	return &Writer{
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

func (m *Writer) Validate() error {
	return m.Body.Validate()
}
