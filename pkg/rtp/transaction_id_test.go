package rtp_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestTransactionID(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	when := time.Date(2015, time.November, 15, 0, 30, 0, 0, eastern)
	participantID := "11021200201"
	transactionID := rtp.TransactionID(when, participantID, "FFFF")

	require.Contains(t, transactionID, "2015111511021200201BFFF")
	require.Len(t, transactionID, 35)

	rr := regexp.MustCompile(`^((19|20)\d{2})(0+[1-9]|1[012])(0+[1-9]|[12]\d|3[01])([a-zA-Z0-9]{11})[B]([a-zA-Z0-9]{15})$`)
	require.True(t, rr.MatchString(transactionID))
}
