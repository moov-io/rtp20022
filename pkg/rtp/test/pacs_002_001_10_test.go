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

var pacs002Constant = &pacs_002_001_10.Document{
	FIToFIPmtStsRpt: pacs_002_001_10.FIToFIPaymentStatusReportV10{
		GrpHdr: pacs_002_001_10.GroupHeader91{
			MsgId:   "M20210701000000010B1B00000008088226",
			CreDtTm: rtp.UnmarshalISODateTime("2021-07-01T10:58:02"),
		},
		OrgnlGrpInfAndSts: []*pacs_002_001_10.OriginalGroupHeader17{
			{
				OrgnlMsgId:   "M20210701000000032A1B00000008088224",
				OrgnlMsgNmId: "pacs.008.001.08",
				OrgnlCreDtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2021-07-01T10:51:00")),
				OrgnlNbOfTxs: rtp.Ptr(pacs_002_001_10.Max15NumericText("1")),
			},
		},
		TxInfAndSts: []*pacs_002_001_10.PaymentTransaction110{
			{
				OrgnlInstrId: rtp.Ptr(pacs_002_001_10.Max35Text("20210701000000032A1B000000000047075")),
				OrgnlTxId:    rtp.Ptr(pacs_002_001_10.Max35Text("20210701000000032A1B000000000047075")),
				TxSts:        rtp.Ptr(pacs_002_001_10.ExternalPaymentTransactionStatus1Code("RJCT")),
				StsRsnInf: []*pacs_002_001_10.StatusReasonInformation12{
					{
						Rsn: &pacs_002_001_10.StatusReason6Choice{
							Cd: rtp.Ptr(pacs_002_001_10.ExternalStatusReason1Code("AC03")),
						},
					},
				},
				AccptncDtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2021-07-01T10:58:02")),
				ClrSysRef:   rtp.Ptr(pacs_002_001_10.Max35Text("002")),
				InstgAgt: &pacs_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_002_001_10.ClearingSystemMemberIdentification2{
							MmbId: "000000010",
						},
					},
				},
				InstdAgt: &pacs_002_001_10.BranchAndFinancialInstitutionIdentification6{
					FinInstnId: pacs_002_001_10.FinancialInstitutionIdentification18{
						ClrSysMmbId: &pacs_002_001_10.ClearingSystemMemberIdentification2{
							MmbId: "000000032",
						},
					},
				},
				OrgnlTxRef: &pacs_002_001_10.OriginalTransactionReference28{
					IntrBkSttlmAmt: &pacs_002_001_10.ActiveOrHistoricCurrencyAndAmount{
						Value: 10.00,
						Ccy:   "USD",
					},
					IntrBkSttlmDt: rtp.Ptr(rtp.UnmarshalISODate("2021-07-01")),
				},
			},
		},
	},
}

func TestReadPacs002(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "pacs002.RTP.xml"))
	require.NoError(t, err)

	pacs002 := &messages.HdrAndData{}
	err = xml.Unmarshal(input, pacs002)
	require.NoError(t, err)

	expected := messages.NewPacs002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.MessageStatusReport = pacs002Constant
	expected.MessageStatusReport.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "MessageStatusReport",
	}

	assert.Equal(t, expected, pacs002)
}

func TestWritePacs002(t *testing.T) {
	input := messages.NewPacs002Message()
	input.MessageStatusReport = pacs002Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "pacs002.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
