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

	"github.com/moov-io/rtp20022/gen/admn_008_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn008Constant = &admn_008_001_01.DocumentTCH{
	DBAvlbtyRpt: admn_008_001_01.DatabaseAvailabilityReportTCH{
		GrpHdr: admn_008_001_01.GrpHdr{
			MsgId:   "20230714151907990000001T11065412146",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-14T15:19:07"),
		},
		DBRptRspn: admn_008_001_01.DatabaseReportReponseTCH{
			OrgnlInstrId: "20230714200000057A1B123882964326398",
			RptCd:        admn_008_001_01.ReportCodeTCHAvlbty,
			InstgAgt: admn_008_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_008_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_008_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_008_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_008_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_008_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			TxSts: admn_008_001_01.TransactionIndividualStatus3CodeEchoActc,
		},
		AvlbtyRpt: admn_008_001_01.AvailabilityReportTCH{
			Cnnctn: &admn_008_001_01.ConnectionTCH{
				CnnctnId: []admn_008_001_01.Max20AlphaNumericText{
					"234",
				},
			},
			AvlbtyPtcpt: &admn_008_001_01.AvailabilityParticipantTCH{
				PtcptSgnOff: &admn_008_001_01.ParticipantSignOffTCH{
					PtcptId: []admn_008_001_01.Min11Max11Text{
						"HDBF546BF64",
					},
				},
				PtcptSspd: &admn_008_001_01.ParticipantSuspendedTCH{
					PtcptId: []admn_008_001_01.Min11Max11Text{
						"KNJBDY5WV4F",
					},
				},
			},
		},
	},
}

func TestReadAdmn008(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn008.xml"))
	require.NoError(t, err)

	admn008 := &messages.Message{}
	err = xml.Unmarshal(input, admn008)
	require.NoError(t, err)

	expected := messages.NewAdmn008Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ParticipantReportResponse = admn008Constant
	expected.ParticipantReportResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ParticipantReportResponse",
	}

	assert.Equal(t, expected, admn008)
}

func TestWriteAdmn008(t *testing.T) {
	input := messages.NewAdmn008Message()
	input.ParticipantReportResponse = admn008Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn008.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
