package hamming

import "errors"

const testVersion = 4

// Distance compares two strings and returns the number of differences
func Distance(a, b string) (int, error) {
	// check length of strings
	if len(a) != len(b) {
		return 0, errors.New("Strings are not of same length.")
	}

	var dist int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
