package clock

import "strconv"

const testVersion = 4

// represents the Clock
type Clock struct {
	h int
	m int
}

// creates new clock from given amount of hours and minutes
func New(hour int, minute int) Clock {
	h, m := hour, minute

	// convert minutes to hours and minutes
	switch {
	case minute >= 60:
		m = minute % 60
		h += minute / 60
	case minute < 0:
		if (-minute)%60 == 0 {
			m = 0
		} else {
			m = 60 - (-minute)%60
		}
		h -= 1 + (-minute)/60
	}

	// convert hours to days and hours
	switch {
	case h >= 24:
		h = h % 24
	case h < 0:
		if (-h)%24 == 0 {
			h = 0
		} else {
			h = 24 - (-h)%24
		}
	}

	return Clock{h, m}
}

func (c Clock) String() string {
	var hour, minute string

	switch {
	case c.h < 10:
		hour = "0" + strconv.Itoa(c.h)
	default:
		hour = strconv.Itoa(c.h)
	}

	switch {
	case c.m < 10:
		minute = "0" + strconv.Itoa(c.m)
	default:
		minute = strconv.Itoa(c.m)
	}

	return hour + ":" + minute
}

func (c Clock) Add(minutes int) Clock {
	h := c.h
	m := c.m + minutes

	return New(h, m)
}
