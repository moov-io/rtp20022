package admn_002_001_01_test

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

var admn002Constant = &admn_002_001_01.Document{
	AdmnSignOnResp: admn_002_001_01.SignOnResponse{
		GrpHdr: admn_002_001_01.GrpHdr{
			MsgId:   "20230110163014990000001T11084940492",
			CreDtTm: rtp.UnmarshalISODateTime("2023-01-10T16:30:14"),
		},
		SignOnResp: admn_002_001_01.SignOnResp{
			OrgnlInstrId: "20230110200000057A1B1238MD3QGUJMED3",
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
						MmbId: "990000001T1",
					},
				},
			},
			Sts: admn_002_001_01.TransactionGroupStatus3CodeAdminActc,
		},
	},
}

func TestReadAdmn002(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn002.RTP.xml"))
	require.NoError(t, err)

	admn002 := &messages.HdrAndData{}
	err = xml.Unmarshal(input, admn002)
	require.NoError(t, err)

	expected := messages.NewAdmn002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SignOnResponse = admn002Constant
	expected.SignOnResponse.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SignOnResponse",
	}

	assert.Equal(t, expected, admn002)
}

func TestWriteAdmn002(t *testing.T) {
	input := messages.NewAdmn002Message()
	input.SignOnResponse = admn002Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admn002.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
