package raindrops

import "strconv"

const testVersion = 2

// Convert test if Number n is divisable by 3, 5 or 7 and adds corresponding
// Pling, Plang or Plong to string
func Convert(n int) string {
	var s string

	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}

	return s
}
