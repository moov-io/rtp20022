package camt_056_001_08

func (d Document) Validate() error {
	return d.FIToFIPmtCxlReq.Validate()
}

func (c FIToFIPaymentCancellationRequestV08) Validate() error {
	return nil
}
