package rtp

import (
	"bytes"
	"cloud.google.com/go/civil"
	"time"
)

type ISODate civil.Date

func UnmarshalISODate(text string) ISODate {
	dateTime := ISODate{}
	_ = dateTime.UnmarshalText([]byte(text))
	return dateTime
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

type ISONormalisedDateTime time.Time

func UnmarshalISONormalisedDateTime(text string) ISONormalisedDateTime {
	dateTime := ISONormalisedDateTime{}
	_ = dateTime.UnmarshalText([]byte(text))
	return dateTime
}

func MarshalISONormalisedDateTime(t ISONormalisedDateTime) string {
	txt, _ := t.MarshalText()
	return string(txt)
}

func (t *ISONormalisedDateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}

func (t ISONormalisedDateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

func (t ISONormalisedDateTime) Validate() error {
	_, err := t.MarshalText()
	return err
}

type ISOTime time.Time

func UnmarshalISOTime(text string) ISOTime {
	dateTime := ISOTime{}
	_ = dateTime.UnmarshalText([]byte(text))
	return dateTime
}

func MarshalISOTime(t ISOTime) string {
	txt, _ := t.MarshalText()
	return string(txt)
}

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}

func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

func (t ISOTime) Validate() error {
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
	d := (civil.Date)(t)
	return d.MarshalText()
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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}

func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
