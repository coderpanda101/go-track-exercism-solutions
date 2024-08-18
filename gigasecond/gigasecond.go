package gigasecond

import (
	"time"
)

// AddGigasecond provides date and time one Gigasecond after the provided date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
