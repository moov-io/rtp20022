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
	"github.com/moov-io/rtp20022/gen/pacs_009_001_08"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var pacs009Constant = &pacs_009_001_08.DocumentTCH{
	FICdtTrf: pacs_009_001_08.FinancialInstitutionCreditTransferV08TCH{
		GrpHdr: pacs_009_001_08.GroupHeader93TCH{
			MsgId:   "M20230713234567891T1BOTS02083825666",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:29:36"),
			NbOfTxs: "1",
			TtlIntrBkSttlmAmt: pacs_009_001_08.ActiveCurrencyAndAmount{
				Value: 46154.49,
				Ccy:   pacs_009_001_08.ActiveCurrencyCodeUsd,
			},
			IntrBkSttlmDt: rtp.UnmarshalISODate("2023-07-13"),
			SttlmInf: pacs_009_001_08.SettlementInstruction7TCH{
				SttlmMtd: pacs_009_001_08.SettlementMethod1CodeClrg,
				ClrSys: pacs_009_001_08.ClearingSystemIdentification3Choice{
					Cd: rtp.Ptr(pacs_009_001_08.ExternalCashClearingSystem1CodeTch),
				},
			},
		},
		CdtTrfTxInf: pacs_009_001_08.CreditTransferTransaction36TCH{
			PmtId: pacs_009_001_08.PaymentIdentification7TCH{
				InstrId:    pacs_009_001_08.Max35TextTCH2("20230713234567891T1BCXIA01868416991"),
				EndToEndId: "EOTS",
				TxId:       pacs_009_001_08.Max35Text("20230713234567891T1BCXIA01868416991"),
				ClrSysRef:  rtp.Ptr(pacs_009_001_08.Max35Text("001")),
			},
			PmtTpInf: pacs_009_001_08.PaymentTypeInformation28TCH{
				SvcLvl: pacs_009_001_08.ServiceLevel8Choice{
					Cd: rtp.Ptr(pacs_009_001_08.ExternalServiceLevel1CodeSdva),
				},
			},
			IntrBkSttlmAmt: pacs_009_001_08.ActiveCurrencyAndAmount{
				Value: 46154.49,
				Ccy:   pacs_009_001_08.ActiveCurrencyCodeUsd,
			},
			InstgAgt: pacs_009_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_009_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_009_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "234567891",
					},
				},
			},
			InstdAgt: pacs_009_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_009_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_009_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "000000010",
					},
				},
			},
			Dbtr: pacs_009_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_009_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_009_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "234567891",
					},
				},
			},
			DbtrAcct: &pacs_009_001_08.CashAccount38{
				Id: pacs_009_001_08.AccountIdentification4Choice{
					Othr: &pacs_009_001_08.GenericAccountIdentification1{
						Id: pacs_009_001_08.Max34Text("US88664715164441"),
					},
				},
			},
			Cdtr: pacs_009_001_08.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_009_001_08.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_009_001_08.ClearingSystemMemberIdentification2TCH{
						MmbId: "000000010",
					},
				},
			},
			CdtrAcct: pacs_009_001_08.CashAccount38{
				Id: pacs_009_001_08.AccountIdentification4Choice{
					Othr: &pacs_009_001_08.GenericAccountIdentification1{
						Id: pacs_009_001_08.Max34Text("42341169728762787"),
					},
				},
			},
		},
	},
}

func TestReadPacs009(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pacs009.xml"))
	require.NoError(t, err)

	pacs009 := &messages.Message{}
	err = xml.Unmarshal(input, pacs009)
	require.NoError(t, err)

	expected := messages.NewPacs009Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.FICreditTransfer = pacs009Constant
	expected.FICreditTransfer.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "FICreditTransfer",
	}

	assert.Equal(t, expected, pacs009)
}

func TestWritePacs009(t *testing.T) {
	input := messages.NewPacs009Message()
	input.FICreditTransfer = pacs009Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pacs009.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.FICreditTransfer.Validate())
}
