package rtp_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestInstructionID(t *testing.T) {
	instID := rtp.InstructionID(time.Now(), "44445566667", "1238")
	if testing.Verbose() {
		t.Log(instID)
	}
	require.Equal(t, 35, len(instID))
	require.Contains(t, instID, "44445566667B1238")

	// Short bankField
	instID = rtp.InstructionID(time.Now(), "44445566667", "238")
	if testing.Verbose() {
		t.Log(instID)
	}
	require.Equal(t, 35, len(instID))
	require.Contains(t, instID, "44445566667B0238")
}
