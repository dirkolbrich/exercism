package triangle

import "math"

const testVersion = 2

type Kind struct {
	name string
}

var NaT = Kind{"NaT"} // not a triangle
var Equ = Kind{"Equ"} // equilateral - all three sides are equal in length
var Iso = Kind{"Iso"} // isosceles - 2 sides are of equaL length
var Sca = Kind{"Sca"} // scalene - all three sides are of different length

// KindFromSides tests the three sides of a triangle
func KindFromSides(a, b, c float64) Kind {
	switch {
	// side is 0
	case (a == 0 || b == 0 || c == 0):
		return NaT
		// sum of two sides shorter than third
	case (a+b < c || b+c < a || a+c < b):
		return NaT
		// side is infinite
	case (a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1)):
		return NaT
		// side is negativ infinite
	case (a == math.Inf(-1) || b == math.Inf(-1) || c == math.Inf(-1)):
		return NaT
	case (a == b && a == c && a != 0):
		return Equ
	case (a == b || a == c || b == c) && (a > 0 || b > 0):
		return Iso
	case (a != b && a != c && b != c) && (a > 0 && b > 0 && c > 0):
		return Sca
	default:
		return NaT
	}
}
