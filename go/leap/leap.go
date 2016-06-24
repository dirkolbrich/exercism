package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// functions takes a year as integer and returns true if year is leap year, otherwise returns false
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
