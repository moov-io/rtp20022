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

	"github.com/moov-io/rtp20022/gen/admn_002_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn002ACTCConstant = &admn_002_001_01.DocumentTCH{
	AdmnSignOnResp: admn_002_001_01.SignOnResponseTCH{
		GrpHdr: admn_002_001_01.GrpHdrTCH{
			MsgId:   "20230713142247990000001T11027889234",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T14:22:47"),
		},
		SignOnResp: admn_002_001_01.SignOnRespTCH{
			OrgnlInstrId: "20230713200000057A1BDCBA88966527298",
			InstgAgt: admn_002_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_002_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_002_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_002_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_002_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_002_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			Sts: admn_002_001_01.TransactionGroupStatus3CodeAdminActc,
		},
	},
}

var admn002RJCTConstant = &admn_002_001_01.DocumentTCH{
	AdmnSignOnResp: admn_002_001_01.SignOnResponseTCH{
		GrpHdr: admn_002_001_01.GrpHdrTCH{
			MsgId:   "20230713142800990000001T10285284366",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T14:28:00"),
		},
		SignOnResp: admn_002_001_01.SignOnRespTCH{
			OrgnlInstrId: "20230713200000057A1BDCBA48308827298",
			InstgAgt: admn_002_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_002_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_002_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_002_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_002_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_002_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			Sts: admn_002_001_01.TransactionGroupStatus3CodeAdminRjct,
			StsRsnInf: &admn_002_001_01.StatusReasonInformation8TCH{
				Rsn: admn_002_001_01.StatusReason6Choice{
					Prtry: rtp.Ptr(admn_002_001_01.ProprietaryReasonCode("9948")),
				},
			},
		},
	},
}

func TestReadAdmn002ACTC(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn002.ACTC.xml"))
	require.NoError(t, err)

	admn002 := &messages.Message{}
	err = xml.Unmarshal(input, admn002)
	require.NoError(t, err)

	expected := messages.NewAdmn002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOnResponse = admn002ACTCConstant
	expected.SignOnResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOnResponse",
	}

	assert.Equal(t, expected, admn002)
}

func TestWriteAdmn002ACTC(t *testing.T) {
	input := messages.NewAdmn002Message()
	input.SignOnResponse = admn002ACTCConstant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn002.ACTC.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOnResponse.Validate())
}

func TestReadAdmn002RJCT(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn002.RJCT.xml"))
	require.NoError(t, err)

	admn002 := &messages.Message{}
	err = xml.Unmarshal(input, admn002)
	require.NoError(t, err)

	expected := messages.NewAdmn002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOnResponse = admn002RJCTConstant
	expected.SignOnResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOnResponse",
	}

	assert.Equal(t, expected, admn002)
}

func TestWriteAdmn002RJCT(t *testing.T) {
	input := messages.NewAdmn002Message()
	input.SignOnResponse = admn002RJCTConstant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn002.RJCT.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOnResponse.Validate())
}
