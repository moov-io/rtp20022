package admi_004_001_02_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/gen/admi_004_001_02"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admi004Constant = &admi_004_001_02.Document{
	SysEvtNtfctn: admi_004_001_02.SystemEventNotificationV02{
		EvtInf: admi_004_001_02.Event2{
			EvtCd: "960",
			EvtParam: []*admi_004_001_02.Max35Text{
				rtp.Ptr(admi_004_001_02.Max35Text("202301181054430BROADCAST01104022138")),
				rtp.Ptr(admi_004_001_02.Max35Text("Broadcast")),
				rtp.Ptr(admi_004_001_02.Max35Text("123456789T1")),
				rtp.Ptr(admi_004_001_02.Max35Text("1234567890101")),
				rtp.Ptr(admi_004_001_02.Max35Text("AVAILABLE")),
			},
			EvtTm: rtp.Ptr(rtp.UnmarshalISODateTime("2023-01-18T10:54:48")),
		},
	},
}

func TestReadAdmi004(t *testing.T) {
	input, err := os.ReadFile(filepath.Join("testdata", "admi004.RTP.xml"))
	require.NoError(t, err)

	admi004 := &messages.HdrAndData{}
	err = xml.Unmarshal(input, admi004)
	require.NoError(t, err)

	expected := messages.NewAdmi004Message()
	expected.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "Message",
	}
	expected.AppHdr.CreDt = rtp.ISONormalisedDateTime(time.Date(1, time.January, 1, 0, 0, 0, 0, rtp.Eastern()))
	expected.SystemNotificationEvent = admi004Constant
	expected.SystemNotificationEvent.XMLName = xml.Name{
		Space: "urn:tch",
		Local: "SystemNotificationEvent",
	}

	assert.Equal(t, expected, admi004)
}

func TestWriteAdmi004(t *testing.T) {
	input := messages.NewAdmi004Message()
	input.SystemNotificationEvent = admi004Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admi004.RTP.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}

func TestParse_Admi004(t *testing.T) {
	doc := readSystemEventNotification(t, "admi_004-960-AVAILABLE.xml")
	require.Equal(t, "960", string(doc.EvtInf.EvtCd))
	require.Equal(t, "Broadcast", string(*doc.EvtInf.EvtParam[1]))
	require.Nil(t, doc.EvtInf.EvtDesc)
	require.Equal(t, "2023-01-18T10:54:48-05:00", time.Time(*doc.EvtInf.EvtTm).Format(time.RFC3339))

	sysEvent := readSystemEventNotification(t, "admi_004-960-UNAVAILABLE.xml")
	require.Equal(t, "960", sysEvent.Code())
	require.Equal(t, "UNAVAILABLE", sysEvent.Parameters()[4])
	require.Equal(t, "", sysEvent.Description())
	require.Equal(t, "2023-01-18T10:54:43-05:00", sysEvent.Time().Format(time.RFC3339))

	sysEvent = readSystemEventNotification(t, "admi_004-971-SUSPENDED.xml")
	require.Equal(t, "971", sysEvent.Code())
	require.Equal(t, "SUSPENDED", sysEvent.Parameters()[2])
	require.Equal(t, "", sysEvent.Description())
	require.Equal(t, "2023-01-18T10:52:21-05:00", sysEvent.Time().Format(time.RFC3339))

	sysEvent = readSystemEventNotification(t, "admi_004_971-NORMAL.xml")
	require.Equal(t, "971", sysEvent.Code())
	require.Equal(t, "NORMAL", sysEvent.Parameters()[2])
	require.Equal(t, "", sysEvent.Description())
	require.Equal(t, "2023-01-18T10:52:26-05:00", sysEvent.Time().Format(time.RFC3339))
}

func readSystemEventNotification(t *testing.T, filename string) admi_004_001_02.SystemEventNotificationV02 {
	t.Helper()

	bs, err := os.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err)

	message := messages.NewAdmi004Message()
	err = xml.Unmarshal(bs, &message)
	require.NoError(t, err)
	require.NotNil(t, message)
	return message.SystemNotificationEvent.SysEvtNtfctn
}
