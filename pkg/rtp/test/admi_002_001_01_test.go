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

	"github.com/moov-io/rtp20022/gen/admi_002_001_01"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admi002Constant = &admi_002_001_01.DocumentTCH{
	Admi00200101: admi_002_001_01.MessageRejectV01TCH{
		RltdRef: admi_002_001_01.MessageReferenceTCH{
			Ref: "20230713145322200000057A11712044729",
		},
		Rsn: admi_002_001_01.RejectionReason2TCH{
			RjctgPtyRsn: admi_002_001_01.Max35Text("690"),
			AddtlData: &rtp.Cdata{
				CDataString: `<?xml version="1.0" encoding="UTF-8"?><Message xmlns:ps="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10" xmlns:head="urn:iso:std:iso:20022:tech:xsd:head.001.001.01" p3:schemaLocation="urn:tch messages.xsd" xmlns:p3="http://www.w3.org/2001/XMLSchema-instance" xmlns="urn:tch"><AppHdr><head:Fr><head:FIId><head:FinInstnId><head:ClrSysMmbId><head:MmbId>990000001T1</head:MmbId></head:ClrSysMmbId></head:FinInstnId></head:FIId></head:Fr><head:To><head:FIId><head:FinInstnId><head:ClrSysMmbId><head:MmbId>200000057A1</head:MmbId></head:ClrSysMmbId></head:FinInstnId></head:FIId></head:To><head:BizMsgIdr>B20230713990000001T1HOTS00382937980</head:BizMsgIdr><head:MsgDefIdr>pacs.002.001.10</head:MsgDefIdr><head:CreDt>2023-07-13T14:53:21</head:CreDt><head:Sgntr><ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:SignedInfo><ds:CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315" /><ds:SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256" /><ds:Reference URI=""><ds:Transforms><ds:Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature" /><ds:Transform Algorithm="http://www.w3.org/2006/12/xml-c14n11" /></ds:Transforms><ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256" /><ds:DigestValue>PHJD8evHEOn6tDmy4Jr3XJ57/lvcWA/J+2UZ5B7WfAk=</ds:DigestValue></ds:Reference></ds:SignedInfo><ds:SignatureValue>Uae3rQJBHAgDODB3DkW41lfX7aSQdFI5NXQGV0dfISqakSZb7JWmPOt0TIUCFFIwEM6WucxxFQe1S5eMRtSH69qkz/7c/58t81ilqEXVN6WoTM8y6FBVZ2vQjdwohclqslyOtoo3kOmiEd+7WOtuF2q2XYz397COOvlI7ry0V4Dz3HLIg9NmKwGOcRMOSLRd2/RK1bF4hcy1xRTPfk3hB4lRRa0y51AC6ByRVF7A4iqAqppdaK/Tgb41x3HNHK8zg4aEEpKZ5Gau8znT8LD4H2tWKDY870/bgJ4cEW5FpwbZugU5rs+uJ+25RJTgdMI1CUXuL7Z/GOMfeM/aaIh7Ow==</ds:SignatureValue><ds:KeyInfo><ds:X509Data><ds:X509SubjectName>CN=Open Test Solutions,OU=OTS,O=FIS,L=Diegem,ST=Vlaams-Brabant,C=BE</ds:X509SubjectName><ds:X509IssuerSerial><ds:X509IssuerName>CN=Open Test Solutions, OU=OTS, O=FIS, L=Diegem, ST=Vlaams-Brabant, C=BE</ds:X509IssuerName><ds:X509SerialNumber>610326338160951572</ds:X509SerialNumber></ds:X509IssuerSerial></ds:X509Data></ds:KeyInfo></ds:Signature></head:Sgntr></AppHdr><MessageStatusReport><ps:FIToFIPmtStsRpt><ps:GrpHdr><ps:MsgId>M20230713BOTS01498664740</ps:MsgId><ps:CreDtTm>2023-07-13T14:53:21</ps:CreDtTm></ps:GrpHdr><ps:OrgnlGrpInfAndSts><ps:OrgnlMsgId>M20230713200000057A1B12813700447298</ps:OrgnlMsgId><ps:OrgnlMsgNmId>camt.029.001.09</ps:OrgnlMsgNmId><ps:OrgnlCreDtTm>2023-07-13T14:53:20</ps:OrgnlCreDtTm><ps:OrgnlNbOfTxs>1</ps:OrgnlNbOfTxs></ps:OrgnlGrpInfAndSts><ps:TxInfAndSts><ps:OrgnlInstrId>M20230713200000057A1B12813700447298</ps:OrgnlInstrId><ps:TxSts>RCVD</ps:TxSts><ps:AccptncDtTm>2023-07-13T14:53:21</ps:AccptncDtTm><ps:InstgAgt><ps:FinInstnId><ps:ClrSysMmbId><ps:MmbId>123456789</ps:MmbId></ps:ClrSysMmbId></ps:FinInstnId></ps:InstgAgt><ps:InstdAgt><ps:FinInstnId><ps:ClrSysMmbId><ps:MmbId>234567891</ps:MmbId></ps:ClrSysMmbId></ps:FinInstnId></ps:InstdAgt></ps:TxInfAndSts></ps:FIToFIPmtStsRpt></MessageStatusReport></Message>`,
			},
		},
	},
}

func TestReadAdmi002(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admi002.xml"))
	require.NoError(t, err)

	admi002 := &messages.Message{}
	err = xml.Unmarshal(input, admi002)
	require.NoError(t, err)

	expected := messages.NewAdmi002Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.MessageReject = admi002Constant
	expected.MessageReject.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "MessageReject",
	}

	assert.Equal(t, expected, admi002)
}

func TestWriteAdmi002(t *testing.T) {
	input := messages.NewAdmi002Message()
	input.MessageReject = admi002Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admi002.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
	assert.NoError(t, input.MessageReject.Validate())
}
