package rtp_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestIDZeroPadding(t *testing.T) {
	when := time.Date(2026, time.November, 24, 10, 0, 0, 0, rtp.Eastern())

	shortSerial := rtp.SerialGenerator(func(int) string { return "42" })
	longSerial := rtp.SerialGenerator(func(int) string { return "123456789012345" })

	cases := map[string]struct {
		got      string
		expected string
	}{
		"MessageID short inputs": {
			got:      rtp.MessageID(when, "12345", "AB", shortSerial),
			expected: "M2026112400000012345B0AB00000000042",
		},
		"MessageID long serial": {
			got:      rtp.MessageID(when, "11021200201", "FFF", longSerial),
			expected: "M2026112411021200201BFFF12345678901",
		},
		"TransactionID short inputs": {
			got:      rtp.TransactionID(when, "12345", "AB", shortSerial),
			expected: "2026112400000012345B00AB00000000042",
		},
		"TransactionID long serial": {
			got:      rtp.TransactionID(when, "11021200201", "FFFF", longSerial),
			expected: "2026112411021200201BFFFF12345678901",
		},
		"BusinessID short inputs": {
			got:      rtp.BusinessID(when, "12345", "AB", shortSerial),
			expected: "B2026112400000012345B0AB00000000042",
		},
		"InstructionID short inputs": {
			got:      rtp.InstructionID(when, "12345", "AB", shortSerial),
			expected: "2026112400000012345B00AB00000000042",
		},
		"CaseAssignmentID short inputs": {
			got:      rtp.CaseAssignmentID(when, "12345", "AB", shortSerial),
			expected: "M2026112400000012345B0AB00000000042",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.got)
			require.Len(t, tc.got, 35)
		})
	}
}
