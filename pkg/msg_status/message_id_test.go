package msg_status

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMessageID(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	when := time.Date(2017, time.November, 12, 0, 0, 0, 0, eastern)
	participantID := "02001000101"
	messageID := MessageID(when, participantID, "MSR")

	require.Contains(t, messageID, "M2017111202001000101BMSR")
	require.Len(t, messageID, 35)

	rr := regexp.MustCompile(`M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*`)
	require.True(t, rr.MatchString(messageID))
}
