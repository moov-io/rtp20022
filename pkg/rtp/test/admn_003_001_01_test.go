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

var admn003Constant = &admn_003_001_01.Document{
	AdmnSignOffReq: admn_003_001_01.SignOffRequest{
		GrpHdr: admn_003_001_01.GrpHdr{
			MsgId:   "2016021810064302120020101B61NHTCSG6",
			CreDtTm: rtp.UnmarshalISODateTime("2016-02-18T10:06:43"),
		},
		SignOffReq: admn_003_001_01.SignOffReq{
			InstrId: "2016021812345678901BD61U00000000006",
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
						MmbId: "XXXXXXXXXXX",
					},
				},
			},
		},
	},
}

func TestReadAdmn003(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn003.RTP.xml"))
	require.NoError(t, err)

	admn003 := &messages.HdrAndData{}
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

	expected, err := os.ReadFile(filepath.Join("testdata", "admn003.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
