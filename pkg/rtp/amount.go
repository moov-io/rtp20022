package rtp

import "fmt"

type Amount float64

func (a Amount) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", a)), nil
}
