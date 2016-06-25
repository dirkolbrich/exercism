package gigasecond

import "time"

const testVersion = 4

// AddGigasecond calculates the gigasecond bases on the birthday
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
