package pascal

// Triangle calculate pascals-triangle given by number of rows n
func Triangle(n int) [][]int {
	var tri [][]int

	// first row starts always with 1
	tri = append(tri, []int{1})

	// each n is a line
	for l := 1; l < n; l++ {
		lineLength := l + 1
		line := make([]int, lineLength)
		prevLi := tri[l-1]

		var middleIndex int
		// get middle index of line
		if lineLength%2 == 0 {
			// line is even
			middleIndex = (lineLength / 2) - 1
		} else {
			// line is odd
			middleIndex = lineLength / 2
		}

		// each line has li elements
		for i := middleIndex; i > 0; i-- {
			a := prevLi[i-1]
			b := prevLi[i]
			line[i], line[(lineLength-1)-i] = a+b, a+b
		}

		// first and last integer are always 1
		line[0], line[lineLength-1] = 1, 1

		// append line to triangle
		tri = append(tri, line)
	}

	return tri
}
