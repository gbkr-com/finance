package finance

import (
	"math"
)

// LinearReturns calculates the linear returns of the given prices p into the
// slice r. Slice r must have length at least equal to the length of p minus one
// otherwise the function will panic. The function returns the number of values
// written into r.
//
// Linear returns are asset additive, i.e. if r is the linear return and w is
// the weight in the portfolio then:
//
//		portfolio return = 𝛴 w . r
//
// Linear returns should be used for portfolio allocation.
//
func LinearReturns(p, r []float64) (n int) {
	n = len(p) - 1
	if len(r) < n {
		panic("insuffient length")
	}
	for i := 1; i < len(p); i++ {
		r[i-1] = (p[i] / p[i-1]) - 1
	}
	return
}

// CompoundReturns calculates the log returns of the given prices p into the
// slice r. Slice r must have length at least equal to the length of p minus one
// otherwise the function will panic. The function returns the number of values
// written into r.
//
// Compound returns are time additive, i.e. if r is the compound return for a
// single period then:
//
//		N period return = 𝛴 r
//
// however, that can be simplified to ln(pn) - ln(p0).
//
func CompoundReturns(p, r []float64) (n int) {
	n = len(p) - 1
	if len(r) < n {
		panic("insuffient length")
	}
	for i := 1; i < len(p); i++ {
		r[i-1] = math.Log(p[i] / p[i-1]) // Division is less costly than two math.Log calls.
	}
	return
}
