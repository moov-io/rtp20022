package rtp_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/moov-io/rtp20022/pkg/rtp"

	"github.com/stretchr/testify/require"
)

func TestCaseAssignment(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	when := time.Date(2015, time.November, 15, 0, 30, 0, 0, eastern)
	participantID := "11021200201"

	caseAssignmentID := rtp.CaseAssignmentID(when, participantID, "FFF", nil)
	require.Contains(t, caseAssignmentID, "M2015111511021200201BFFF")

	rr := regexp.MustCompile(`^M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*$`)
	require.True(t, rr.MatchString(caseAssignmentID))
}
