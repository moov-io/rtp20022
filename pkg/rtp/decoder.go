package rtp

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/rtp20022"
)

// maxTokenLimit is the upper limit of how many XML tokens to consume.
// The XML documents we receive should only have a couple tokens to iterate
// through. (e.g. AppHdr and SignOnResponse)
const maxTokenLimit = 10

var (
	ErrUnmatchedDocument = errors.New("exceeded max number of decode attempts")
)

func decodeXML[H any](data []byte, head H, factories map[string]rtp20022.Document) (H, rtp20022.Document, error) {
	var body rtp20022.Document

	var start *xml.StartElement
	var unhandledTokens []string

	tokensConsumed := 0
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		tokensConsumed += 1
		if tokensConsumed >= maxTokenLimit {
			tagNames := strings.Join(unhandledTokens, ",")
			return head, body, fmt.Errorf("%w: unhandled tokens: %v", ErrUnmatchedDocument, tagNames)
		}
		if token == nil {
			continue
		}

		switch elm := token.(type) {
		case xml.StartElement:
			switch elm.Name.Local {
			case "AppHdr":
				err = decoder.DecodeElement(&head, &elm)
				if err != nil {
					return head, nil, fmt.Errorf("error decoding AppHdr: %v", err)
				}

			default:
				factory, exists := factories[elm.Name.Local]
				if exists {
					body = factory
					start = &elm
					goto decode
				} else {
					unhandledTokens = append(unhandledTokens, elm.Name.Local)
				}
			}
		}
	}

decode:
	err := decoder.DecodeElement(&body, start)
	if err != nil {
		return head, nil, fmt.Errorf("error decoding %T as body: %v", body, err)
	}
	return head, body, nil
}
