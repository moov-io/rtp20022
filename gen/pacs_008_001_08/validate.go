package pacs_008_001_08

func (d Document) Validate() error {
	return d.FIToFICstmrCdtTrf.Validate()
}

func (c FIToFICustomerCreditTransferV08) Validate() error {
	return nil // TODO(adam):
}
