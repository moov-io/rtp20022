package dt

import (
	"fmt"
	"sync"
	"time"
)

var (
	eastern       *time.Location
	locationSetup sync.Once
)

// Eastern returns a *time.Location for the US Eastern region.
// RTP requires all timestamps are in this location.
func Eastern() *time.Location {
	// RTP specifies that all timestamps are in Eastern (Daylight or Standard)
	// so we will need to use this *time.Location many times. Rather than load
	// it every time we are going to lazy cache it when needed.
	locationSetup.Do(func() {
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			panic(fmt.Sprintf("Unable to load America/New_York from tzdata: %v", err))
		}
		eastern = loc
	})
	return eastern
}
