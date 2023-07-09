package test

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/gen/admn_004_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn004Constant = &admn_004_001_01.DocumentTCH{
	AdmnSignOffResp: admn_004_001_01.SignOffResponseTCH{
		GrpHdr: admn_004_001_01.GrpHdrTCH{
			MsgId:   "20230113125830990000001T11296198522",
			CreDtTm: rtp.UnmarshalISODateTime("2023-01-13T12:58:30"),
		},
		SignOffResp: admn_004_001_01.SignOffRespTCH{
			OrgnlInstrId: "20230113200000057A1B123862501723637",
			InstgAgt: admn_004_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_004_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_004_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_004_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_004_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_004_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			Sts: admn_004_001_01.TransactionGroupStatus3CodeAdminActc,
		},
	},
}

func TestReadAdmn004(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn004.RTP.xml"))
	require.NoError(t, err)

	admn004 := &messages.Message{}
	err = xml.Unmarshal(input, admn004)
	require.NoError(t, err)

	expected := messages.NewAdmn004Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOffResponse = admn004Constant
	expected.SignOffResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOffResponse",
	}

	assert.Equal(t, expected, admn004)
}

func TestWriteAdmn004(t *testing.T) {
	input := messages.NewAdmn004Message()
	input.SignOffResponse = admn004Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn004.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
