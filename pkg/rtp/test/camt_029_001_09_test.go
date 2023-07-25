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

	"github.com/moov-io/rtp20022/gen/camt_029_001_09"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var camt029Constant = &camt_029_001_09.DocumentTCH{
	RsltnOfInvstgtn: camt_029_001_09.ResolutionOfInvestigationV09TCH{
		Assgnmt: camt_029_001_09.CaseAssignment5TCH{
			Id: "M20230713234567891T1BOTS00670150991",
			Assgnr: camt_029_001_09.Party40ChoiceTCH{
				Agt: &camt_029_001_09.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_029_001_09.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_029_001_09.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
			Assgne: camt_029_001_09.Party40ChoiceTCH{
				Agt: &camt_029_001_09.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_029_001_09.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_029_001_09.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000010",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:49:09"),
		},
		Sts: camt_029_001_09.InvestigationStatus5Choice{
			Conf: rtp.Ptr(camt_029_001_09.ExternalInvestigationExecutionConfirmation1CodeIpay),
		},
		CxlDtls: camt_029_001_09.UnderlyingTransaction22TCH{
			OrgnlGrpInfAndSts: camt_029_001_09.OriginalGroupHeader14TCH{
				RslvdCase: camt_029_001_09.Case5TCH{
					Id: "M20230713BOTS00856763956",
					Cretr: camt_029_001_09.Party40ChoiceTCH{
						Agt: &camt_029_001_09.BranchAndFinancialInstitutionIdentification6TCH{
							FinInstnId: camt_029_001_09.FinancialInstitutionIdentification18TCH{
								ClrSysMmbId: camt_029_001_09.ClearingSystemMemberIdentification2TCH{
									MmbId: "000000010",
								},
							},
						},
					},
				},
				OrgnlMsgId:   "M20230713BOTS00856763956",
				OrgnlMsgNmId: camt_029_001_09.OrigMsgNameCamt05600108,
				OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:49:09"),
			},
			TxInfAndSts: &camt_029_001_09.PaymentTransaction102TCH{
				CxlStsId: rtp.Ptr(camt_029_001_09.Max35Text("M20230713234567891T1BOTS00893154088")),
				RsltnRltdInf: &camt_029_001_09.ResolutionData1TCH{
					IntrBkSttlmAmt: camt_029_001_09.ActiveOrHistoricCurrencyAndAmount{
						Value: 17272.16,
						Ccy:   camt_029_001_09.ActiveOrHistoricCurrencyCodeUsd,
					},
				},
			},
		},
	},
}

func TestReadCamt029(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt029.xml"))
	require.NoError(t, err)

	camt029 := &messages.Message{}
	err = xml.Unmarshal(input, camt029)
	require.NoError(t, err)

	expected := messages.NewCamt029Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ResponseReturnOfFunds = camt029Constant
	expected.ResponseReturnOfFunds.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ResponseReturnOfFunds",
	}

	assert.Equal(t, expected, camt029)
}

func TestWriteCamt029(t *testing.T) {
	input := messages.NewCamt029Message()
	input.ResponseReturnOfFunds = camt029Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "camt029.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.ResponseReturnOfFunds.Validate())
}
