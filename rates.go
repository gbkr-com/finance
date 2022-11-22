package finance

import (
	"math"
)

// Discount returns the factor to discount a value given the continuous
// compounded rate over unit time. That time is typically fractional years.
//
func Discount(rate, time float64) float64 {
	return math.Exp(-rate * time)
}

// Compound returns the factor to compound at the given continuous rate over
// unit time. That time is typically fractional years.
//
func Compound(rate, time float64) float64 {
	return math.Exp(rate * time)
}

// ToContinuous returns the continuous compounded rate equivalent to the
// given single compounded rate.
//
func ToContinuous(rate float64) float64 {
	return math.Log(rate + 1)
}

// ToSingle returns the single compunded rate equivalent to the given continuous
// compounded rate.
//
func ToSingle(rate float64) float64 {
	return Compound(rate, 1) - 1
}
