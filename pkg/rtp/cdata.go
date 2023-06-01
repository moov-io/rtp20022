package rtp

// May be no more than 20000 items long
type Max20000Text string

type Cdata struct {
	Max20000Text `xml:",cdata"`
}
