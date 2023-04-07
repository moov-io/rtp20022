package header

import (
	"bytes"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWriter(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	req := NewWriter(WriteParams{
		FromMemberID:      "990000001T1",
		ToMemberID:        "200000057A1",
		BusinessID:        "B20221221990000001T1HOTS01101645161",
		MessageDefinition: "admn.002.001.01",
		CreatedOn:         time.Date(2022, 12, 21, 15, 31, 19, 0, eastern),
	})

	generated, err := xml.MarshalIndent(req, "", "  ")
	require.NoError(t, err)
	generated = bytes.TrimSpace(generated)

	expected, err := os.ReadFile(filepath.Join("testdata", "header-no-signature.xml"))
	require.NoError(t, err)
	expected = bytes.TrimSpace(expected)

	require.Equal(t, string(expected), string(generated))
}

func TestWriterDUPL(t *testing.T) {
	eastern, _ := time.LoadLocation("America/New_York")
	req := NewWriter(WriteParams{
		FromMemberID:      "990000001T1",
		ToMemberID:        "200000057A1",
		BusinessID:        "B20221221990000001T1HOTS01101645161",
		MessageDefinition: "admn.002.001.01",
		CreatedOn:         time.Date(2022, 12, 21, 15, 31, 19, 0, eastern),
		Duplicate:         true,
	})

	generated, err := xml.MarshalIndent(req, "", "  ")
	require.NoError(t, err)
	generated = bytes.TrimSpace(generated)

	expected, err := os.ReadFile(filepath.Join("testdata", "header-dupl-no-signature.xml"))
	require.NoError(t, err)
	expected = bytes.TrimSpace(expected)

	require.Equal(t, string(expected), string(generated))
}
