// Package gigasecond can add a gigasecond
package gigasecond

import "time"

// AddGigasecond adds 1,000,000,000 seconds to the given time t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000000000))
}
