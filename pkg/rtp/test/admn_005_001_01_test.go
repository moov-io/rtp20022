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

	"github.com/moov-io/rtp20022/gen/admn_005_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admn005Constant = &admn_005_001_01.Document{
	AdmnEchoReq: admn_005_001_01.EchoRequest{
		GrpHdr: admn_005_001_01.GrpHdr{
			MsgId:   "20230125125943990000001T10589310786",
			CreDtTm: rtp.UnmarshalISODateTime("2023-01-25T12:59:43"),
		},
		EchoTxInf: admn_005_001_01.EchoTxInf{
			FnctnCd: admn_005_001_01.EchoCode731,
			InstrId: "20230125990000001T1HQYBV00589310786",
			InstgAgt: admn_005_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_005_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_005_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_005_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_005_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_005_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "200000057A1",
					},
				},
			},
		},
	},
}

func TestReadAdmn005(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn005.RTP.xml"))
	require.NoError(t, err)

	admn005 := &messages.HdrAndData{}
	err = xml.Unmarshal(input, admn005)
	require.NoError(t, err)

	expected := messages.NewAdmn005Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.EchoRequest = admn005Constant
	expected.EchoRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "EchoRequest",
	}

	assert.Equal(t, expected, admn005)
}

func TestWriteAdmn005(t *testing.T) {
	input := messages.NewAdmn005Message()
	input.EchoRequest = admn005Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn005.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
