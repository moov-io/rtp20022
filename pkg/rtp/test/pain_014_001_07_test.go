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
	"github.com/moov-io/rtp20022/gen/pain_014_001_07"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var pain014ConstantACTC = &pain_014_001_07.DocumentTCH{
	CdtrPmtActvtnReqStsRpt: pain_014_001_07.CreditorPaymentActivationRequestStatusReportV07TCH{
		GrpHdr: pain_014_001_07.GroupHeader87TCH{
			MsgId:   "M20230713234567891T1BOTS01353800076",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:23:33"),
			InitgPty: pain_014_001_07.PartyIdentification135TCH{
				Id: pain_014_001_07.Party38ChoiceTCH{
					OrgId: &pain_014_001_07.OrganisationIdentification29TCH{
						Othr: pain_014_001_07.GenericOrganisationIdentification1TCH{
							Id: "234567891",
						},
					},
				},
			},
		},
		OrgnlGrpInfAndSts: pain_014_001_07.OriginalGroupInformation30TCH{
			OrgnlMsgId:   "M20230713200000057A1BOTS00812976160",
			OrgnlMsgNmId: pain_014_001_07.OrigMsgNamePain01300107,
			OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:23:33"),
			OrgnlNbOfTxs: "1",
		},
		OrgnlPmtInfAndSts: pain_014_001_07.OriginalPaymentInstruction31TCH{
			OrgnlPmtInfId: "20230713200000057A1BIWUY00608958679",
			TxInfAndSts: pain_014_001_07.PaymentTransaction104TCH{
				OrgnlEndToEndId: rtp.Ptr(pain_014_001_07.Max35Text("EOTS")),
				TxSts:           pain_014_001_07.ExternalPaymentTransactionStatus1CodeActc,
				PmtCondSts: &pain_014_001_07.PaymentConditionStatus1TCH{
					AccptdAmt: pain_014_001_07.ActiveCurrencyAndAmount{
						Value: 8352.45,
						Ccy:   pain_014_001_07.ActiveCurrencyCodeUsd,
					},
					GrntedPmt: true,
					EarlyPmt:  true,
				},
				OrgnlTxRef: pain_014_001_07.OriginalTransactionReference29TCH{
					Amt: &pain_014_001_07.AmountType4Choice{
						InstdAmt: &pain_014_001_07.ActiveOrHistoricCurrencyAndAmount{
							Value: 8352.45,
							Ccy:   pain_014_001_07.ActiveOrHistoricCurrencyCodeUsd,
						},
					},
					ReqdExctnDt: pain_014_001_07.DateAndDateTime2Choice{
						DtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2023-07-13T15:23:33")),
					},
					CdtrAgt: pain_014_001_07.BranchAndFinancialInstitutionIdentification6TCH{
						FinInstnId: pain_014_001_07.FinancialInstitutionIdentification18TCH{
							ClrSysMmbId: pain_014_001_07.ClearingSystemMemberIdentification2TCH{
								MmbId: "000000010",
							},
						},
					},
					Cdtr: pain_014_001_07.PartyIdentification135TCH2{
						Nm: "Participant Valid model",
						Id: &pain_014_001_07.Party38ChoiceTCH2{
							PrvtId: &pain_014_001_07.PersonIdentification13TCH{
								DtAndPlcOfBirth: pain_014_001_07.DateAndPlaceOfBirth1{
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
}

var pain014ConstantRJCT = &pain_014_001_07.DocumentTCH{
	CdtrPmtActvtnReqStsRpt: pain_014_001_07.CreditorPaymentActivationRequestStatusReportV07TCH{
		GrpHdr: pain_014_001_07.GroupHeader87TCH{
			MsgId:   "M20230713234567891T1BOTS01745252915",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:53:17"),
			InitgPty: pain_014_001_07.PartyIdentification135TCH{
				Id: pain_014_001_07.Party38ChoiceTCH{
					OrgId: &pain_014_001_07.OrganisationIdentification29TCH{
						Othr: pain_014_001_07.GenericOrganisationIdentification1TCH{
							Id: "234567891",
						},
					},
				},
			},
		},
		OrgnlGrpInfAndSts: pain_014_001_07.OriginalGroupInformation30TCH{
			OrgnlMsgId:   "M20230713200000057A1BOTS00364792520",
			OrgnlMsgNmId: pain_014_001_07.OrigMsgNamePain01300107,
			OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:53:17"),
			OrgnlNbOfTxs: "1",
		},
		OrgnlPmtInfAndSts: pain_014_001_07.OriginalPaymentInstruction31TCH{
			OrgnlPmtInfId: "20230713200000057A1BYNYE01407854635",
			TxInfAndSts: pain_014_001_07.PaymentTransaction104TCH{
				OrgnlEndToEndId: rtp.Ptr(pain_014_001_07.Max35Text("EOTS")),
				TxSts:           pain_014_001_07.ExternalPaymentTransactionStatus1CodeRjct,
				StsRsnInf: &pain_014_001_07.StatusReasonInformation12TCH{
					Rsn: pain_014_001_07.StatusReason6ChoiceTCH{
						Cd: rtp.Ptr(pain_014_001_07.ExternalStatusReason1CodeTCHAc06),
					},
				},
				OrgnlTxRef: pain_014_001_07.OriginalTransactionReference29TCH{
					Amt: &pain_014_001_07.AmountType4Choice{
						InstdAmt: &pain_014_001_07.ActiveOrHistoricCurrencyAndAmount{
							Value: 59483.54,
							Ccy:   pain_014_001_07.ActiveOrHistoricCurrencyCodeUsd,
						},
					},
					ReqdExctnDt: pain_014_001_07.DateAndDateTime2Choice{
						DtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2023-07-13T15:53:17")),
					},
					CdtrAgt: pain_014_001_07.BranchAndFinancialInstitutionIdentification6TCH{
						FinInstnId: pain_014_001_07.FinancialInstitutionIdentification18TCH{
							ClrSysMmbId: pain_014_001_07.ClearingSystemMemberIdentification2TCH{
								MmbId: "000000010",
							},
						},
					},
					Cdtr: pain_014_001_07.PartyIdentification135TCH2{
						Nm: "Participant Valid model",
						Id: &pain_014_001_07.Party38ChoiceTCH2{
							PrvtId: &pain_014_001_07.PersonIdentification13TCH{
								DtAndPlcOfBirth: pain_014_001_07.DateAndPlaceOfBirth1{
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
}

func TestReadPain014ACTC(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pain014.ACTC.xml"))
	require.NoError(t, err)

	pain014 := &messages.Message{}
	err = xml.Unmarshal(input, pain014)
	require.NoError(t, err)

	expected := messages.NewPain014Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ResponsePaymentRequest = pain014ConstantACTC
	expected.ResponsePaymentRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ResponsePaymentRequest",
	}

	assert.Equal(t, expected, pain014)
}

func TestWritePain014ACTC(t *testing.T) {
	input := messages.NewPain014Message()
	input.ResponsePaymentRequest = pain014ConstantACTC

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pain014.ACTC.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}

func TestReadPain014RJCT(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pain014.RJCT.xml"))
	require.NoError(t, err)

	pain014 := &messages.Message{}
	err = xml.Unmarshal(input, pain014)
	require.NoError(t, err)

	expected := messages.NewPain014Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.ResponsePaymentRequest = pain014ConstantRJCT
	expected.ResponsePaymentRequest.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "ResponsePaymentRequest",
	}

	assert.Equal(t, expected, pain014)
}

func TestWritePain014RJCT(t *testing.T) {
	input := messages.NewPain014Message()
	input.ResponsePaymentRequest = pain014ConstantRJCT

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pain014.RJCT.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
