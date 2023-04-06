package credit_transfer_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/moov-io/rtp20022/pkg/credit_transfer"

	"github.com/stretchr/testify/require"
)

func TestMessageID(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	when := time.Date(2015, time.November, 15, 0, 30, 0, 0, eastern)
	participantID := "11021200201"
	messageID := credit_transfer.MessageID(when, participantID, "FFF")

	require.Contains(t, messageID, "M2015111511021200201BFFF")
	require.Len(t, messageID, 35)

	rr := regexp.MustCompile(`^[M]((19|20)\d{2})(0+[1-9]|1[012])(0+[1-9]|[12]\d|3[01])([a-zA-Z0-9]{11})[B][a-zA-Z0-9]{3}[0-9]{11}$`)
	require.True(t, rr.MatchString(messageID))
}
