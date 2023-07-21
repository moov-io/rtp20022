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

	"github.com/moov-io/rtp20022/gen/admn_007_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn007Constant = &admn_007_001_01.DocumentTCH{
	DBRptReq: admn_007_001_01.DatabaseReportRequestTCH{
		GrpHdr: admn_007_001_01.GrpHdrTCH{
			MsgId:   "20230714151738200000057A1ODZIGYJNED",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-14T15:17:38"),
		},
		DBRptInf: admn_007_001_01.DatabaseReportInformationTCH{
			RptCd:   admn_007_001_01.ReportCodeTCHAvlbty,
			InstrId: "20230714200000057A1B123872285226398",
			InstgAgt: admn_007_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_007_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_007_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			InstdAgt: admn_007_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_007_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_007_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
		},
	},
}

func TestReadAdmn007(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn007.xml"))
	require.NoError(t, err)

	admn007 := &messages.Message{}
	err = xml.Unmarshal(input, admn007)
	require.NoError(t, err)

	expected := messages.NewAdmn007Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ParticipantReport = admn007Constant
	expected.ParticipantReport.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ParticipantReport",
	}

	assert.Equal(t, expected, admn007)
}

func TestWriteAdmn007(t *testing.T) {
	input := messages.NewAdmn007Message()
	input.ParticipantReport = admn007Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn007.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.ParticipantReport.Validate())
}
