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

	"github.com/moov-io/rtp20022/gen/admn_003_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn003Constant = &admn_003_001_01.DocumentTCH{
	AdmnSignOffReq: admn_003_001_01.SignOffRequestTCH{
		GrpHdr: admn_003_001_01.GrpHdrTCH{
			MsgId:   "20230711100610200000057A1ADYAGRZNGD",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:06:10"),
		},
		SignOffReq: admn_003_001_01.SignOffReqTCH{
			InstrId: "20230711200000057A1B123800017348098",
			InstgAgt: admn_003_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_003_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_003_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_003_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_003_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_003_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
		},
	},
}

func TestReadAdmn003(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn003.xml"))
	require.NoError(t, err)

	admn003 := &messages.Message{}
	err = xml.Unmarshal(input, admn003)
	require.NoError(t, err)

	expected := messages.NewAdmn003Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOffRequest = admn003Constant
	expected.SignOffRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOffRequest",
	}

	assert.Equal(t, expected, admn003)
}

func TestWriteAdmn003(t *testing.T) {
	input := messages.NewAdmn003Message()
	input.SignOffRequest = admn003Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn003.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOffRequest.Validate())
}
