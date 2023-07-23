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

func TestISOTimeFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	when := time.Date(0, time.January, 1, 10, 36, 19, 0, loc)

	require.Equal(t, rtp.ISOTime(when), rtp.UnmarshalISOTime("10:36:19"))

	out, err := rtp.ISOTime(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "10:36:19", string(out))

	out, err = xml.Marshal(rtp.ISOTime(when))
	require.NoError(t, err)
	require.Equal(t, "<ISOTime>10:36:19</ISOTime>", string(out))

	var read rtp.ISOTime
	err = xml.Unmarshal([]byte("<ISOTime>10:36:19</ISOTime>"), &read)
	require.NoError(t, err)
	require.True(t, when.Equal(time.Time(read)))
}

func TestISONormalisedDateTimeFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	when := time.Date(2019, time.March, 21, 10, 36, 19, 0, loc)

	require.Equal(t, rtp.ISONormalisedDateTime(when), rtp.UnmarshalISONormalisedDateTime("2019-03-21T10:36:19"))

	out, err := rtp.ISONormalisedDateTime(when).MarshalText()
	require.NoError(t, err)
	require.Equal(t, "2019-03-21T10:36:19", string(out))

	out, err = xml.Marshal(rtp.ISONormalisedDateTime(when))
	require.NoError(t, err)
	require.Equal(t, "<ISONormalisedDateTime>2019-03-21T10:36:19</ISONormalisedDateTime>", string(out))

	var read rtp.ISONormalisedDateTime
	err = xml.Unmarshal([]byte("<ISONormalisedDateTime>2019-03-21T10:36:19</ISONormalisedDateTime>"), &read)
	require.NoError(t, err)
	require.True(t, when.Equal(time.Time(read)))
}
