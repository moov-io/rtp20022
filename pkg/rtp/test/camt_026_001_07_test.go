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

	"github.com/moov-io/rtp20022/gen/camt_026_001_07"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var camt026Constant = &camt_026_001_07.DocumentTCH{
	UblToApply: camt_026_001_07.UnableToApplyV07TCH{
		Assgnmt: camt_026_001_07.CaseAssignment5TCH{
			Id: "M20230713234567891T1BOTS02112879685",
			Assgnr: camt_026_001_07.Party40ChoiceTCH{
				Agt: &camt_026_001_07.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_026_001_07.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_026_001_07.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
			Assgne: camt_026_001_07.Party40ChoiceTCH{
				Agt: &camt_026_001_07.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_026_001_07.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_026_001_07.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000010",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:25:46"),
		},
		Case: camt_026_001_07.Case5TCH{
			Id: "M20230713234567891T1BOTS02112879685",
			Cretr: camt_026_001_07.Party40ChoiceTCH2{
				Pty: &camt_026_001_07.PartyIdentification135TCH{
					Nm: "MRS. GREEN",
				},
			},
		},
		Undrlyg: camt_026_001_07.UnderlyingTransaction5ChoiceTCH{
			IntrBk: &camt_026_001_07.UnderlyingPaymentTransaction4TCH{
				OrgnlGrpInf: camt_026_001_07.UnderlyingGroupInformation1TCH2{
					OrgnlMsgId:   "M20230713234567891T1BOTS01225082200",
					OrgnlMsgNmId: camt_026_001_07.OrigMsgNameTCH2(camt_026_001_07.OrigMsgNamePacs00800108),
					OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:25:46"),
				},
				OrgnlInstrId:    "20230713234567891T1BKHAH00545741922",
				OrgnlEndToEndId: rtp.Ptr(camt_026_001_07.Max35Text("EOTS")),
				OrgnlTxId:       "20230713234567891T1BKHAH00545741922",
				OrgnlIntrBkSttlmAmt: camt_026_001_07.ActiveOrHistoricCurrencyAndAmount{
					Value: 53468.47,
					Ccy:   camt_026_001_07.ActiveOrHistoricCurrencyCodeUsd,
				},
				OrgnlIntrBkSttlmDt: rtp.UnmarshalISODate("2023-07-13"),
			},
		},
		Justfn: camt_026_001_07.UnableToApplyJustification3Choice{
			MssngOrIncrrctInf: &camt_026_001_07.MissingOrIncorrectInformation3{
				MssngInf: []*camt_026_001_07.UnableToApplyMissing1{
					{
						Cd: "MS01",
					},
				},
			},
		},
	},
}

func TestReadCamt026(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt026.xml"))
	require.NoError(t, err)

	camt026 := &messages.Message{}
	err = xml.Unmarshal(input, camt026)
	require.NoError(t, err)

	expected := messages.NewCamt026Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.RequestForInformation = camt026Constant
	expected.RequestForInformation.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "RequestForInformation",
	}

	assert.Equal(t, expected, camt026)
}

func TestWriteCamt026(t *testing.T) {
	input := messages.NewCamt026Message()
	input.RequestForInformation = camt026Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "camt026.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.RequestForInformation.Validate())
}
