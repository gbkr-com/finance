package finance

import (
	"math"
)

// CDF returns the cumulative distribution of the standard normal distribution
// for the given z.
//
func CDF(z float64) float64 {
	z = -z * InverseSqrt2
	return 0.5 * math.Erfc(z)
}

// PDF returns the probability density of the standard normal distribution for
// the given z.
//
func PDF(z float64) float64 {
	z = -z * z / 2
	return InverseSqrt2Pi * math.Exp(z)
}
