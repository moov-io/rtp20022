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
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var pacs008Constant = &pacs_008_001_08.DocumentTCH{
	FIToFICstmrCdtTrf: pacs_008_001_08.FIToFICustomerCreditTransferV08TCH{
		GrpHdr: pacs_008_001_08.GroupHeader93TCH{
			MsgId:   "M20230713234567891T1BOTS02083825666",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:29:36"),
			NbOfTxs: "1",
			TtlIntrBkSttlmAmt: pacs_008_001_08.ActiveCurrencyAndAmount{
				Value: 46154.49,
				Ccy:   pacs_008_001_08.ActiveCurrencyCodeUsd,
			},
			IntrBkSttlmDt: rtp.UnmarshalISODate("2023-07-13"),
			SttlmInf: pacs_008_001_08.SettlementInstruction7TCH{
				SttlmMtd: pacs_008_001_08.SettlementMethod1CodeClrg,
				ClrSys: pacs_008_001_08.ClearingSystemIdentification3Choice{
					Cd: rtp.Ptr(pacs_008_001_08.ExternalCashClearingSystem1CodeTch),
				},
			},
		},
		CdtTrfTxInf: pacs_008_001_08.CreditTransferTransaction39TCH{
			PmtId: pacs_008_001_08.PaymentIdentification7TCH{
				InstrId:    pacs_008_001_08.Max35TextTCH2("20230713234567891T1BCXIA01868416991"),
				EndToEndId: "EOTS",
				TxId:       pacs_008_001_08.Max35Text("20230713234567891T1BCXIA01868416991"),
				ClrSysRef:  rtp.Ptr(pacs_008_001_08.Max35Text("001")),
			},
			PmtTpInf: pacs_008_001_08.PaymentTypeInformation28TCH{
				SvcLvl: pacs_008_001_08.ServiceLevel8Choice{
					Cd: rtp.Ptr(pacs_008_001_08.ExternalServiceLevel1CodeSdva),
				},
				LclInstrm: pacs_008_001_08.LocalInstrument2Choice{
					Prtry: rtp.Ptr(pacs_008_001_08.LocalPropStandard),
				},
				CtgyPurp: pacs_008_001_08.CategoryPurpose1Choice{
					Prtry: rtp.Ptr(pacs_008_001_08.CatePurpPropBusiness),
				},
			},
			IntrBkSttlmAmt: pacs_008_001_08.ActiveCurrencyAndAmount{
				Value: 46154.49,
				Ccy:   pacs_008_001_08.ActiveCurrencyCodeUsd,
			},
			ChrgBr: pacs_008_001_08.ChargeBearerType1CodeSlev,
			InstgAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_008_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "234567891",
					},
				},
			},
			InstdAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_008_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "000000010",
					},
				},
			},
			Dbtr: pacs_008_001_08.PartyIdentification135TCH3{
				Nm: pacs_008_001_08.Max140Text("MRS. GREEN"),
				PstlAdr: &pacs_008_001_08.PostalAddress24TCH{
					StrtNm:      pacs_008_001_08.Max70Text("Broadway"),
					BldgNb:      rtp.Ptr(pacs_008_001_08.Max16Text("1500")),
					PstCd:       pacs_008_001_08.Max16Text("NY 12345"),
					TwnNm:       pacs_008_001_08.Max35Text("New York"),
					CtrySubDvsn: pacs_008_001_08.Max35Text("NY"),
					Ctry:        pacs_008_001_08.CountryCode("US"),
				},
				Id: &pacs_008_001_08.Party38ChoiceTCH3{
					PrvtId: &pacs_008_001_08.PersonIdentification13TCH2{
						DtAndPlcOfBirth: pacs_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     rtp.UnmarshalISODate("1984-01-01"),
							CityOfBirth: "New York",
							CtryOfBirth: "US",
						},
					},
				},
			},
			DbtrAcct: pacs_008_001_08.CashAccount38TCH{
				Id: pacs_008_001_08.AccountIdentification4Choice{
					Othr: &pacs_008_001_08.GenericAccountIdentification1{
						Id: pacs_008_001_08.Max34Text("US88664715164441"),
					},
				},
			},
			DbtrAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_008_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "234567891",
					},
				},
			},
			CdtrAgt: pacs_008_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_008_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_008_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "000000010",
					},
				},
			},
			Cdtr: pacs_008_001_08.PartyIdentification135TCH3{
				Nm: pacs_008_001_08.Max140Text("Participant Valid model"),
				PstlAdr: &pacs_008_001_08.PostalAddress24TCH{
					StrtNm:      pacs_008_001_08.Max70Text("NORTH AVE"),
					BldgNb:      rtp.Ptr(pacs_008_001_08.Max16Text("1123")),
					PstCd:       pacs_008_001_08.Max16Text("12344"),
					TwnNm:       pacs_008_001_08.Max35Text("LOS ANGELES"),
					CtrySubDvsn: pacs_008_001_08.Max35Text("LA"),
					Ctry:        pacs_008_001_08.CountryCode("US"),
				},
				Id: &pacs_008_001_08.Party38ChoiceTCH3{
					PrvtId: &pacs_008_001_08.PersonIdentification13TCH2{
						DtAndPlcOfBirth: pacs_008_001_08.DateAndPlaceOfBirth1{
							BirthDt:     rtp.UnmarshalISODate("1989-01-09"),
							CityOfBirth: "LOS ANGELES",
							CtryOfBirth: "US",
						},
					},
				},
			},
			CdtrAcct: pacs_008_001_08.CashAccount38TCH2{
				Id: pacs_008_001_08.AccountIdentification4Choice{
					Othr: &pacs_008_001_08.GenericAccountIdentification1{
						Id: pacs_008_001_08.Max34Text("42341169728762787"),
					},
				},
			},
			RltdRmtInf: &pacs_008_001_08.RemittanceLocation7TCH{
				RmtId: rtp.Ptr(pacs_008_001_08.Max35Text("20230713152936OTS00752775034")),
			},
		},
	},
}

func TestReadPacs008(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pacs008.xml"))
	require.NoError(t, err)

	pacs008 := &messages.Message{}
	err = xml.Unmarshal(input, pacs008)
	require.NoError(t, err)

	expected := messages.NewPacs008Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.CreditTransfer = pacs008Constant
	expected.CreditTransfer.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "CreditTransfer",
	}

	assert.Equal(t, expected, pacs008)
}

func TestWritePacs008(t *testing.T) {
	input := messages.NewPacs008Message()
	input.CreditTransfer = pacs008Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pacs008.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.CreditTransfer.Validate())
}
