package rtp

import (
	"bytes"
	"cloud.google.com/go/civil"
	"time"
)

type ISODate civil.Date

func UnmarshalISODate(text string) ISODate {
	date := ISODate{}
	_ = date.UnmarshalText([]byte(text))
	return date
}

func MarshalISODate(t ISODate) string {
	txt, _ := t.MarshalText()
	return string(txt)
}

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}

func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

func (t ISODate) Validate() error {
	_, err := t.MarshalText()
	return err
}

type ISODateTime time.Time

func UnmarshalISODateTime(text string) ISODateTime {
	dateTime := ISODateTime{}
	_ = dateTime.UnmarshalText([]byte(text))
	return dateTime
}

func MarshalISODateTime(t ISODateTime) string {
	txt, _ := t.MarshalText()
	return string(txt)
}

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}

func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

func (t ISODateTime) Validate() error {
	_, err := t.MarshalText()
	return err
}

type xsdDate civil.Date

func (t *xsdDate) UnmarshalText(text []byte) error {
	s := string(bytes.TrimSpace(text))
	d := (*civil.Date)(t)
	var err error
	*d, err = civil.ParseDate(s)
	if err != nil {
		return err
	}
	return nil
}

func (t xsdDate) MarshalText() ([]byte, error) {
	return (civil.Date)(t).MarshalText()
}

func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.ParseInLocation(format, s, Eastern())
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.ParseInLocation(format+"Z07:00", s, Eastern())
	}
	return err
}

func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format)), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05")
}

func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05")
}
