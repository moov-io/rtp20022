package tests

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/gen/messages"

	"github.com/stretchr/testify/require"
)

func TestPacs008(t *testing.T) {
	input := `<?xml version="1.0" encoding="UTF-8"?><Message xmlns="urn:tch" xmlns:head="urn:iso:std:iso:20022:tech:xsd:head.001.001.01" xmlns:ct="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"><AppHdr><head:Fr><head:FIId><head:FinInstnId><head:ClrSysMmbId><head:MmbId>000000010B1</head:MmbId></head:ClrSysMmbId></head:FinInstnId></head:FIId></head:Fr><head:To><head:FIId><head:FinInstnId><head:ClrSysMmbId><head:MmbId>200000058T1</head:MmbId></head:ClrSysMmbId></head:FinInstnId></head:FIId></head:To><head:BizMsgIdr>B20231206000000010B1B00335081122910</head:BizMsgIdr><head:MsgDefIdr>pacs.008.001.08</head:MsgDefIdr><head:CreDt>2023-12-06T23:08:38</head:CreDt><head:Sgntr><ds:Signature xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:SignedInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:CanonicalizationMethod Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/><ds:SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"/><ds:Reference URI=""><ds:Transforms><ds:Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/><ds:Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/></ds:Transforms><ds:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"/><ds:DigestValue>FZ5WNW1KttP2ZoEkvu5YliFKRxzuwxAPSTeAafm4HCA=</ds:DigestValue></ds:Reference></ds:SignedInfo><ds:SignatureValue>kqN8hzwC0KmMKVVXbwOWWxdLBhn3gb0jA0a9wpQz85CcgPE0hR1zA5A7gv4WWvFooqnMLotzLQ/tzaGKjbbPm1Bf08XLc+BcWVzXjCfRSMNtpR+gMp50ODGlvNafrMHnsummGkFWF3+HGivCFpkLWSzau2+20DKsCxAQGGdhWCMfDqcYWfRVrkVWKVfz+YDUF7K7U7hXVM3DuWhJMcxmcIwGern/14KX+fIecMxsE2+a8R+sOPr5gYQ+TOwwFo/bwGCHTT5MLSgOpa+1+fLYnxY0QVYXihIGMYuy+JyW3JS+1FkzRU4X3oWayfe8FEVnkTTHZ0kTx/pFuOdLY5pZvg==</ds:SignatureValue><ds:KeyInfo><ds:X509Data><ds:X509IssuerSerial><ds:X509IssuerName>CN=DigiCert Global G2 TLS RSA SHA256 2020 CA1,O=DigiCert Inc,C=US</ds:X509IssuerName><ds:X509SerialNumber>7733185569732290893228546868636904811</ds:X509SerialNumber></ds:X509IssuerSerial><ds:X509SubjectName>CN=bc553361-9e6a-468f-9941-e7acb234513d.moov.io,O=Moov Financial Inc.,L=Cedar Falls,ST=Iowa,C=US</ds:X509SubjectName></ds:X509Data></ds:KeyInfo></ds:Signature></head:Sgntr></AppHdr><CreditTransfer><ct:FIToFICstmrCdtTrf><ct:GrpHdr><ct:MsgId>M20231206000000010B1B00335081122910</ct:MsgId><ct:CreDtTm>2023-12-06T23:08:38</ct:CreDtTm><ct:NbOfTxs>1</ct:NbOfTxs><ct:TtlIntrBkSttlmAmt Ccy="USD">526.33</ct:TtlIntrBkSttlmAmt><ct:IntrBkSttlmDt>2023-12-06</ct:IntrBkSttlmDt><ct:SttlmInf><ct:SttlmMtd>CLRG</ct:SttlmMtd><ct:ClrSys><ct:Cd>TCH</ct:Cd></ct:ClrSys></ct:SttlmInf></ct:GrpHdr><ct:CdtTrfTxInf><ct:PmtId><ct:InstrId>20231206200000057T1B123840081122910</ct:InstrId><ct:EndToEndId>NOREF</ct:EndToEndId><ct:TxId>20231206200000057T1B123840081122910</ct:TxId></ct:PmtId><ct:PmtTpInf><ct:SvcLvl><ct:Cd>SDVA</ct:Cd></ct:SvcLvl><ct:LclInstrm><ct:Prtry>INTERMEDIARY</ct:Prtry></ct:LclInstrm><ct:CtgyPurp><ct:Prtry>CONSUMER</ct:Prtry></ct:CtgyPurp></ct:PmtTpInf><ct:IntrBkSttlmAmt Ccy="USD">526.33</ct:IntrBkSttlmAmt><ct:ChrgBr>SLEV</ct:ChrgBr><ct:InstgAgt><ct:FinInstnId><ct:ClrSysMmbId><ct:MmbId>123456780</ct:MmbId></ct:ClrSysMmbId></ct:FinInstnId></ct:InstgAgt><ct:InstdAgt><ct:FinInstnId><ct:ClrSysMmbId><ct:MmbId>987654320</ct:MmbId></ct:ClrSysMmbId></ct:FinInstnId></ct:InstdAgt><ct:Dbtr><ct:Nm>Moov, Inc</ct:Nm><ct:PstlAdr><ct:StrtNm/><ct:PstCd/><ct:TwnNm/><ct:CtrySubDvsn/><ct:Ctry/></ct:PstlAdr></ct:Dbtr><ct:DbtrAcct><ct:Id><ct:Othr><ct:Id>65434577456456</ct:Id></ct:Othr></ct:Id></ct:DbtrAcct><ct:DbtrAgt><ct:FinInstnId><ct:ClrSysMmbId><ct:MmbId>123456780</ct:MmbId></ct:ClrSysMmbId></ct:FinInstnId></ct:DbtrAgt><ct:CdtrAgt><ct:FinInstnId><ct:ClrSysMmbId><ct:MmbId>987654320</ct:MmbId></ct:ClrSysMmbId></ct:FinInstnId></ct:CdtrAgt><ct:Cdtr><ct:Nm>Jane Doe</ct:Nm><ct:PstlAdr><ct:StrtNm>842 Other Ave</ct:StrtNm><ct:PstCd>89415</ct:PstCd><ct:TwnNm>Green Stream</ct:TwnNm><ct:CtrySubDvsn>IA</ct:CtrySubDvsn><ct:Ctry>US</ct:Ctry></ct:PstlAdr></ct:Cdtr><ct:CdtrAcct><ct:Id><ct:Othr><ct:Id>74625332505100907</ct:Id></ct:Othr></ct:Id></ct:CdtrAcct></ct:CdtTrfTxInf></ct:FIToFICstmrCdtTrf></CreditTransfer></Message>`

	t.Run("parse and validate, while expecting errors", func(t *testing.T) {
		var m messages.Message
		err := xml.NewDecoder(strings.NewReader(input)).Decode(&m)
		require.NoError(t, err)

		err = m.Validate()
		require.Error(t, err)

		el, ok := err.(base.ErrorList)
		require.True(t, ok)
		require.Len(t, el, 5)

		// Check each error
		require.ErrorContains(t, el[0], "PostalAddress24TCH.StrtNm:  fails validation with length 0 >= required minLength 1")
		require.ErrorContains(t, el[1], "PostalAddress24TCH.PstCd:  fails validation with length 0 >= required minLength 1")
		require.ErrorContains(t, el[2], "PostalAddress24TCH.TwnNm:  fails validation with length 0 >= required minLength 1")
		require.ErrorContains(t, el[3], "PostalAddress24TCH.CtrySubDvsn:  fails validation with length 0 >= required minLength 1")
		require.ErrorContains(t, el[4], "PostalAddress24TCH.Ctry:  fails validation with pattern [A-Z]{2,2}")
	})
}
