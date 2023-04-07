package tch

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInstructionID(t *testing.T) {
	instID := InstructionID(time.Now(), "44445566667", "1238")
	if testing.Verbose() {
		t.Log(instID)
	}
	require.Equal(t, 35, len(instID))
	require.Contains(t, instID, "44445566667B1238")
}
