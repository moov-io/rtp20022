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

	"github.com/moov-io/rtp20022/gen/admn_001_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn001Constant = &admn_001_001_01.DocumentTCH{
	AdmnSignOnReq: admn_001_001_01.SignOnRequestTCH{
		GrpHdr: admn_001_001_01.GrpHdrTCH{
			MsgId:   "20230713142246200000057A1QD4EHWRNKD",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T14:22:46"),
		},
		SignOnReq: admn_001_001_01.SignOnReqTCH{
			InstrId: "20230713200000057A1BDCBA88966527298",
			InstgAgt: admn_001_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_001_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_001_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			InstdAgt: admn_001_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_001_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_001_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
		},
	},
}

func TestReadAdmn001(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn001.xml"))
	require.NoError(t, err)

	admn001 := &messages.Message{}
	err = xml.Unmarshal(input, admn001)
	require.NoError(t, err)

	expected := messages.NewAdmn001Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOnRequest = admn001Constant
	expected.SignOnRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOnRequest",
	}

	assert.Equal(t, expected, admn001)
}

func TestWriteAdmn001(t *testing.T) {
	input := messages.NewAdmn001Message()
	input.SignOnRequest = admn001Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn001.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOnRequest.Validate())
}
