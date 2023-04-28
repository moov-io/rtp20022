package pacs_002_001_10

func (d Document) Validate() error {
	return d.FIToFIPmtStsRpt.Validate()
}
func (p *FIToFIPaymentStatusReportV10) Validate() error {
	return nil // TODO(adam):
}
