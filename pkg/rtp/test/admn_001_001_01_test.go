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
			MsgId:   "2016021810064302120020101B61NHTCSG6",
			CreDtTm: rtp.UnmarshalISODateTime("2016-02-18T10:06:43"),
		},
		SignOnReq: admn_001_001_01.SignOnReqTCH{
			InstrId: "2016021812345678901BD61U00000000006",
			InstgAgt: admn_001_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_001_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_001_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "990000001T1",
					},
				},
			},
			InstdAgt: admn_001_001_01.BranchAndFinancialInstitutionIdentification4ADMN{
				FinInstnId: admn_001_001_01.FinancialInstitutionIdentification7ADMN{
					ClrSysMmbId: admn_001_001_01.ClearingSystemMemberIdentification2ADMN{
						MmbId: "XXXXXXXXXXX",
					},
				},
			},
		},
	},
}

func TestReadAdmn001(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admn001.RTP.xml"))
	require.NoError(t, err)

	admn001 := &messages.Message{}
	err = xml.Unmarshal(input, admn001)
	require.NoError(t, err)

	expected := messages.NewAdmn001Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
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

	expected, err := os.ReadFile(filepath.Join("testdata", "admn001.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
