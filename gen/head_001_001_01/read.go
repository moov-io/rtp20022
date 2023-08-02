package head_001_001_01

func (r BusinessApplicationHeaderV01TCH) CopyDuplicateCode() string {
	if r.CpyDplct != nil {
		return string(*r.CpyDplct)
	}
	return ""
}

func (r BusinessApplicationHeaderV01TCH) FromMemberID() string {
	if r.Fr.FIId != nil {
		return string(r.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	return ""
}

func (r BusinessApplicationHeaderV01TCH) ToMemberID() string {
	if r.To.FIId != nil {
		return string(r.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	return ""
}
