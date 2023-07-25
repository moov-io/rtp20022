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

	"github.com/moov-io/rtp20022/gen/camt_056_001_08"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var camt056Constant = &camt_056_001_08.DocumentTCH{
	FIToFIPmtCxlReq: camt_056_001_08.FIToFIPaymentCancellationRequestV08TCH{
		Assgnmt: camt_056_001_08.CaseAssignment5TCH{
			Id: "M20230713234567891T1BOTS00384697493",
			Assgnr: camt_056_001_08.Party40ChoiceTCH{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
			Assgne: camt_056_001_08.Party40ChoiceTCH{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2TCH{
							MmbId: "000000032",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:28:43"),
		},
		Case: camt_056_001_08.Case5TCH{
			Id: "M20230713234567891T1BOTS00384697493",
			Cretr: camt_056_001_08.Party40ChoiceTCH{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
		},
		Undrlyg: camt_056_001_08.UnderlyingTransaction23TCH{
			OrgnlGrpInfAndCxl: camt_056_001_08.OriginalGroupHeader15{
				OrgnlMsgId:   "M20230713234567891T1BOTS00694297350",
				OrgnlMsgNmId: camt_056_001_08.OrigMsgNamePacs00800108,
				OrgnlCreDtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2023-07-13T15:28:42")),
			},
			TxInf: camt_056_001_08.PaymentTransaction106TCH{
				OrgnlInstrId:    "20230713234567891T1BLFFV00017549968",
				OrgnlEndToEndId: rtp.Ptr(camt_056_001_08.Max35Text("EOTS")),
				OrgnlTxId:       "20230713234567891T1BLFFV00017549968",
				OrgnlClrSysRef:  "001",
				OrgnlIntrBkSttlmAmt: camt_056_001_08.ActiveOrHistoricCurrencyAndAmount{
					Value: 42389.38,
					Ccy:   camt_056_001_08.ActiveOrHistoricCurrencyCodeUsd,
				},
				OrgnlIntrBkSttlmDt: rtp.UnmarshalISODate("2023-07-13"),
				CxlRsnInf: camt_056_001_08.PaymentCancellationReason5TCH{
					Orgtr: &camt_056_001_08.PartyIdentification135TCH{
						Nm: rtp.Ptr(camt_056_001_08.Max140Text("MRS. GREEN")),
						Id: &camt_056_001_08.Party38ChoiceTCH{
							OrgId: &camt_056_001_08.OrganisationIdentification29TCH{
								Othr: camt_056_001_08.GenericOrganisationIdentification1TCH{
									Id: "234567891",
								},
							},
						},
					},
					Rsn: camt_056_001_08.CancellationReason33ChoiceTCH{
						Cd: rtp.Ptr(camt_056_001_08.ExternalCancellationReason1CodeCust),
					},
				},
				OrgnlTxRef: &camt_056_001_08.OriginalTransactionReference28TCH{
					Dbtr: &camt_056_001_08.Party40ChoiceTCH2{
						Pty: &camt_056_001_08.PartyIdentification135TCH2{
							Nm: "MRS. GREEN",
							PstlAdr: &camt_056_001_08.PostalAddress24TCH{
								StrtNm:      "Broadway",
								BldgNb:      rtp.Ptr(camt_056_001_08.Max16Text("1500")),
								PstCd:       "NY 12345",
								TwnNm:       "New York",
								CtrySubDvsn: "NY",
								Ctry:        "US",
							},
							Id: &camt_056_001_08.Party38ChoiceTCH2{
								PrvtId: &camt_056_001_08.PersonIdentification13TCH{
									DtAndPlcOfBirth: camt_056_001_08.DateAndPlaceOfBirth1{
										BirthDt:     rtp.UnmarshalISODate("1984-01-01"),
										CityOfBirth: "New York",
										CtryOfBirth: "US",
									},
								},
							},
						},
					},
					Cdtr: &camt_056_001_08.Party40ChoiceTCH3{
						Pty: &camt_056_001_08.PartyIdentification135TCH3{
							Nm: "Participant Valid model",
							PstlAdr: &camt_056_001_08.PostalAddress24TCH{
								StrtNm:      "NORTH AVE",
								BldgNb:      rtp.Ptr(camt_056_001_08.Max16Text("1123")),
								PstCd:       "12344",
								TwnNm:       "LOS ANGELES",
								CtrySubDvsn: "LA",
								Ctry:        "US",
							},
							Id: &camt_056_001_08.Party38ChoiceTCH2{
								PrvtId: &camt_056_001_08.PersonIdentification13TCH{
									DtAndPlcOfBirth: camt_056_001_08.DateAndPlaceOfBirth1{
										BirthDt:     rtp.UnmarshalISODate("1989-01-09"),
										CityOfBirth: "LOS ANGELES",
										CtryOfBirth: "US",
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestReadCamt056(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt056.refund.xml"))
	require.NoError(t, err)

	camt056 := &messages.Message{}
	err = xml.Unmarshal(input, camt056)
	require.NoError(t, err)

	expected := messages.NewCamt056Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ReturnOfFunds = camt056Constant
	expected.ReturnOfFunds.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ReturnOfFunds",
	}

	assert.Equal(t, expected, camt056)
}

func TestWriteCamt056(t *testing.T) {
	input := messages.NewCamt056Message()
	input.ReturnOfFunds = camt056Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "camt056.refund.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.ReturnOfFunds.Validate())
}
