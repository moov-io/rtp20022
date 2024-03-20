package rtp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaddedSerial(t *testing.T) {
	output := PaddedSerial("abcd")(8)
	require.Equal(t, "0000abcd", output)
}
