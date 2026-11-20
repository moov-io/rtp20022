package rtp_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func Example_parse() {
	input, err := os.ReadFile(filepath.Join("test", "testdata", "pacs008.xml"))
	if err != nil {
		panic(err)
	}

	var msg messages.Message
	if err := xml.Unmarshal(input, &msg); err != nil {
		panic(err)
	}

	fmt.Println(msg.CreditTransfer.FIToFICstmrCdtTrf.GrpHdr.MsgId)
	fmt.Println(msg.CreditTransfer.Validate() == nil)
	// Output:
	// M20230713234567891T1BOTS02083825666
	// true
}

func Example_generate() {
	now := time.Date(2023, 7, 13, 15, 29, 36, 0, rtp.Eastern())
	serial := func(length int) string {
		return fmt.Sprintf("%0*d", length, 1)
	}

	msg := messages.NewPacs008Message()
	msg.AppHdr.CreDt = rtp.ISODateTime(now)
	msg.CreditTransfer = &pacs_008_001_08.DocumentTCH{
		FIToFICstmrCdtTrf: pacs_008_001_08.FIToFICustomerCreditTransferV08TCH{
			GrpHdr: pacs_008_001_08.GroupHeader93TCH{
				MsgId:   pacs_008_001_08.Max35TextTCH(rtp.MessageID(now, "234567891T1", "OTS", serial)),
				CreDtTm: rtp.ISODateTime(now),
				NbOfTxs: "1",
			},
			CdtTrfTxInf: pacs_008_001_08.CreditTransferTransaction39TCH{
				PmtId: pacs_008_001_08.PaymentIdentification7TCH{
					InstrId: pacs_008_001_08.Max35TextTCH2(rtp.InstructionID(now, "234567891T1", "CXIA", serial)),
					TxId:    pacs_008_001_08.Max35Text(rtp.TransactionID(now, "234567891T1", "CXIA", serial)),
				},
			},
		},
	}

	out, err := xml.MarshalIndent(msg, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(rtp.MessageID(now, "234567891T1", "OTS", serial))
	fmt.Println(strings.Contains(string(out), `xmlns:head="urn:iso:std:iso:20022:tech:xsd:head.001.001.01"`))
	fmt.Println(strings.Contains(string(out), `xmlns:ct="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"`))
	// Output:
	// M20230713234567891T1BOTS00000000001
	// true
	// true
}
