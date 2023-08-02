package admi_002_001_01

func (r MessageRejectV01TCH) Reference() string {
	return string(r.RltdRef.Ref)
}

func (r MessageRejectV01TCH) Reason() string {
	return string(r.Rsn.RjctgPtyRsn)
}

func (r MessageRejectV01TCH) AdditionalData() string {
	if r.Rsn.AddtlData != nil {
		return r.Rsn.AddtlData.CDataString
	}
	return ""
}
