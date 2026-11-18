package rtp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestRemapReasonFromAdditionalInfo(t *testing.T) {
	bankcore2005 := "2005:Bankcore service internal server error, 5001:Credit inquiry failed:Fatal error occurred. Please report"

	tests := []struct {
		name   string
		reason string
		addtl  string
		want   string
	}{
		{
			name:   "1100 with 2005 becomes 9912",
			reason: "1100",
			addtl:  bankcore2005,
			want:   "9912",
		},
		{
			name:   "1100 with invalid account becomes AC03",
			reason: "1100",
			addtl:  "20008:Invalid Account Number",
			want:   "AC03",
		},
		{
			name:   "1100 with invalid-account needle becomes AC03",
			reason: "1100",
			addtl:  "2003:Rejected by bankcore service, 20008:Invalid Account Number",
			want:   "AC03",
		},
		{
			name:   "invalid account needle",
			reason: "1100",
			addtl:  "Invalid Account Number",
			want:   "AC03",
		},
		{
			name:   "internal server error needle becomes 9912",
			reason: "1100",
			addtl:  "Bankcore service internal server error",
			want:   "9912",
		},
		{
			name:   "account-fatal wins over retryable",
			reason: "1100",
			addtl:  "2005:Bankcore service internal server error, 20008:Invalid Account Number",
			want:   "AC03",
		},
		{
			name:   "unmapped 1100 stays 1100",
			reason: "1100",
			addtl:  "Any Other Reasons.",
			want:   "1100",
		},
		{
			name:   "5001 without 2005 stays 1100",
			reason: "1100",
			addtl:  "5001:Credit inquiry failed:Fatal error occurred",
			want:   "1100",
		},
		{
			name:   "empty AddtlInf stays 1100",
			reason: "1100",
			addtl:  "",
			want:   "1100",
		},
		{
			name:   "9910 is not remapped from AddtlInf",
			reason: "9910",
			addtl:  bankcore2005,
			want:   "9910",
		},
		{
			name:   "AC03 is not remapped from AddtlInf",
			reason: "AC03",
			addtl:  bankcore2005,
			want:   "AC03",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rtp.RemapReasonFromAdditionalInfo(tt.reason, tt.addtl)
			require.Equal(t, tt.want, got)
		})
	}
}
