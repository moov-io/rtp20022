package admi_002_001_01

func (r Admi00200101) Reference() string {
	return string(r.RltdRef.Ref)
}

func (r Admi00200101) Reason() string {
	return string(r.Rsn.RjctgPtyRsn)
}

func (r Admi00200101) AdditionalData() string {
	if r.Rsn.AddtlData != nil {
		return r.Rsn.AddtlData.CDataString
	}
	return ""
}
