package rtp_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestBusinessID(t *testing.T) {
	businessID := rtp.BusinessID(time.Now(), "44445566667", "128")
	if testing.Verbose() {
		t.Log(businessID)
	}
	require.Equal(t, 35, len(businessID))
	require.Contains(t, businessID, "44445566667B128")

	rr := regexp.MustCompile(`^[B]((19|20)\d{2})(0+[1-9]|1[012])(0+[1-9]|[12]\d|3[01])([a-zA-Z0-9]{11})[BH][a-zA-Z0-9]{3}[0-9]{11}$`)
	require.True(t, rr.MatchString(businessID))
}
