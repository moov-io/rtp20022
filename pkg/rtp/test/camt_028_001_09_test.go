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

	"github.com/moov-io/rtp20022/gen/camt_028_001_09"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var camt028Constant = &camt_028_001_09.DocumentTCH{
	AddtlPmtInf: camt_028_001_09.AdditionalPaymentInformationV09TCH{
		Assgnmt: camt_028_001_09.CaseAssignment5TCH{
			Id: "M20230713234567891T1BOTS01520718933",
			Assgnr: camt_028_001_09.Party40ChoiceTCH{
				Agt: &camt_028_001_09.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_028_001_09.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_028_001_09.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
			Assgne: camt_028_001_09.Party40ChoiceTCH{
				Agt: &camt_028_001_09.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_028_001_09.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_028_001_09.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000010",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:28:05"),
		},
		Case: camt_028_001_09.Case5TCH{
			Id: "M20230713BOTS00075450887",
			Cretr: camt_028_001_09.Party40ChoiceTCH2{
				Pty: &camt_028_001_09.PartyIdentification135TCH{
					Nm: "Participant Valid model",
				},
			},
		},
		Undrlyg: camt_028_001_09.UnderlyingTransaction5ChoiceTCH{
			IntrBk: &camt_028_001_09.UnderlyingPaymentTransaction4TCH{
				OrgnlGrpInf: camt_028_001_09.UnderlyingGroupInformation1TCH2{
					OrgnlMsgId:   "M20230713200000057A1BOTS00781766046",
					OrgnlMsgNmId: camt_028_001_09.OrigMsgNameTCH2(camt_028_001_09.OrigMsgNamePacs00800108),
				},
				OrgnlInstrId:    "20230713200000057A1BWLMG01424005165",
				OrgnlEndToEndId: rtp.Ptr(camt_028_001_09.Max35Text("EOTS")),
				OrgnlTxId:       "20230713200000057A1BWLMG01424005165",
				OrgnlIntrBkSttlmAmt: camt_028_001_09.ActiveOrHistoricCurrencyAndAmount{
					Value: 69622.06,
					Ccy:   camt_028_001_09.ActiveOrHistoricCurrencyCodeUsd,
				},
				OrgnlIntrBkSttlmDt: rtp.UnmarshalISODate("2023-07-13"),
			},
		},
		Inf: camt_028_001_09.PaymentComplementaryInformation8TCH{
			InstrId:    rtp.Ptr(camt_028_001_09.Max35Text("20230713200000057A1BWLMG01424005165")),
			EndToEndId: rtp.Ptr(camt_028_001_09.Max35Text("EOTS")),
			TxId:       rtp.Ptr(camt_028_001_09.Max35Text("20230713200000057A1BWLMG01424005165")),
		},
	},
}

func TestReadCamt028(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt028.xml"))
	require.NoError(t, err)

	camt028 := &messages.Message{}
	err = xml.Unmarshal(input, camt028)
	require.NoError(t, err)

	expected := messages.NewCamt028Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ResponseRequestForInformation = camt028Constant
	expected.ResponseRequestForInformation.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ResponseRequestForInformation",
	}

	assert.Equal(t, expected, camt028)
}

func TestWriteCamt028(t *testing.T) {
	input := messages.NewCamt028Message()
	input.ResponseRequestForInformation = camt028Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "camt028.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
