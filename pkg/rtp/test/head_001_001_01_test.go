package test

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/gen/head_001_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/gen/xmldsig"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var head001Constant = head_001_001_01.BusinessApplicationHeaderV01{
	Fr: head_001_001_01.Party9ChoiceBAH{
		FIId: &head_001_001_01.BranchAndFinancialInstitutionIdentification5BAH{
			FinInstnId: head_001_001_01.FinancialInstitutionIdentification8BAH{
				ClrSysMmbId: head_001_001_01.ClearingSystemMemberIdentification2ADMN{
					MmbId: "990000001T1",
				},
			},
		},
	},
	To: head_001_001_01.Party9ChoiceBAH{
		FIId: &head_001_001_01.BranchAndFinancialInstitutionIdentification5BAH{
			FinInstnId: head_001_001_01.FinancialInstitutionIdentification8BAH{
				ClrSysMmbId: head_001_001_01.ClearingSystemMemberIdentification2ADMN{
					MmbId: "200000057A1",
				},
			},
		},
	},
	BizMsgIdr: "B20221221990000001T1HOTS01101645161",
	MsgDefIdr: head_001_001_01.OrigMsgNameAdmn00200101,
	CreDt:     rtp.UnmarshalISONormalisedDateTime("2022-12-21T15:31:19"),
	CpyDplct:  rtp.Ptr(head_001_001_01.CopyDuplicate1CodeDupl),
}
var head001Signature = &head_001_001_01.Sgntr{
	Signature: &xmldsig.Signature{
		XMLName: xml.Name{
			Space: "http://www.w3.org/2000/09/xmldsig#",
			Local: "Signature",
		},
		SignedInfo: xmldsig.SignedInfoType{
			CanonicalizationMethod: xmldsig.CanonicalizationMethodType{
				Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315",
			},
			SignatureMethod: xmldsig.SignatureMethodType{
				Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
			},
			Reference: []xmldsig.ReferenceType{
				{
					URI: rtp.Ptr(""),
					Transforms: &xmldsig.TransformsType{
						Transform: []xmldsig.TransformType{
							{
								Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
							},
							{
								Algorithm: "http://www.w3.org/2006/12/xml-c14n11",
							},
						},
					},
					DigestMethod: xmldsig.DigestMethodType{
						Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
					},
					DigestValue: "1WvBpJiQuDdiXBqQLqmysAw1W0b1E0vlzCqRReoCnxI=",
				},
			},
		},
		SignatureValue: xmldsig.SignatureValueType{
			Data: "GFkk2bhGHd7kbsRBa+s8HvADwsUTm6N2Mro06O2yCJxl7SkXp/l5OY8QBYw/NlRjBHQsQzpymoNz5+p7OPTu/YGZ5s29iYD8gCA55rbQjr+K+m4dzVAcZQoLs7vSe3QoqQlUugkzFlMwjAGLW0WHZQA9PvJFbVAMVbZk+mMIJH8pyP1kSd+iu3sccav5d9qkG/ESBXUzCzQwrg9OTwBZXop5bs2SIAlG90III0J31l6ARq4eUtzbHaDGKDlbcHkkwnKdev/uYkvvjx1yr/XzMuVvpwbbOfhmcgTlAxKsKpx7sM1Pxq5Jx66Vt7k9AzKX0W/PSPDdRJjQfolEW1Aptg==",
		},
		KeyInfo: &xmldsig.KeyInfoType{
			X509Data: []*xmldsig.X509DataType{
				{
					X509SubjectName: rtp.Ptr("CN=demo.tchbt.tchrtp.org,OU=Technology Integration Management,O=The Clearing House Payments Company LLC,L=New York,S=New York,C=US"),
					X509IssuerSerial: &xmldsig.X509IssuerSerialType{
						X509IssuerName:   "CN=The Clearing House Enterprise Issuing CA 1, O=The Clearing House Payments Company, C=US",
						X509SerialNumber: "303111756915795778601924",
					},
				},
			},
		},
	},
}

func TestReadHead001_signature(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "head001.RTP.signature.xml"))
	require.NoError(t, err)

	head001 := &messages.Message{}
	err = xml.Unmarshal(input, head001)
	require.NoError(t, err)

	expected := messages.NewPacs008Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr = head001Constant
	expected.AppHdr.Sgntr = head001Signature

	assert.Equal(t, expected, head001)
}

func TestReadHead001_nosignature(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "head001.RTP.nosignature.xml"))
	require.NoError(t, err)

	head001 := &messages.Message{}
	err = xml.Unmarshal(input, head001)
	require.NoError(t, err)

	expected := messages.NewPacs008Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr = head001Constant

	assert.Equal(t, expected, head001)
}

func TestWriteHead001_signature(t *testing.T) {
	input := messages.NewPacs008Message()
	input.AppHdr = head001Constant
	input.AppHdr.Sgntr = head001Signature

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "head001.RTP.signature.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}

func TestWriteHead001_nosignature(t *testing.T) {
	input := messages.NewPacs008Message()
	input.AppHdr = head001Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "head001.RTP.nosignature.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}
