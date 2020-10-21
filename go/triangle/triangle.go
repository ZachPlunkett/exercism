// Package triangle provides a triangle type finding function
package triangle

import "math"

// Kind is either a triangle type or not a triangle
type Kind string

// Constants for types of triangles as well as not a triangle
const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// realNumber determines if three given numbers are realnumbers
func realNumber(a, b, c float64) bool {
	nums := []float64{a, b, c}

	for _, i := range nums {
		if math.IsNaN(i) || math.IsInf(i, 0) {
			return false
		}
	}
	return true
}

// KindFromSides returns the type of a triangle based upon its three sides.
func KindFromSides(a, b, c float64) Kind {

	// Not a Triangle
	if !(a > 0 && b > 0 && c > 0 && a+b >= c && a+c >= b && b+c >= a) ||
		!realNumber(a, b, c) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
