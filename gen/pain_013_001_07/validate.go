package pain_013_001_07

func (d DocumentTCH) Validate() error {
	return d.CdtrPmtActvtnReq.Validate()
}

func (c CreditorPaymentActivationRequestV07TCH) Validate() error {
	return nil // TODO(adam):
}
