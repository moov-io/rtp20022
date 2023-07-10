package camt_056_001_08

func (d DocumentTCH) Validate() error {
	return d.FIToFIPmtCxlReq.Validate()
}

func (c FIToFIPaymentCancellationRequestV08TCH) Validate() error {
	return nil
}
