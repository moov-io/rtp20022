package camt_029_001_09

func (d Document) Validate() error {
	return d.RsltnOfInvstgtn.Validate()
}

func (c ResolutionOfInvestigationV09) Validate() error {
	return nil
}
