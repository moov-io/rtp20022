package msg_status

import (
	"bytes"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWriterTxStatusReceived(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")

	req := NewWriter(WriteParams{
		MessageID:                    "M20230222200000057A1BFFF87758970177",
		CreatedOn:                    time.Date(2023, 2, 22, 18, 19, 45, 0, eastern),
		OriginalInstructionId:        "20230222200000057A1B123856846970177",
		OriginalTxId:                 "20230222200000057A1B123856846970177",
		InstructingAgentMemberId:     "123456780",
		InstructedAgentMemberId:      "234567891",
		OriginalMessageId:            "M20230222234567891T1BOTS01850058244",
		OriginalMessageNameId:        "camt.035.001.05",
		OriginalCreatedOn:            time.Date(2023, 2, 22, 18, 19, 25, 0, eastern),
		OriginalNumberOfTransactions: 1,
		TxSts:                        TxStatusReceived,
		AcceptanceDateTime:           time.Date(2023, 2, 22, 18, 19, 25, 0, eastern),
	})

	generated, err := xml.MarshalIndent(req, "", "  ")
	require.NoError(t, err)
	generated = bytes.TrimSpace(generated)

	expected, err := os.ReadFile(filepath.Join("testdata", "msg-status-recv.xml"))
	require.NoError(t, err)
	expected = bytes.TrimSpace(expected)

	require.Equal(t, string(expected), string(generated))
}

func TestWriterTxStatusRejected(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")

	req := NewWriter(WriteParams{
		MessageID:                    "M20230222200000057A1BFFF87758970177",
		CreatedOn:                    time.Date(2023, 2, 22, 18, 19, 45, 0, eastern),
		OriginalInstructionId:        "20230222200000057A1B123856846970177",
		OriginalTxId:                 "20230222200000057A1B123856846970177",
		InstructingAgentMemberId:     "123456780",
		InstructedAgentMemberId:      "234567891",
		OriginalMessageId:            "M20230222234567891T1BOTS01850058244",
		OriginalMessageNameId:        "pacs.008.001.08",
		OriginalCreatedOn:            time.Date(2023, 2, 22, 18, 19, 25, 0, eastern),
		OriginalNumberOfTransactions: 1,
		TxSts:                        TxStatusRejected,
		StsRsnCd:                     "AC03",
		AcceptanceDateTime:           time.Date(2023, 2, 22, 18, 19, 25, 0, eastern),
	})

	generated, err := xml.MarshalIndent(req, "", "  ")
	require.NoError(t, err)
	generated = bytes.TrimSpace(generated)

	expected, err := os.ReadFile(filepath.Join("testdata", "msg-status-rjct.xml"))
	require.NoError(t, err)
	expected = bytes.TrimSpace(expected)

	require.Equal(t, string(expected), string(generated))
}
