package rtp

type Cdata struct {
	CDataString string `xml:",cdata"`
}

func (c Cdata) Validate() error {
	return nil
}
