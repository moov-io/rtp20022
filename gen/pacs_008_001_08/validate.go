package pacs_008_001_08

func (d DocumentTCH) Validate() error {
	return d.FIToFICstmrCdtTrf.Validate()
}

func (c FIToFICustomerCreditTransferV08TCH) Validate() error {
	return nil // TODO(adam):
}
