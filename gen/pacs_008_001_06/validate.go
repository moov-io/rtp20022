package pacs_008_001_06

func (d Document) Validate() error {
	return d.FIToFICstmrCdtTrf.Validate()
}

func (c FIToFICustomerCreditTransferV06) Validate() error {
	return nil // TODO(adam):
}
