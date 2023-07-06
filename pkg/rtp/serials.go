package rtp

import (
	"encoding/base32"
	"fmt"
	"time"
)

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
