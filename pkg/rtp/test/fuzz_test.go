// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package test

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func FuzzRTPMessage(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}
		data := []byte(contents)

		var msg messages.Message
		if err := xml.Unmarshal(data, &msg); err != nil {
			return
		}
		_, _ = xml.Marshal(&msg)
	})
}

func FuzzRTPTypes(f *testing.F) {
	f.Add("2019-03-21")
	f.Add("2019-03-21T10:36:19")
	f.Add("0")
	f.Add("1234.56")
	f.Add("")
	f.Add("not-a-date")
	f.Add("20230713234567891T1BOTS02083825666")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 256 {
			t.Skip()
		}

		d := rtp.UnmarshalISODate(s)
		_ = d.Validate()
		_, _ = d.MarshalText()

		dt := rtp.UnmarshalISODateTime(s)
		_ = dt.Validate()
		_, _ = dt.MarshalText()

		_ = rtp.ValidatePattern(s, `^[A-Za-z0-9]+$`)
		_ = rtp.ValidateLength(s, 35)
		_ = rtp.ValidateMinLength(s, 1)
		_ = rtp.ValidateMaxLength(s, 35)
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("<Message></Message>")

	_ = filepath.Walk("testdata", func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".xml") {
			bs, err := os.ReadFile(path)
			if err != nil || len(bs) > 512*1024 {
				return nil
			}
			f.Add(string(bs))
		}
		return nil
	})
}
