package slice

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var all []string

	if n > len(s) {
		return all
	}

	for i := 0; i <= len(s)-n; i++ {
		all = append(all, s[i:i+n])
	}

	return all
}

// Frist returns the first substring of s with length n.
func Frist(n int, s string) string {
	return string(s[0:n])
}

// First returns first substring and checks for length of n vs. s
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "n is longer than s", false
	}
	return Frist(n, s), true
}
