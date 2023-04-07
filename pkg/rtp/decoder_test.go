package rtp

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/rtp20022"
	"github.com/moov-io/rtp20022/pkg/header"
	"github.com/moov-io/rtp20022/pkg/msg_status"
	"github.com/moov-io/rtp20022/pkg/tch"

	"github.com/stretchr/testify/require"
)

func parse(data []byte) (*tch.Reader, error) {
	head, body, err := DecodeXML(data, header.Reader{}, map[string]rtp20022.Document{
		"MessageStatusReport": &msg_status.Reader{},
	})
	r := tch.NewReader(head, body)
	if err != nil {
		return r, fmt.Errorf("%w: parsing TCH message", err)
	}
	return r, nil
}

func TestDecode(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "..", "testdata", "credit_transfer_reject_response.xml"))
	require.NoError(t, err)

	msg, err := parse(bs)
	require.NoError(t, err)

	require.Equal(t, "990000001T1", msg.Header.FromMemberID())
	require.Equal(t, "200000057A1", msg.Header.ToMemberID())

	msgStatus, ok := msg.Body.(*msg_status.Reader)
	require.True(t, ok)
	require.Equal(t, "M20230112BOTS00697603054", msgStatus.MsgId())
}
