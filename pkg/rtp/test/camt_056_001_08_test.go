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

var camt056Constant = &camt_056_001_08.Document{
	FIToFIPmtCxlReq: camt_056_001_08.FIToFIPaymentCancellationRequestV08{
		Assgnmt: camt_056_001_08.CaseAssignment5{
			Id: "M20210701000000010B1B00000008088677",
			Assgnr: camt_056_001_08.Party40Choice{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2{
							MmbId: "000000010",
						},
					},
				},
			},
			Assgne: camt_056_001_08.Party40Choice{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2{
							MmbId: "000000032",
						},
					},
				},
			},
			CreDtTm: rtp.UnmarshalISODateTime("2021-07-01T11:09:04"),
		},
		Case: camt_056_001_08.Case5{
			Id: "M20210701000000010B1B00000008088677",
			Cretr: camt_056_001_08.Party40Choice{
				Agt: &camt_056_001_08.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: camt_056_001_08.FinancialInstitutionIdentification18{
						ClrSysMmbId: camt_056_001_08.ClearingSystemMemberIdentification2{
							MmbId: "000000010",
						},
					},
				},
			},
		},
		Undrlyg: camt_056_001_08.UnderlyingTransaction23{
			OrgnlGrpInfAndCxl: camt_056_001_08.OriginalGroupHeader15{
				OrgnlMsgId:   "M20210701000000032A1B00000008088224",
				OrgnlMsgNmId: camt_056_001_08.OrigMsgNamePacs00800108,
				OrgnlCreDtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2021-06-21T11:09:04")),
			},
			TxInf: camt_056_001_08.PaymentTransaction106{
				OrgnlInstrId:    "20210701000000032A1B000000000047075",
				OrgnlEndToEndId: rtp.Ptr(camt_056_001_08.Max35Text("MYREF123")),
				OrgnlTxId:       "20210701000000032A1B000000000047075",
				OrgnlClrSysRef:  "001",
				OrgnlIntrBkSttlmAmt: camt_056_001_08.ActiveOrHistoricCurrencyAndAmount{
					Value: 10.00,
					Ccy:   camt_056_001_08.ActiveOrHistoricCurrencyCodeUsd,
				},
				OrgnlIntrBkSttlmDt: rtp.UnmarshalISODate("2021-07-01"),
				CxlRsnInf: camt_056_001_08.PaymentCancellationReason5{
					Orgtr: &camt_056_001_08.PartyIdentification135{
						Id: &camt_056_001_08.Party38Choice{
							OrgId: &camt_056_001_08.OrganisationIdentification29{
								Othr: &camt_056_001_08.GenericOrganisationIdentification1{
									Id: "000000010",
								},
							},
						},
					},
					Rsn: camt_056_001_08.CancellationReason33Choice{
						Cd: rtp.Ptr(camt_056_001_08.ExternalCancellationReason1Code(camt_056_001_08.ExternalCancellationReason1CodeCust)),
					},
					AddtlInf: rtp.Ptr(camt_056_001_08.Max105Text("CUST")),
				},
			},
		},
	},
}

func TestReadCamt056(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "camt056.RTP.xml"))
	require.NoError(t, err)

	camt056 := &messages.HdrAndData{}
	err = xml.Unmarshal(input, camt056)
	require.NoError(t, err)

	expected := messages.NewCamt056Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
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

	expected, err := os.ReadFile(filepath.Join("testdata", "camt056.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
