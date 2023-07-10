package camt_035_001_05

func (d DocumentTCH) Validate() error {
	return d.PrtryFrmtInvstgtn.Validate()
}

func (c ProprietaryFormatInvestigationV05TCH) Validate() error {
	return nil // TODO(nikhil):
}
