package secret

import "strconv"

func Handshake(n int) []string {
	var hs []string

	if n < 0 {
		return hs
	}

	// convert int to binary
	binary := strconv.FormatInt(int64(n), 2)

	// reverse the string, so we can check the binary backwards
	var reverse string
	for _, b := range binary {
		reverse = string(b) + reverse
	}

	// check each binary fo 0 or 1
	for pos, b := range reverse {
		switch {
		case pos == 0 && string(b) == "1":
			hs = append(hs, "wink")
		case pos == 1 && string(b) == "1":
			hs = append(hs, "double blink")
		case pos == 2 && string(b) == "1":
			hs = append(hs, "close your eyes")
		case pos == 3 && string(b) == "1":
			hs = append(hs, "jump")
		// reverse slice
		case pos == 4 && string(b) == "1":
			for i, j := 0, len(hs)-1; i < j; i, j = i+1, j-1 {
				hs[i], hs[j] = hs[j], hs[i]
			}
		}
	}

	return hs
}
