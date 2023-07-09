package admi_002_001_01

func (d Document) Validate() error {
	return d.Admi00200101.Validate()
}

func (s MessageRejectV01) Validate() error {
	return nil // TODO(adam):
}
