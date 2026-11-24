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

func TestISODateTimeConvertsToEastern(t *testing.T) {
	cases := map[string]struct {
		when     time.Time
		expected string
	}{
		"UTC during EST": {
			when:     time.Date(2026, time.November, 24, 3, 30, 0, 0, time.UTC),
			expected: "2026-11-23T22:30:00", // UTC-5
		},
		"UTC during EDT": {
			when:     time.Date(2026, time.July, 4, 16, 0, 0, 0, time.UTC),
			expected: "2026-07-04T12:00:00", // UTC-4
		},
		"fixed offset": {
			when:     time.Date(2026, time.March, 2, 9, 15, 0, 0, time.FixedZone("BRT", -3*60*60)),
			expected: "2026-03-02T07:15:00",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			out := rtp.MarshalISODateTime(rtp.ISODateTime(tc.when))
			require.Equal(t, tc.expected, out)

			// reading the value back must give the same instant
			read := rtp.UnmarshalISODateTime(out)
			require.True(t, tc.when.Equal(time.Time(read)), "got %v", time.Time(read))
		})
	}

	t.Run("zero time is unchanged", func(t *testing.T) {
		require.Equal(t, "0001-01-01T00:00:00", rtp.MarshalISODateTime(rtp.ISODateTime(time.Time{})))
	})
}
