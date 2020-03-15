// Package triangle has a method that checks the type of a triangle
package triangle

import (
	"math"
)

// Kind is our custom type
type Kind int

const (
	// NaT show return when it's not a triangle
	NaT Kind = iota
	// Equ when triangle has all three sides the same length
	Equ
	//Iso triangle has at least two sides the same length AND consequently two equal angles
	Iso
	//Sca triangle has all sides of different lengths
	Sca
	//Deg sum of the lengths of two sides equals the third
	Deg
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	if a+b == c || a+c == b || b+c == a {
		return Deg
	}
	return Sca
}
