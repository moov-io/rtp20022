package rtp

import (
	"fmt"
	"time"
)

// TransactionID
//
// Format:  YYYYMMDDbbbbbbbbbbbBRRRRnnnnnnnnnnn
// Example: 2015111502120020101BCPAA00000000001
func TransactionID(ts time.Time, participantID string, serial string, messageSerialFunc SerialGenerator) string {
	creationDate := ts.Format("20060102")
	messageSerial := pickGenerator(messageSerialFunc, NumericSerialNumber)(11)

	return fmt.Sprintf("%s%011.11sB%04.4s%011.11s", creationDate, participantID, serial, messageSerial)
}
