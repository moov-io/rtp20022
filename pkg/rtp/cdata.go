package rtp

type Cdata struct {
	CDataString string `xml:",cdata"`
}

func (c Cdata) Validate() error {
	// TODO JB:
	return nil
}
