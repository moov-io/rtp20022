## v0.11.0 (Released 2024-03-20)

This release of moov-io/signedxml changes how IDs are generated. Previously IDs were generated with the current time
being used as randomness. Now IDs have their serial numbers generated with functions. The previous behavior can be
restored by passing `NumericSerialNumber` to each instance of an ID function.

**Before**
```go
messageID := rtp.MessageID(time.Now(), participantID, serial)
```

**After**
```go
messageID := rtp.MessageID(time.Now(), participantID, serial, rtp.NumericSerialNumber)
```

BREAKNIG CHANGES

- rtp: expose serial generator as a function, with helpers

BUILD

- fix(deps): update module cloud.google.com/go to v0.112.1
- fix(deps): update module github.com/stretchr/testify to v1.9.0

## v0.10.2 (Released 2024-02-12)

IMPROVEMENTS

- fix: correct typo in the validate template for length restrictions
- test: verify rtp.Amount marshals large amounts in decimal notation

BUILD

- fix(deps): update module cloud.google.com/go to v0.112.0
- fix(deps): update module github.com/moov-io/base to v0.48.5

## v0.10.1 (Released 2024-01-11)

IMPROVEMENTS

- fix: make ValidateLength error more clear

BUILD

- fix(deps): update module github.com/moov-io/base to v0.48.4

## v0.10.0 (Released 2023-12-18)

IMPROVEMENTS

- feat: add camt.029 Case Assignment ID

BUILD

- fix(deps): update module github.com/moov-io/base to v0.48.3

## v0.9.5 (Released 2023-12-14)

IMPROVEMENTS

- fix: pad bankField in InstructionID when short

## v0.9.4 (Released 2023-12-07)

IMPROVEMENTS

- feat: include field names in Validate() errors
- tests: add pacs.008 invalid check

BUILD

- fix(deps): update module cloud.google.com/go to v0.111.0

## v0.9.3 (Released 2023-11-28)

IMPROVEMENTS

- rtp: cache regexes in memory for reuse

BUILD

- fix(deps): update module cloud.google.com/go to v0.110.10
- fix(deps): update module github.com/moov-io/base to v0.48.2

## v0.9.2 (Released 2023-08-04)

ADDITIONS

- feat: Generated code improvements (no functional changes)

## v0.9.0 v0.9.1 (Released 2023-07-25)

- feat: convert ISODate types underlying type to civil.Date
- feat: Remove unnecessary types ISOTime and ISONormalisedDateTime

## v0.8.0 (Released 2023-07-21)

ADDITIONS

- feat: generate the struct for the signature from XSD
- feat: rename schemas to match 3.0 versions
- feat: XSD version 3.0
- feat: add and update unit tests for models
- feat: add Context method for some models
- feat: auto-generate Validate functions
- feat: add test coverage for validation

BUILD

- fix(deps): use the latest stable Go release
- fix(deps): update module github.com/moov-io/base to v0.45.1

## v0.7.0 (Released 2023-07-06)

ADDITIONS

This release includes a reworked version of the XML structs and related code for rtp20022 schemas based on a new generator. This release of rtp20022 has been verified to comply with TCH testing and verification processes.

- feat: add CDATA support
- feat: update XML struct generation
- feat: add `admn.001.001.01`
- feat: add `admn.002.001.01`
- feat: add `admn.003.001.01`
- feat: add `admn.004.001.01`
- feat: add `admn.005.001.01`
- feat: add `admn.006.001.01`
- feat: add `head.001.001.01`

BUILD

- fix(deps): update module github.com/moov-io/base to v0.45.0
- fix(deps): update module github.com/stretchr/testify to v1.8.4

## v0.6.0 (Released 2023-05-17)

ADDITIONS

- feat: response return of funds writer

## v0.5.0 (Released 2023-05-09)

ADDITIONS

- feat: Changes for rtp-test-harness
- feat: camt.056.001.08
- feat: camt.029.001.09

BUILD

- fix(deps): update module github.com/moov-io/base to v0.42.0
- fix(deps): update module github.com/go-logfmt/logfmt to v0.6.0

## v0.4.0 (Released 2023-04-19)

- feat: status reason

## v0.3.0 (Released 2023-04-07)

ADDITIONS

- feat: generate pain.013.001.07
- gen/pain_013_001_07: add Validate()
- feat: add AppHdr (head.001) reader and writer
- feat: add credit_transfer (pacs.008) reader and writer
- feat: add msg_status (pacs.002) reader and writer
- feat: add setup for reading pacs.008 and decoding xml

BUILD

- fix(deps): update module github.com/stretchr/testify to v1.8.2

## v0.2.0 (Released 2023-02-14)

This release of rtp20022 has been verified to comply with TCH testing and verification processes. We have had to modify the generated code to comply with TCH alterations of ISO20022.

ADDITIONS

- feat: add `admi.002.001.01`
- feat: add `admi.004.001.02`
- feat: add `camt.035.001.05`
- feat: add `pacs.008.001.08`

## v0.1.0 (Released 2023-01-09)

Initial release of rtp20022 with pacs.002.001.07 model.
