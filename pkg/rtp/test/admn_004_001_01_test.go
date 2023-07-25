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

var admn004ACTCConstant = &admn_004_001_01.DocumentTCH{
	AdmnSignOffResp: admn_004_001_01.SignOffResponseTCH{
		GrpHdr: admn_004_001_01.GrpHdrTCH{
			MsgId:   "20230711100611990000001T10310971405",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:06:11"),
		},
		SignOffResp: admn_004_001_01.SignOffRespTCH{
			OrgnlInstrId: "20230711200000057A1B123800017348098",
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

var admn004RJCTConstant = &admn_004_001_01.DocumentTCH{
	AdmnSignOffResp: admn_004_001_01.SignOffResponseTCH{
		GrpHdr: admn_004_001_01.GrpHdrTCH{
			MsgId:   "20230711100706990000001T11713135631",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:07:06"),
		},
		SignOffResp: admn_004_001_01.SignOffRespTCH{
			OrgnlInstrId: "20230711200000057A1B123857452448098",
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
			Sts: admn_004_001_01.TransactionGroupStatus3CodeAdminRjct,
			StsRsnInf: &admn_004_001_01.StatusReasonInformation8TCH{
				Rsn: admn_004_001_01.StatusReason6Choice{
					Prtry: rtp.Ptr(admn_004_001_01.ProprietaryReasonCode("9948")),
				},
			},
		},
	},
}

func TestReadAdmn004ACTC(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn004.ACTC.xml"))
	require.NoError(t, err)

	admn004 := &messages.Message{}
	err = xml.Unmarshal(input, admn004)
	require.NoError(t, err)

	expected := messages.NewAdmn004Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOffResponse = admn004ACTCConstant
	expected.SignOffResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOffResponse",
	}

	assert.Equal(t, expected, admn004)
}

func TestWriteAdmn004ACTC(t *testing.T) {
	input := messages.NewAdmn004Message()
	input.SignOffResponse = admn004ACTCConstant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn004.ACTC.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOffResponse.Validate())
}

func TestReadAdmn004RJCT(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn004.RJCT.xml"))
	require.NoError(t, err)

	admn004 := &messages.Message{}
	err = xml.Unmarshal(input, admn004)
	require.NoError(t, err)

	expected := messages.NewAdmn004Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOffResponse = admn004RJCTConstant
	expected.SignOffResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOffResponse",
	}

	assert.Equal(t, expected, admn004)
}

func TestWriteAdmn004RJCT(t *testing.T) {
	input := messages.NewAdmn004Message()
	input.SignOffResponse = admn004RJCTConstant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn004.RJCT.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.SignOffResponse.Validate())
}
