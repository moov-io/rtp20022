package rtp

import (
	"fmt"
	"time"
)

// TransactionID
//
// Format:  YYYYMMDDbbbbbbbbbbbBRRRRnnnnnnnnnnn
// Example: 2015111502120020101BCPAA00000000001
func TransactionID(ts time.Time, participantID string, serial string) string {
	creationDate := ts.Format("20060102")
	messageSerial := NumericSerialNumber(11)

	return fmt.Sprintf("%s%11.11sB%4.4s%s", creationDate, participantID, serial, messageSerial)
}
