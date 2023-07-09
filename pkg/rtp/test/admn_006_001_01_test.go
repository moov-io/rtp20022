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

	"github.com/moov-io/rtp20022/gen/admn_006_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn006Constant = &admn_006_001_01.DocumentTCH{
	AdmnEchoResp: admn_006_001_01.EchoResponseTCH{
		GrpHdr: admn_006_001_01.GrpHdrTCH{
			MsgId:   "20230125125943990000001T10589310786",
			CreDtTm: rtp.UnmarshalISODateTime("2023-01-25T12:59:43"),
		},
		EchoResponse: admn_006_001_01.EchoResp{
			InstgAgt: admn_006_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_006_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_006_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_006_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_006_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_006_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
			OrgnlInstrId: "20230125990000001T1HQYBV00589310786",
			FnctnCd:      admn_006_001_01.EchoCode731,
			TxSts:        admn_006_001_01.TransactionIndividualStatus3CodeEchoActc,
		},
	},
}

func TestReadAdmn006(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn006.RTP.xml"))
	require.NoError(t, err)

	admn006 := &messages.Message{}
	err = xml.Unmarshal(input, admn006)
	require.NoError(t, err)

	expected := messages.NewAdmn006Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.EchoResponse = admn006Constant
	expected.EchoResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "EchoResponse",
	}

	assert.Equal(t, expected, admn006)
}

func TestWriteAdmn006(t *testing.T) {
	input := messages.NewAdmn006Message()
	input.EchoResponse = admn006Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn006.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
