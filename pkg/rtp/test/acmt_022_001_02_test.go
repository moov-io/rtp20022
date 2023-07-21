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

	"github.com/moov-io/rtp20022/gen/acmt_022_001_02"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var acmt022Constant = &acmt_022_001_02.DocumentTCH{
	IdModAdvc: acmt_022_001_02.IdentificationModificationAdviceV02TCH{
		Assgnmt: acmt_022_001_02.IdentificationAssignment2TCH{
			MsgId:   "M20230714990000001T1HOTS01278645046",
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-14T15:02:32"),
			Assgnr: acmt_022_001_02.Party12ChoiceTCH{
				Agt: &acmt_022_001_02.BranchAndFinancialInstitutionIdentification5TCH{
					FinInstnId: acmt_022_001_02.FinancialInstitutionIdentification8TCH{
						ClrSysMmbId: acmt_022_001_02.ClearingSystemMemberIdentification2TCH{
							MmbId: "990000001",
						},
					},
				},
			},
			Assgne: acmt_022_001_02.Party12ChoiceTCH{
				Agt: &acmt_022_001_02.BranchAndFinancialInstitutionIdentification5TCH{
					FinInstnId: acmt_022_001_02.FinancialInstitutionIdentification8TCH{
						ClrSysMmbId: acmt_022_001_02.ClearingSystemMemberIdentification2TCH{
							MmbId: "234567891",
						},
					},
				},
			},
		},
		OrgnlTxRef: acmt_022_001_02.OriginalTransactionReference18TCH{
			MsgId:   "M20230714234567891T1BOTS01558775951",
			MsgNmId: acmt_022_001_02.OrigMsgNamePacs00800108,
			CreDtTm: rtp.UnmarshalISODateTime("2023-07-14T14:01:48"),
			OrgnlTx: acmt_022_001_02.PaymentIdentification4TCH{
				InstrId:    "20230714234567891T1BQPZY00718083704",
				EndToEndId: "EOTS",
				TxId:       "20230714234567891T1BQPZY00718083704",
			},
		},
		Mod: acmt_022_001_02.IdentificationModification2TCH{
			Id: "20230714990000001T1HVMBR00558027681",
			OrgnlPtyAndAcctId: &acmt_022_001_02.IdentificationInformation2TCH{
				Acct: acmt_022_001_02.AccountIdentification4Choice{
					Othr: &acmt_022_001_02.GenericAccountIdentification1{
						Id: "23234346324234234",
					},
				},
				Agt: acmt_022_001_02.BranchAndFinancialInstitutionIdentification5TCH{
					FinInstnId: acmt_022_001_02.FinancialInstitutionIdentification8TCH{
						ClrSysMmbId: acmt_022_001_02.ClearingSystemMemberIdentification2TCH{
							MmbId: "123456780",
						},
					},
				},
			},
			UpdtdPtyAndAcctId: acmt_022_001_02.IdentificationInformation2TCH{
				Acct: acmt_022_001_02.AccountIdentification4Choice{
					Othr: &acmt_022_001_02.GenericAccountIdentification1{
						Id: "00000000558027681",
					},
				},
				Agt: acmt_022_001_02.BranchAndFinancialInstitutionIdentification5TCH{
					FinInstnId: acmt_022_001_02.FinancialInstitutionIdentification8TCH{
						ClrSysMmbId: acmt_022_001_02.ClearingSystemMemberIdentification2TCH{
							MmbId: "123456780",
						},
					},
				},
			},
		},
	},
}

func TestReadAcmt022(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "acmt022.xml"))
	require.NoError(t, err)

	acmt022 := &messages.Message{}
	err = xml.Unmarshal(input, acmt022)
	require.NoError(t, err)

	expected := messages.NewAcmt022Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.TokenIdentification = acmt022Constant
	expected.TokenIdentification.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "TokenIdentification",
	}

	assert.Equal(t, expected, acmt022)
}

func TestWriteAcmt022(t *testing.T) {
	input := messages.NewAcmt022Message()
	input.TokenIdentification = acmt022Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "acmt022.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.TokenIdentification.Validate())
}
