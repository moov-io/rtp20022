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
			Id: "M20210701000000032A1B00000008088704",
			Assgnr: camt_029_001_09.Party40ChoiceTCH{
				Agt: &camt_029_001_09.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_029_001_09.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_029_001_09.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000032",
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
			CreDtTm: rtp.UnmarshalISODateTime("2021-07-01T11:09:58"),
		},
		Sts: camt_029_001_09.InvestigationStatus5Choice{
			Conf: rtp.Ptr(camt_029_001_09.ExternalInvestigationExecutionConfirmation1CodeIpay),
		},
		CxlDtls: camt_029_001_09.UnderlyingTransaction22TCH{
			OrgnlGrpInfAndSts: camt_029_001_09.OriginalGroupHeader14TCH{
				RslvdCase: camt_029_001_09.Case5TCH{
					Id: "M20210701000000010B1B00000008088677",
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
				OrgnlMsgId:   "M20210701000000010B1B00000008088677",
				OrgnlMsgNmId: camt_029_001_09.OrigMsgNameCamt05600108,
				OrgnlCreDtTm: rtp.UnmarshalISODateTime("2021-06-21T11:09:58"),
			},
		},
	},
}

func TestReadCamt029(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt029.RTP.xml"))
	require.NoError(t, err)

	camt029 := &messages.Message{}
	err = xml.Unmarshal(input, camt029)
	require.NoError(t, err)

	expected := messages.NewCamt029Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
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

	expected, err := os.ReadFile(filepath.Join("testdata", "camt029.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
