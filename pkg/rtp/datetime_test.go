package rtp_test

import (
	"cloud.google.com/go/civil"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestISODateFormat(t *testing.T) {
	when := civil.Date{
		Year:  2019,
		Month: time.March,
		Day:   21,
	}

	require.Equal(t, rtp.ISODate(when), rtp.UnmarshalISODate("2019-03-21"))

	out, err := rtp.ISODate(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21", string(out))

	out, err = xml.Marshal(rtp.ISODate(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODate>2019-03-21</ISODate>", string(out))

	var read rtp.ISODate
	err = xml.Unmarshal([]byte("<ISODate>2019-03-21</ISODate>"), &read)
	require.NoError(t, err)
	require.True(t, when == (civil.Date)(read))
}

func TestISODateTimeFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	when := time.Date(2019, time.March, 21, 10, 36, 19, 0, loc)

	require.Equal(t, rtp.ISODateTime(when), rtp.UnmarshalISODateTime("2019-03-21T10:36:19"))

	out, err := rtp.ISODateTime(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21T10:36:19", string(out))

	out, err = xml.Marshal(rtp.ISODateTime(when))
	require.NoError(t, err)
	require.Equal(t, "<ISODateTime>2019-03-21T10:36:19</ISODateTime>", string(out))

	var read rtp.ISODateTime
	err = xml.Unmarshal([]byte("<ISODateTime>2019-03-21T10:36:19</ISODateTime>"), &read)
	require.NoError(t, err)
	require.True(t, when.Equal(time.Time(read)))
}
