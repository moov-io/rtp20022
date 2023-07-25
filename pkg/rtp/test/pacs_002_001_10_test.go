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
	"github.com/moov-io/rtp20022/gen/pacs_002_001_10"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var pacs002ConstantACTC = &pacs_002_001_10.DocumentTCH{
	FIToFIPmtStsRpt: pacs_002_001_10.FIToFIPaymentStatusReportV10TCH{
		GrpHdr: pacs_002_001_10.GroupHeader91TCH{
			MsgId:   "M20230711BOTS00162724540",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:53:03"),
		},
		OrgnlGrpInfAndSts: pacs_002_001_10.OriginalGroupHeader17TCH{
			OrgnlMsgId:   "M20230711200000057A1BFFF64928178098",
			OrgnlMsgNmId: pacs_002_001_10.OrigMsgNamePacs00800108,
			OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:53:02"),
			OrgnlNbOfTxs: pacs_002_001_10.Max1NumericText("1"),
		},
		TxInfAndSts: pacs_002_001_10.PaymentTransaction110TCH{
			OrgnlInstrId: pacs_002_001_10.Max35Text("20230711200000057A1B123864928178098"),
			OrgnlTxId:    rtp.Ptr(pacs_002_001_10.Max35Text("20230711200000057A1B123864928178098")),
			TxSts:        pacs_002_001_10.ExternalPaymentTransactionStatus1CodeActc,
			AccptncDtTm:  rtp.UnmarshalISODateTime("2023-07-11T10:53:03"),
			ClrSysRef:    rtp.Ptr(pacs_002_001_10.Max35Text("001")),
			InstgAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2TCH{
						MmbId: "987654320",
					},
				},
			},
			InstdAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2TCH{
						MmbId: "123456780",
					},
				},
			},
			OrgnlTxRef: &pacs_002_001_10.OriginalTransactionReference28{
				IntrBkSttlmAmt: &pacs_002_001_10.ActiveOrHistoricCurrencyAndAmount{
					Value: 2000.00,
					Ccy:   pacs_002_001_10.ActiveOrHistoricCurrencyCodeUsd,
				},
				IntrBkSttlmDt: rtp.Ptr(rtp.UnmarshalISODate("2023-07-11")),
			},
		},
	},
}

var pacs002ConstantRJCT = &pacs_002_001_10.DocumentTCH{
	FIToFIPmtStsRpt: pacs_002_001_10.FIToFIPaymentStatusReportV10TCH{
		GrpHdr: pacs_002_001_10.GroupHeader91TCH{
			MsgId:   "M20230711BOTS00714217617",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:53:06"),
		},
		OrgnlGrpInfAndSts: pacs_002_001_10.OriginalGroupHeader17TCH{
			OrgnlMsgId:   "M20230711200000057A1BFFF88058178098",
			OrgnlMsgNmId: pacs_002_001_10.OrigMsgNamePacs00800108,
			OrgnlCreDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:53:05"),
			OrgnlNbOfTxs: pacs_002_001_10.Max1NumericText("1"),
		},
		TxInfAndSts: pacs_002_001_10.PaymentTransaction110TCH{
			OrgnlInstrId: pacs_002_001_10.Max35Text("20230711200000057A1B123888058178098"),
			OrgnlTxId:    rtp.Ptr(pacs_002_001_10.Max35Text("20230711200000057A1B123888058178098")),
			TxSts:        pacs_002_001_10.ExternalPaymentTransactionStatus1CodeRjct,
			StsRsnInf: &pacs_002_001_10.StatusReasonInformation12TCH{
				Rsn: &pacs_002_001_10.StatusReason6ChoiceTCH{
					Cd: rtp.Ptr(pacs_002_001_10.ExternalStatusReason1CodeAc03),
				},
			},
			AccptncDtTm: rtp.UnmarshalISODateTime("2023-07-11T10:53:06"),
			InstgAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2TCH{
						MmbId: "987654320",
					},
				},
			},
			InstdAgt: pacs_002_001_10.BranchAndFinancialInstitutionIdentification6TCH{
				FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18TCH{
					ClrSysMmbId: pacs_002_001_10.ClearingSystemMemberIdentification2TCH{
						MmbId: "123456780",
					},
				},
			},
			OrgnlTxRef: &pacs_002_001_10.OriginalTransactionReference28{
				IntrBkSttlmAmt: &pacs_002_001_10.ActiveOrHistoricCurrencyAndAmount{
					Value: 10000.00,
					Ccy:   pacs_002_001_10.ActiveOrHistoricCurrencyCodeUsd,
				},
				IntrBkSttlmDt: rtp.Ptr(rtp.UnmarshalISODate("2023-07-11")),
			},
		},
	},
}

func TestReadPacs002ACTC(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pacs002.ACTC.xml"))
	require.NoError(t, err)

	pacs002 := &messages.Message{}
	err = xml.Unmarshal(input, pacs002)
	require.NoError(t, err)

	expected := messages.NewPacs002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.MessageStatusReport = pacs002ConstantACTC
	expected.MessageStatusReport.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "MessageStatusReport",
	}

	assert.Equal(t, expected, pacs002)
}

func TestWritePacs002ACTC(t *testing.T) {
	input := messages.NewPacs002Message()
	input.MessageStatusReport = pacs002ConstantACTC

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pacs002.ACTC.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.MessageStatusReport.Validate())
}

func TestReadPacs002RJCT(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pacs002.RJCT.xml"))
	require.NoError(t, err)

	pacs002 := &messages.Message{}
	err = xml.Unmarshal(input, pacs002)
	require.NoError(t, err)

	expected := messages.NewPacs002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.MessageStatusReport = pacs002ConstantRJCT
	expected.MessageStatusReport.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "MessageStatusReport",
	}

	assert.Equal(t, expected, pacs002)
}

func TestWritePacs002RJCT(t *testing.T) {
	input := messages.NewPacs002Message()
	input.MessageStatusReport = pacs002ConstantRJCT

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pacs002.RJCT.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.MessageStatusReport.Validate())
}
