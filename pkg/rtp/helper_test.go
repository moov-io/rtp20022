package rtp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func OnesSerial(length int) string {
	return strings.Repeat("1", length)
}

func TestPickGenerator(t *testing.T) {
	serial := pickGenerator(nil, OnesSerial)(11)
	require.Equal(t, "11111111111", serial)

	serial = pickGenerator(RandomSerial, OnesSerial)(11)
	require.NotEqual(t, "11111111111", serial)
}
