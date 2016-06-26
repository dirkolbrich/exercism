package diffsquares

// Difference calculates what it says
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SquareOfSums calculates the square of the summ of each int from n
func SquareOfSums(n int) int {
	var sum int
	for i := 0; i < n+1; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares summs up the squares of ech int from n
func SumOfSquares(n int) int {
	var sum int
	for i := 0; i < n+1; i++ {
		sum += i * i
	}
	return sum
}
