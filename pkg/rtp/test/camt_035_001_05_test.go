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

	"github.com/moov-io/rtp20022/gen/camt_035_001_05"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var camt035Constant = &camt_035_001_05.DocumentTCH{
	PrtryFrmtInvstgtn: camt_035_001_05.ProprietaryFormatInvestigationV05TCH{
		Assgnmt: camt_035_001_05.CaseAssignment5TCH{
			Id: "M20230713234567891T1BOTS00251021992",
			Assgnr: camt_035_001_05.Party40ChoiceTCH{
				Agt: &camt_035_001_05.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_035_001_05.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_035_001_05.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
			Assgne: camt_035_001_05.Party40ChoiceTCH{
				Agt: &camt_035_001_05.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_035_001_05.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_035_001_05.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000010",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:48:38"),
		},
		Case: camt_035_001_05.Case5TCH{
			Id: "M20230713234567891T1BOTS00251021992",
			Cretr: camt_035_001_05.Party40ChoiceTCH2{
				Pty: &camt_035_001_05.PartyIdentification135TCH{
					Nm: "MRS. GREEN",
				},
			},
		},
		PrtryData: camt_035_001_05.ProprietaryData7TCHTCH{
			Tp: camt_035_001_05.Max35TextTCH3Ack,
			Data: camt_035_001_05.ProprietaryData6ReducedTCH{
				OrigRefs: camt_035_001_05.TransactionReferences8ReducedTCH{
					InstrId:    "20230713200000057A1BFSTF01198898349",
					EndToEndId: "EOTS",
					TxId:       "20230713200000057A1BFSTF01198898349",
				},
			},
		},
	},
}

func TestReadCamt035(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt035.xml"))
	require.NoError(t, err)

	camt035 := &messages.Message{}
	err = xml.Unmarshal(input, camt035)
	require.NoError(t, err)

	expected := messages.NewCamt035Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.Acknowledgement = camt035Constant
	expected.Acknowledgement.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Acknowledgement",
	}

	assert.Equal(t, expected, camt035)
}

func TestWriteCamt035(t *testing.T) {
	input := messages.NewCamt035Message()
	input.Acknowledgement = camt035Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "camt035.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.Acknowledgement.Validate())
}
