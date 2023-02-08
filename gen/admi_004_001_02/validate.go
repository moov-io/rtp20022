package admi_004_001_02

func (d Document) Validate() error {
	return d.SysEvtNtfctn.Validate()
}

func (s SystemEventNotificationV02) Validate() error {
	return nil // TODO(adam):
}
