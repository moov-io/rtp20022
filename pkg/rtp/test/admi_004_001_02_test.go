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

	"github.com/moov-io/rtp20022/gen/admi_004_001_02"
	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

var admi004Constant = &admi_004_001_02.DocumentTCH{
	SysEvtNtfctn: admi_004_001_02.SystemEventNotificationV02TCH{
		EvtInf: admi_004_001_02.Event2TCH{
			EvtCd: "960",
			EvtParam: []admi_004_001_02.Max35Text{
				admi_004_001_02.Max35Text("202307111045580BROADCAST01867895274"),
				admi_004_001_02.Max35Text("Broadcast"),
				admi_004_001_02.Max35Text("123456789T1"),
				admi_004_001_02.Max35Text("1234567890101"),
				admi_004_001_02.Max35Text("AVAILABLE"),
			},
			EvtTm: rtp.UnmarshalISODateTime("2023-07-11T10:46:03"),
		},
	},
}

func TestWriteAdmi004(t *testing.T) {
	input := messages.NewAdmi004Message()
	input.SystemNotificationEvent = admi004Constant

	output, err := xml.MarshalIndent(input, "", "    ")
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join("testdata", "admi004.960-AVAILABLE.xml"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), fmt.Sprintf("%s%s\n", xml.Header, string(output)))
}

func TestReadAdmi004(t *testing.T) {
	t.Run("960-AVAILABLE", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.960-AVAILABLE.xml")
		require.Equal(t, "960", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "Broadcast", string(sysEvent.EvtInf.EvtParam[1]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:46:03-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("960-UNAVAILABLE", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.960-UNAVAILABLE.xml")
		require.Equal(t, "960", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "UNAVAILABLE", string(sysEvent.EvtInf.EvtParam[4]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:45:58-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("971-SUSPENDED", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.971-SUSPENDED.xml")
		require.Equal(t, "971", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "SUSPENDED", string(sysEvent.EvtInf.EvtParam[2]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:47:01-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("971-NORMAL", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.971-NORMAL.xml")
		require.Equal(t, "971", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "NORMAL", string(sysEvent.EvtInf.EvtParam[2]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:47:06-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("972-DEFAULTBOTH", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.972-DEFAULTBOTH.xml")
		require.Equal(t, "972", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "DEFAULTBOTH", string(sysEvent.EvtInf.EvtParam[5]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:48:07-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("972-NORMAL", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.972-NORMAL.xml")
		require.Equal(t, "972", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "NORMAL", string(sysEvent.EvtInf.EvtParam[5]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:48:12-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("975-ACTIVE", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.975-ACTIVE.xml")
		require.Equal(t, "975", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "ACTIVE", string(sysEvent.EvtInf.EvtParam[3]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:49:08-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("976-ACTIVE", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.976-ACTIVE.xml")
		require.Equal(t, "976", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "ACTIVE", string(sysEvent.EvtInf.EvtParam[4]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:49:18-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("976-INACTIVE", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.976-INACTIVE.xml")
		require.Equal(t, "976", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "INACTIVE", string(sysEvent.EvtInf.EvtParam[4]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:49:08-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("981-Broadcast", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.981-Broadcast.xml")
		require.Equal(t, "981", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "Broadcast", string(sysEvent.EvtInf.EvtParam[1]))
		require.Equal(t, "Broadcast Message", string(*sysEvent.EvtInf.EvtDesc))
		require.Equal(t, "2023-07-11T10:50:38-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("981-Notification", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.981-Notification.xml")
		require.Equal(t, "981", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "Notification", string(sysEvent.EvtInf.EvtParam[1]))
		require.Equal(t, "Notification Message", string(*sysEvent.EvtInf.EvtDesc))
		require.Equal(t, "2023-07-11T10:51:24-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("982-SIGNON", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.982-SIGNON.xml")
		require.Equal(t, "982", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "SIGNON", string(sysEvent.EvtInf.EvtParam[4]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:52:04-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("982-SIGNOFF", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.982-SIGNOFF.xml")
		require.Equal(t, "982", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "SIGNOFF", string(sysEvent.EvtInf.EvtParam[4]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:52:09-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("993-EXCEEDED", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.993-EXCEEDED.xml")
		require.Equal(t, "993", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "EXCEEDED", string(sysEvent.EvtInf.EvtParam[5]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:53:05-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("993-NORMAL", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.993-NORMAL.xml")
		require.Equal(t, "993", string(sysEvent.EvtInf.EvtCd))
		require.Equal(t, "NORMAL", string(sysEvent.EvtInf.EvtParam[5]))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:53:06-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("998", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.998.xml")
		require.Equal(t, "998", string(sysEvent.EvtInf.EvtCd))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:53:06-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})

	t.Run("999", func(t *testing.T) {
		sysEvent := readSystemEventNotification(t, "admi004.999.xml")
		require.Equal(t, "999", string(sysEvent.EvtInf.EvtCd))
		require.Nil(t, sysEvent.EvtInf.EvtDesc)
		require.Equal(t, "2023-07-11T10:53:01-04:00", time.Time(sysEvent.EvtInf.EvtTm).Format(time.RFC3339))
	})
}

func readSystemEventNotification(t *testing.T, filename string) admi_004_001_02.SystemEventNotificationV02TCH {
	t.Helper()

	bs, err := os.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err)

	message := messages.NewAdmi004Message()
	err = xml.Unmarshal(bs, &message)
	require.NoError(t, err)
	require.NotNil(t, message)
	return message.SystemNotificationEvent.SysEvtNtfctn
}
