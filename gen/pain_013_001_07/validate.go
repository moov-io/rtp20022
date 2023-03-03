package pain_013_001_07

func (d Document) Validate() error {
	return d.CdtrPmtActvtnReq.Validate()
}

func (c CreditorPaymentActivationRequestV07) Validate() error {
	return nil // TODO(adam):
}
