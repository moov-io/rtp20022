package camt_035_001_05

func (d Document) Validate() error {
	return d.PrtryFrmtInvstgtn.Validate()
}

func (c ProprietaryFormatInvestigationV05) Validate() error {
	return nil // TODO(nikhil):
}
