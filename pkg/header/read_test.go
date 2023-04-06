package header

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponse(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("testdata", "header.xml"))
	require.NoError(t, err)

	head := NewReader()
	err = xml.Unmarshal(bs, &head)
	require.NoError(t, err)

	require.Equal(t, "990000001T1", head.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	require.Equal(t, "990000001T1", head.FromMemberID())
	require.Equal(t, "200000057A1", head.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
	require.Equal(t, "200000057A1", head.ToMemberID())

	require.Equal(t, "B20221221990000001T1HOTS01101645161", head.BizMsgIdr.Text)
	require.Equal(t, "admn.002.001.01", head.MsgDefIdr.Text)
	require.Equal(t, "2022-12-21T15:31:19", head.CreDt.Text)

	// require.Equal(t, "http://www.w3.org/2001/04/xmlenc#sha256", head.Sgntr.Signature.SignedInfo.Reference.DigestMethod.Algorithm)
	// require.Equal(t, "jq4sQQEO21TBd4I0NHIos18VJjxEvJ9FY1Cypvk6wgg=", head.Sgntr.Signature.SignedInfo.Reference.DigestValue)

	// require.Equal(t, "Ix+4AspM39k6+ohi3K/Z0lAOMxJ7U57Ex0wmSxRXbrEDulOdtakWXXwMhHwhmrJZoUxe37xfRcOfx5hKyPGjso+b7X1EGBVjUZYUspajWX66hu++lZDaTmNVIjCvdLK/z4XxEzlA+SP7nC+KWozFf1rNAZwBT5Tw/9lmH1/ytVzAUo8YQF/9mrxQw/ulF0ZiA6nUzB2QsMWFs7UAXsVdtvqHNEjpn1id6uXVURGsBAWlR5fNgOo4JBSs6fQ8432K8Fw7QDmo7xEcokh+Wc21DAdkyKQo9RnIbNFY72iUyx/xyjDrGjkjaqS0lLQjjtmP3RsUj/46mqvVkObBJob9mg==", head.Sgntr.Signature.SignatureValue)
	// require.Equal(t, "CN=Open Test Solutions,OU=OTS,O=FIS,L=Diegem,ST=Vlaams-Brabant,C=BE", head.Sgntr.Signature.KeyInfo.X509Data.X509SubjectName)
	// require.Equal(t, "CN=Open Test Solutions, OU=OTS, O=FIS, L=Diegem, ST=Vlaams-Brabant, C=BE", head.Sgntr.Signature.KeyInfo.X509Data.X509IssuerSerial.X509IssuerName)
	// require.Equal(t, "610326338160951572", head.Sgntr.Signature.KeyInfo.X509Data.X509IssuerSerial.X509SerialNumber)
}

func TestResponseDUPL(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("testdata", "header-dupl.xml"))
	require.NoError(t, err)

	head := NewReader()
	err = xml.Unmarshal(bs, &head)
	require.NoError(t, err)

	require.Equal(t, "990000001T1", head.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	require.Equal(t, "200000057A1", head.To.FIId.FinInstnId.ClrSysMmbId.MmbId)

	require.Equal(t, "B20221221990000001T1HOTS01101645161", head.BizMsgIdr.Text)
	require.Equal(t, "admn.002.001.01", head.MsgDefIdr.Text)
	require.Equal(t, "2022-12-21T15:31:19", head.CreDt.Text)
	require.Equal(t, "DUPL", head.CpyDplct.Text)

	// require.Equal(t, "http://www.w3.org/2001/04/xmlenc#sha256", head.Sgntr.Signature.SignedInfo.Reference.DigestMethod.Algorithm)
	// require.Equal(t, "jq4sQQEO21TBd4I0NHIos18VJjxEvJ9FY1Cypvk6wgg=", head.Sgntr.Signature.SignedInfo.Reference.DigestValue)

	// require.Equal(t, "Ix+4AspM39k6+ohi3K/Z0lAOMxJ7U57Ex0wmSxRXbrEDulOdtakWXXwMhHwhmrJZoUxe37xfRcOfx5hKyPGjso+b7X1EGBVjUZYUspajWX66hu++lZDaTmNVIjCvdLK/z4XxEzlA+SP7nC+KWozFf1rNAZwBT5Tw/9lmH1/ytVzAUo8YQF/9mrxQw/ulF0ZiA6nUzB2QsMWFs7UAXsVdtvqHNEjpn1id6uXVURGsBAWlR5fNgOo4JBSs6fQ8432K8Fw7QDmo7xEcokh+Wc21DAdkyKQo9RnIbNFY72iUyx/xyjDrGjkjaqS0lLQjjtmP3RsUj/46mqvVkObBJob9mg==", head.Sgntr.Signature.SignatureValue)
	// require.Equal(t, "CN=Open Test Solutions,OU=OTS,O=FIS,L=Diegem,ST=Vlaams-Brabant,C=BE", head.Sgntr.Signature.KeyInfo.X509Data.X509SubjectName)
	// require.Equal(t, "CN=Open Test Solutions, OU=OTS, O=FIS, L=Diegem, ST=Vlaams-Brabant, C=BE", head.Sgntr.Signature.KeyInfo.X509Data.X509IssuerSerial.X509IssuerName)
	// require.Equal(t, "610326338160951572", head.Sgntr.Signature.KeyInfo.X509Data.X509IssuerSerial.X509SerialNumber)
}
