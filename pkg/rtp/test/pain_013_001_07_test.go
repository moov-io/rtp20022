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

	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/gen/pain_013_001_07"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var pain013Constant = &pain_013_001_07.DocumentTCH{
	CdtrPmtActvtnReq: pain_013_001_07.CreditorPaymentActivationRequestV07TCH{
		GrpHdr: pain_013_001_07.GroupHeader78TCH{
			MsgId:   "M20210701000000032A1B00000008088839",
			CreDtTm: rtp.UnmarshalISODateTime("2021-07-01T11:13:14"),
			NbOfTxs: "1",
			InitgPty: pain_013_001_07.PartyIdentification135TCH{
				Id: pain_013_001_07.Party38ChoiceTCH{
					OrgId: &pain_013_001_07.OrganisationIdentification29TCH{
						Othr: []pain_013_001_07.GenericOrganisationIdentification1TCH{
							{
								Id: "000000032",
							},
						},
					},
				},
			},
		},
		PmtInf: pain_013_001_07.PaymentInstruction31TCH{
			PmtInfId: pain_013_001_07.Max35TextTCH3("20210701000000032A1B000000000047080"),
			PmtMtd:   pain_013_001_07.PaymentMethod7CodeTrf,
			ReqdExctnDt: pain_013_001_07.DateAndDateTime2Choice{
				Dt: rtp.Ptr(rtp.UnmarshalISODate("2021-07-01")),
			},
			XpryDt: pain_013_001_07.DateAndDateTime2Choice{
				Dt: rtp.Ptr(rtp.UnmarshalISODate("2021-07-01")),
			},
			Dbtr: pain_013_001_07.PartyIdentification135TCH2{
				Nm: pain_013_001_07.Max140Text("000000010"),
				Id: &pain_013_001_07.Party38ChoiceTCH2{
					PrvtId: &pain_013_001_07.PersonIdentification13TCH{
						DtAndPlcOfBirth: pain_013_001_07.DateAndPlaceOfBirth1{
							BirthDt:     rtp.UnmarshalISODate("1980-12-06"),
							CityOfBirth: "NY",
							CtryOfBirth: "US",
						},
					},
				},
			},
			DbtrAcct: pain_013_001_07.CashAccount38TCH{
				Id: pain_013_001_07.AccountIdentification4Choice{
					Othr: &pain_013_001_07.GenericAccountIdentification1{
						Id: "66906867108802622",
					},
				},
			},
			DbtrAgt: pain_013_001_07.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pain_013_001_07.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pain_013_001_07.ClearingSystemMemberIdentification2TCH{
						MmbId: "000000010",
					},
				},
			},
			CdtTrfTx: pain_013_001_07.CreditTransferTransaction35TCH{
				PmtId: pain_013_001_07.PaymentIdentification6{
					InstrId:    rtp.Ptr(pain_013_001_07.Max35Text("20210701000000032A1B000000000047080")),
					EndToEndId: "E2E-Ref001",
				},
				PmtTpInf: pain_013_001_07.PaymentTypeInformation26TCH{
					SvcLvl: pain_013_001_07.ServiceLevel8Choice{
						Cd: rtp.Ptr(pain_013_001_07.ExternalServiceLevel1CodeSdva),
					},
					LclInstrm: pain_013_001_07.LocalInstrument2ChoiceTCH{
						Prtry: rtp.Ptr(pain_013_001_07.LocalPropTCHStandard),
					},
					CtgyPurp: pain_013_001_07.CategoryPurpose1Choice{
						Prtry: rtp.Ptr(pain_013_001_07.CatePurpPropConsumer),
					},
				},
				Amt: pain_013_001_07.AmountType4Choice{
					InstdAmt: &pain_013_001_07.ActiveOrHistoricCurrencyAndAmount{
						Value: 15.00,
						Ccy:   pain_013_001_07.ActiveOrHistoricCurrencyCodeUsd,
					},
				},
				ChrgBr: pain_013_001_07.ChargeBearerType1CodeSlev,
				CdtrAgt: pain_013_001_07.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: pain_013_001_07.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: pain_013_001_07.ClearingSystemMemberIdentification2TCH{
							MmbId: pain_013_001_07.Max35TextTCH2("000000032"),
						},
					},
				},
				Cdtr: pain_013_001_07.PartyIdentification135TCH4{
					Nm: pain_013_001_07.Max140Text("000000032"),
					Id: &pain_013_001_07.Party38ChoiceTCH2{
						PrvtId: &pain_013_001_07.PersonIdentification13TCH{
							DtAndPlcOfBirth: pain_013_001_07.DateAndPlaceOfBirth1{
								BirthDt:     rtp.UnmarshalISODate("1976-02-14"),
								CityOfBirth: "GL",
								CtryOfBirth: "RO",
							},
						},
					},
				},
				CdtrAcct: pain_013_001_07.CashAccount38TCH2{
					Id: pain_013_001_07.AccountIdentification4Choice{
						Othr: &pain_013_001_07.GenericAccountIdentification1{
							Id: pain_013_001_07.Max34Text("46684938893141836"),
						},
					},
				},
				InstrForCdtrAgt: []*pain_013_001_07.InstructionForCreditorAgent1TCH{
					{
						Cd:       pain_013_001_07.Instruction3CodeTCHReci,
						InstrInf: pain_013_001_07.Max140Text("VLTK"),
					},
				},
			},
		},
	},
}

func TestReadPain013(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pain013.RTP.xml"))
	require.NoError(t, err)

	pain013 := &messages.Message{}
	err = xml.Unmarshal(input, pain013)
	require.NoError(t, err)

	expected := messages.NewPain013Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.PaymentRequest = pain013Constant
	expected.PaymentRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "PaymentRequest",
	}

	assert.Equal(t, expected, pain013)
}

func TestWritePain013(t *testing.T) {
	input := messages.NewPain013Message()
	input.PaymentRequest = pain013Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pain013.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
