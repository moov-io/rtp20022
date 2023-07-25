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
	"github.com/moov-io/rtp20022/gen/remt_001_001_04"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var remt001Constant = &remt_001_001_04.DocumentTCH{
	RmtAdvc: remt_001_001_04.RemittanceAdviceV04TCH{
		GrpHdr: remt_001_001_04.GroupHeader79TCH{
			MsgId:   "M20230713234567891T1BOTS01522107521",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-13T15:29:36"),
			InitgPty: remt_001_001_04.PartyIdentification135TCH{
				Id: remt_001_001_04.Party38ChoiceTCH{
					OrgId: &remt_001_001_04.OrganisationIdentification29TCH{
						Othr: remt_001_001_04.GenericOrganisationIdentification1TCH{
							Id: "234567891",
						},
					},
				},
			},
			MsgRcpt: remt_001_001_04.PartyIdentification135TCH{
				Id: remt_001_001_04.Party38ChoiceTCH{
					OrgId: &remt_001_001_04.OrganisationIdentification29TCH{
						Othr: remt_001_001_04.GenericOrganisationIdentification1TCH{
							Id: "000000010",
						},
					},
				},
			},
		},
		RmtInf: remt_001_001_04.RemittanceInformation19TCH{
			RmtId: "20230713152936OTS00752775034",
			OrgnlPmtInf: remt_001_001_04.OriginalPaymentInformation8TCH{
				Refs: remt_001_001_04.TransactionReferences5{
					InstrId:    rtp.Ptr(remt_001_001_04.Max35Text("20230713234567891T1BCXIA01868416991")),
					EndToEndId: "EOTS",
					TxId:       rtp.Ptr(remt_001_001_04.Max35Text("20230713234567891T1BCXIA01868416991")),
				},
				Amt: remt_001_001_04.AmountType3Choice{
					InstdAmt: &remt_001_001_04.ActiveOrHistoricCurrencyAndAmount{
						Value: 46154.49,
						Ccy:   remt_001_001_04.ActiveOrHistoricCurrencyCodeUsd,
					},
				},
				Dbtr: &remt_001_001_04.PartyIdentification135TCH3{
					Nm: rtp.Ptr(remt_001_001_04.Max140Text("MRS. GREEN")),
				},
				DbtrAcct: &remt_001_001_04.CashAccount38{
					Id: remt_001_001_04.AccountIdentification4Choice{
						Othr: &remt_001_001_04.GenericAccountIdentification1{
							Id: "US88664715164441",
						},
					},
				},
				DbtrAgt: &remt_001_001_04.BranchAndFinancialInstitutionIdentification6TCH{
					FinInstnId: remt_001_001_04.FinancialInstitutionIdentification18TCH{
						ClrSysMmbId: remt_001_001_04.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
				Cdtr: &remt_001_001_04.PartyIdentification135TCH3{
					Nm: rtp.Ptr(remt_001_001_04.Max140Text("Participant Valid model")),
				},
			},
		},
	},
}

func TestReadRemt001(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "remt001.xml"))
	require.NoError(t, err)

	remt001 := &messages.Message{}
	err = xml.Unmarshal(input, remt001)
	require.NoError(t, err)

	expected := messages.NewRemt001Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISODateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.StandaloneRemittance = remt001Constant
	expected.StandaloneRemittance.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "StandaloneRemittance",
	}

	assert.Equal(t, expected, remt001)
}

func TestWriteRemt001(t *testing.T) {
	input := messages.NewRemt001Message()
	input.StandaloneRemittance = remt001Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "remt001.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.StandaloneRemittance.Validate())
}
