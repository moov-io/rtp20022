package admi_002_001_01

func (d DocumentTCH) Validate() error {
	return d.Admi00200101.Validate()
}

func (s MessageRejectV01TCH) Validate() error {
	return nil // TODO(adam):
}
