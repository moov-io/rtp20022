package response_return_of_funds

import (
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	now := time.Date(2023, 5, 16, 17, 20, 27, 0, eastern)

	req := NewWriter(WriteParams{
		CreatedOn:             now,
		CaseAssignmentId:      "M20230516200000057A1B12884072027248",
		AssignorAgentMemberId: "234567891",
		AssigneeAgentMemberId: "123456789",
		StsConfCd:             "IPAY",
		OriginalMessageId:     "M20230516234567891T1BOTS01331895026",
		OriginalMessageNameId: "camt.056.001.08",
		OriginalCreatedOn:     now,
	})

	generated, err := xml.MarshalIndent(req, "", "  ")
	require.NoError(t, err)
	generated = bytes.TrimSpace(generated)

	expected, err := os.ReadFile(filepath.Join("testdata", "response-return-of-funds-ipay.xml"))
	require.NoError(t, err)
	expected = bytes.TrimSpace(expected)

	require.Equal(t, string(expected), string(generated))
}
