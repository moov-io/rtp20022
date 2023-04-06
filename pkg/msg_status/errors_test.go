package msg_status

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorCodes(t *testing.T) {
	ec := IsError("9910")
	require.NotNil(t, ec)
	require.Equal(t, "9910", ec.Code)
	require.Equal(t, ErrTemporary, ec.Level)

	ec = IsError("be07")
	require.NotNil(t, ec)
	require.Equal(t, "BE07", ec.Code)
	require.Equal(t, ErrFatal, ec.Level)

	ec = IsError("ac03")
	require.NotNil(t, ec)
	require.Equal(t, "AC03", ec.Code)
	require.Equal(t, ErrAccountFatal, ec.Level)

	ec = IsError("AM13")
	require.NotNil(t, ec)
	require.Equal(t, "AM13", ec.Code)
	require.Equal(t, ErrTemporary, ec.Level)

	ec = IsError("AM09")
	require.NotNil(t, ec)
	require.Equal(t, "AM09", ec.Code)
	require.Equal(t, ErrLogic, ec.Level)

	ec = IsError("")
	require.Nil(t, ec)

	ec = IsError("9999")
	require.Nil(t, ec)
}

func TestErrorLevel(t *testing.T) {
	// just make sure .Error() doesn't recusively panic
	level := ErrNetwork
	require.Equal(t, "Network issue", level.Error())
	require.Equal(t, "Network issue", fmt.Sprintf("%v", level))
	require.Equal(t, "Network issue", fmt.Errorf("%w", level).Error())
}
