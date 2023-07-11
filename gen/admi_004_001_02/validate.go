package admi_004_001_02

func (d DocumentTCH) Validate() error {
	return d.SysEvtNtfctn.Validate()
}

func (s SystemEventNotificationV02TCH) Validate() error {
	return nil // TODO(adam):
}
