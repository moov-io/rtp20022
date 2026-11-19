[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/rtp20022?status.svg)](https://godoc.org/github.com/moov-io/rtp20022)
[![Build Status](https://github.com/moov-io/rtp20022/workflows/Go/badge.svg)](https://github.com/moov-io/rtp20022/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/rtp20022/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/rtp20022)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/rtp20022)](https://goreportcard.com/report/github.com/moov-io/rtp20022)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/rtp20022?label=project%20size)](https://github.com/moov-io/rtp20022)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/rtp20022/master/LICENSE)
[![Slack Channel](https://img.shields.io/badge/Slack-Moov-e01563)](https://slack.moov.io/)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/rtp20022)](https://github.com/moov-io/rtp20022)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

## moov-io/rtp20022

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

This repository contains a subset of ISO 20022 messages for RTP payments. ISO 20022 is a standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions, and other financial information.

## Project Status

The Go types in `gen/` are generated from the TCH XSDs with [moov-io/xsd2go](https://github.com/moov-io/xsd2go). This project is used to generate messages and pass certification of 2.9 and 3.0 RTP/TCH releases. Please star the project if you are interested in its progress, and open an issue if you encounter bugs, unclear documentation, or have feature suggestions. Thanks!

This repository is a Go library. It does not ship a CLI or HTTP server.

## Usage

### Go Library

Install the module and run `go mod tidy` so `go.sum` includes the library's dependencies (including `github.com/moov-io/base`). A bare `go get github.com/moov-io/rtp20022@latest` in a new module can leave the dependency marked `// indirect` and later builds may fail with a missing `go.sum` entry.

```
go get github.com/moov-io/rtp20022@latest
go mod tidy
```

This project requires **Go 1.26 or higher**. The current `go.mod` also sets `toolchain go1.27.4`.

The `messages.Message` type can read and write XML-encoded TCH RTP messages. Use the `NewXxxMessage()` constructors when writing so the document is created with the required `xmlns` prefixes. Reading TCH messages that already use prefixes (for example `head:` and `ct:`) works with `encoding/xml` and this package's custom unmarshaler. Go still has [general limitations around XML namespace prefixes](https://github.com/golang/go/issues/13400); the constructors exist so you do not have to assemble those prefixes yourself.

Runnable parse and generate examples live in `pkg/rtp/example_test.go`. They use `messages.Message`, `messages.NewPacs008Message()`, `Validate()`, and the ID helpers in `pkg/rtp`.

```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/moov-io/rtp20022/gen/messages"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func main() {
	// Parse a prefixed TCH sample
	input, err := os.ReadFile("pkg/rtp/test/testdata/pacs008.xml")
	if err != nil {
		panic(err)
	}
	var parsed messages.Message
	if err := xml.Unmarshal(input, &parsed); err != nil {
		panic(err)
	}
	fmt.Println(parsed.CreditTransfer.FIToFICstmrCdtTrf.GrpHdr.MsgId)
	if err := parsed.CreditTransfer.Validate(); err != nil {
		panic(err)
	}

	// Generate a pacs.008 document with the required namespace prefixes
	now := time.Now()
	msg := messages.NewPacs008Message()
	msg.AppHdr.CreDt = rtp.ISODateTime(now)
	msg.CreditTransfer = &pacs_008_001_08.DocumentTCH{
		FIToFICstmrCdtTrf: pacs_008_001_08.FIToFICustomerCreditTransferV08TCH{
			GrpHdr: pacs_008_001_08.GroupHeader93TCH{
				MsgId:   pacs_008_001_08.Max35TextTCH(rtp.MessageID(now, "234567891T1", "OTS", rtp.NumericSerialNumber)),
				CreDtTm: rtp.ISODateTime(now),
				NbOfTxs: "1",
			},
		},
	}
	out, err := xml.MarshalIndent(msg, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
```

### Supported messages

The library covers the following TCH message types. Each constructor returns a `*messages.Message` with the correct XML namespaces for writing.

| Message | Constructor |
| --- | --- |
| acmt.022.001.02 (`TokenIdentification`) | `NewAcmt022Message()` |
| admi.002.001.01 (`MessageReject`) | `NewAdmi002Message()` |
| admi.004.001.02 (`SystemNotificationEvent`) | `NewAdmi004Message()` |
| admn.001.001.01 (`SignOnRequest`) | `NewAdmn001Message()` |
| admn.002.001.01 (`SignOnResponse`) | `NewAdmn002Message()` |
| admn.003.001.01 (`SignOffRequest`) | `NewAdmn003Message()` |
| admn.004.001.01 (`SignOffResponse`) | `NewAdmn004Message()` |
| admn.005.001.01 (`EchoRequest`) | `NewAdmn005Message()` |
| admn.006.001.01 (`EchoResponse`) | `NewAdmn006Message()` |
| admn.007.001.01 (`ParticipantReport`) | `NewAdmn007Message()` |
| admn.008.001.01 (`ParticipantReportResponse`) | `NewAdmn008Message()` |
| camt.026.001.07 (`RequestForInformation`) | `NewCamt026Message()` |
| camt.028.001.09 (`ResponseRequestForInformation`) | `NewCamt028Message()` |
| camt.029.001.09 (`ResponseReturnOfFunds`) | `NewCamt029Message()` |
| camt.035.001.05 (`Acknowledgement`) | `NewCamt035Message()` |
| camt.056.001.08 (`ReturnOfFunds`) | `NewCamt056Message()` |
| pacs.002.001.10 (`MessageStatusReport`) | `NewPacs002Message()` |
| pacs.008.001.08 (`CreditTransfer`) | `NewPacs008Message()` |
| pacs.009.001.08 (`FICreditTransfer`) | `NewPacs009Message()` |
| pacs.028.001.03 (`PaymentStatusRequest`) | `NewPacs028Message()` |
| pain.013.001.07 (`PaymentRequest`) | `NewPain013Message()` |
| pain.014.001.07 (`ResponsePaymentRequest`) | `NewPain014Message()` |
| remt.001.001.04 (`StandaloneRemittance`) | `NewRemt001Message()` |

ID helpers for writing messages are in `pkg/rtp`, including `MessageID`, `InstructionID`, `TransactionID`, and `AdmnMessageID`.

### Regenerating code

The types in `gen/` are produced by [moov-io/xsd2go](https://github.com/moov-io/xsd2go) from the XSDs in `xsd/`, not by `aqwari.net/xml/cmd/xsdgen`. To regenerate:

```
make install-generator
make generate
```

`make install-generator` installs `moovio_xsd2go` and `goimports`. `make generate` runs `scripts/generate.sh`, then formats the generated files.

## Getting Help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov) | You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/rtp20022/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#rtp20022`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
