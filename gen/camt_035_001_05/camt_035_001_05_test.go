package camt_035_001_05_test

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

var camt035Constant = &camt_035_001_05.Document{
	PrtryFrmtInvstgtn: camt_035_001_05.ProprietaryFormatInvestigationV05{
		Assgnmt: camt_035_001_05.CaseAssignment5{
			Id: "M20210701000000032A1B00000008088595",
			Assgnr: camt_035_001_05.Party40Choice{
				Agt: &camt_035_001_05.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: camt_035_001_05.FinancialInstitutionIdentification18{
						ClrSysMmbId: camt_035_001_05.ClearingSystemMemberIdentification2{
							MmbId: "000000032",
						},
					},
				},
			},
			Assgne: camt_035_001_05.Party40Choice{
				Agt: &camt_035_001_05.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: camt_035_001_05.FinancialInstitutionIdentification18{
						ClrSysMmbId: camt_035_001_05.ClearingSystemMemberIdentification2{
							MmbId: "000000010",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2021-07-01T11:07:14"),
		},
		Case: camt_035_001_05.Case5{
			Id: "M20210701000000032A1B00000008088595",
			Cretr: camt_035_001_05.Party40Choice{
				Pty: &camt_035_001_05.PartyIdentification135{
					Nm: "TCH",
				},
			},
		},
		PrtryData: camt_035_001_05.ProprietaryData7TCH{
			Tp: camt_035_001_05.Max35Text(camt_035_001_05.Max35TextTCH3Ack),
			Data: camt_035_001_05.ProprietaryData6Reduced{
				Ustrd: rtp.Ptr(camt_035_001_05.Max140Text("Information to the ACK")),
				OrigRefs: camt_035_001_05.TransactionReferences8Reduced{
					InstrId:    "20210701000000032A1B000000000047075",
					EndToEndId: "E2E-Ref001",
					TxId:       "20210701000000032A1B000000000047075",
				},
			},
		},
	},
}

func TestReadCamt035(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt035.RTP.xml"))
	require.NoError(t, err)

	camt035 := &messages.HdrAndData{}
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

	expected, err := os.ReadFile(filepath.Join("testdata", "camt035.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
