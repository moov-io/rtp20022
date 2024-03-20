package rtp

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"math/big"
	"time"
)

type SerialGenerator func(length int) string

// AlphaSerialNumber will generate an alphanumeric serial number from the current timestamp
func AlphaSerialNumber(length int) string {
	bs := []byte(fmt.Sprint(time.Now().UnixMilli()))
	ss := base32.StdEncoding.EncodeToString(bs)
	var serialNumber string
	for i := len(ss) - 1; i >= 0; i-- {
		if (ss[i] >= '0' && ss[i] <= '9') || (ss[i] >= 'A' && ss[i] <= 'Z') {
			serialNumber += string(ss[i])
		}
	}
	if len(serialNumber) > length {
		serialNumber = serialNumber[:length]
	}
	return serialNumber
}

// NumericSerialNumber will generate an numeric serial number from the current timestamp
func NumericSerialNumber(length int) string {
	ss := fmt.Sprint(time.Now().UnixMilli())
	var serial string
	for i := len(ss) - 1; i >= 0; i-- {
		serial += string(ss[i])
	}
	if len(serial) > length {
		serial = serial[:length]
	}
	return serial
}

const (
	serialCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// RandomSerial will generate an alphanumeric serial number from cryptographic random data
func RandomSerial(length int) string {
	max := big.NewInt(int64(len(serialCharset)))
	serialNumber := make([]byte, length)
	for i := range serialNumber {
		randIndex, _ := rand.Int(rand.Reader, max)
		serialNumber[i] = serialCharset[randIndex.Int64()]
	}
	return string(serialNumber)
}

// PaddedSerial will format a serial number to the intended length
func PaddedSerial(input string) SerialGenerator {
	return func(length int) string {
		if len(input) == length {
			return input
		}
		// padding := fmt.Sprintf("%d.%d", length, length
		return fmt.Sprintf(fmt.Sprintf("%%0%d.%dv", length, length), input)
	}
}

func pickGenerator(fn SerialGenerator, fallback SerialGenerator) SerialGenerator {
	if fn != nil {
		return fn
	}
	return fallback
}
