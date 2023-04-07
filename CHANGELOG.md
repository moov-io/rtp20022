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
