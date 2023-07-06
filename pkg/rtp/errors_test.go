package rtp_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestErrorCodes(t *testing.T) {
	ec := rtp.IsError("9910")
	require.NotNil(t, ec)
	require.Equal(t, "9910", ec.Code)
	require.Equal(t, rtp.ErrTemporary, ec.Level)

	ec = rtp.IsError("be07")
	require.NotNil(t, ec)
	require.Equal(t, "BE07", ec.Code)
	require.Equal(t, rtp.ErrFatal, ec.Level)

	ec = rtp.IsError("ac03")
	require.NotNil(t, ec)
	require.Equal(t, "AC03", ec.Code)
	require.Equal(t, rtp.ErrAccountFatal, ec.Level)

	ec = rtp.IsError("AM13")
	require.NotNil(t, ec)
	require.Equal(t, "AM13", ec.Code)
	require.Equal(t, rtp.ErrTemporary, ec.Level)

	ec = rtp.IsError("AM09")
	require.NotNil(t, ec)
	require.Equal(t, "AM09", ec.Code)
	require.Equal(t, rtp.ErrLogic, ec.Level)

	ec = rtp.IsError("")
	require.Nil(t, ec)

	ec = rtp.IsError("9999")
	require.Nil(t, ec)
}

func TestErrorLevel(t *testing.T) {
	// just make sure .Error() doesn't recusively panic
	level := rtp.ErrNetwork
	require.Equal(t, "Network issue", level.Error())
	require.Equal(t, "Network issue", fmt.Sprintf("%v", level))
	require.Equal(t, "Network issue", fmt.Errorf("%w", level).Error())
}
