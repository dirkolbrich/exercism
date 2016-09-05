package grains

import (
	"errors"
	"math"
)

var ErrInvalid = errors.New("Invalid field number given.")

func validField(f int) bool {
	return f >= 1 && f <= 64
}

// Square calculates the grains on a given field
func Square(field int) (uint64, error) {
	if !validField(field) {
		return 0, ErrInvalid
	}
	return uint64(math.Exp2(float64(field - 1))), nil
}

// Total calculates the grains on all fields of a chess board
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		total = total + uint64(math.Exp2(float64(i-1)))
	}
	return total
}
