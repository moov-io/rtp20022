package camt_029_001_09

func (d DocumentTCH) Validate() error {
	return d.RsltnOfInvstgtn.Validate()
}

func (c ResolutionOfInvestigationV09TCH) Validate() error {
	return nil
}
