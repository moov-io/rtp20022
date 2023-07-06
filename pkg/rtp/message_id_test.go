package rtp_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestMessageID(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	when := time.Date(2015, time.November, 15, 0, 30, 0, 0, eastern)
	participantID := "11021200201"
	messageID := rtp.MessageID(when, participantID, "FFF")

	require.Contains(t, messageID, "M2015111511021200201BFFF")
	require.Len(t, messageID, 35)

	rr := regexp.MustCompile(`^[M]((19|20)\d{2})(0+[1-9]|1[012])(0+[1-9]|[12]\d|3[01])([a-zA-Z0-9]{11})[B][a-zA-Z0-9]{3}[0-9]{11}$`)
	require.True(t, rr.MatchString(messageID))
}

func TestAdmnMessageID(t *testing.T) {
	messageID := rtp.AdmnMessageID(time.Now(), "44445566667")
	if testing.Verbose() {
		t.Log(messageID)
	}
	require.Equal(t, 35, len(messageID))
	require.Contains(t, messageID, "44445566667")
}
