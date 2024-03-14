package rtp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// AlphaSerialNumber generates a random alphanumeric serial number of the specified length
// without depending on the current time.
func AlphaSerialNumber(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	serialNumber := make([]byte, length)

	for i := range serialNumber {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		serialNumber[i] = charset[randIndex.Int64()]
	}
	return string(serialNumber)
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
