package pacs_002_001_10

func (d DocumentTCH) Validate() error {
	return d.FIToFIPmtStsRpt.Validate()
}

func (p *FIToFIPaymentStatusReportV10TCH) Validate() error {
	return nil // TODO(adam):
}
