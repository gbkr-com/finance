package finance

import (
	"math"
)

// Series the principal statistics on a series of numbers, for example price
// returns.
//
type Series struct {
	v        []float64
	n        float64
	sample   bool
	Sum      float64 // Sum of the series.
	Mean     float64 // Mean of the series.
	MAD      float64 // Mean absolute deviation of the series.
	SumD2    float64 // Sum of the square differences from the mean.
	Variance float64 // Variance of the series.
	StdDev   float64 // Standard deviation of the series, a.k.a. volatility when this is a series of returns.
}

// NewSeries makes a new Series for the given []float64, calculating the principal
// statistics on those numbers. If the series is a sample then the variance is
// adjusted accordingly.
//
// The given slice must contain at least two numbers, otherwise this function
// panics.
//
func NewSeries(v []float64, sample bool) (x Series) {
	if len(v) < 2 {
		panic("insufficient data")
	}
	x.v = v
	x.n = float64(len(v))
	x.sample = sample
	//
	// Calculate mean.
	//
	for _, v := range v {
		x.Sum += v
	}
	x.Mean = x.Sum / x.n
	//
	// Calculate statistics dependent upon the mean.
	//
	for _, v := range v {
		d := v - x.Mean
		x.MAD += math.Abs(d)
		x.SumD2 += d * d
	}
	x.MAD /= x.n
	if sample {
		x.Variance = x.SumD2 / (x.n - 1)
	} else {
		x.Variance = x.SumD2 / x.n
	}
	x.Variance = x.SumD2 / x.n
	x.StdDev = math.Sqrt(x.Variance)
	return
}

// Analysis between two series.
//
type Analysis struct {

	// The Covariance of the two series.
	//
	Covariance float64

	// The correlation coefficient.
	//
	R float64

	// The coefficient of determination.
	//
	R2 float64

	// The Gradient from the linear regression of the two series. When x are the
	// returns from an index and y are the returns from an index constituent,
	// then Gradient is the Beta factor for that constituent.
	//
	Gradient float64

	// The Intercept from the linear regression.
	//
	Intercept float64
}

// Compare returns the analysis of this series compared to the given series. If
// this series was made as a sample then the covariance will also be calculated
// as a sample.
//
// If the series differ in length then this function will panic.
//
func (x Series) Compare(y Series) (xy Analysis) {
	if y.n != x.n {
		panic("incompatible series")
	}
	var sumdxdy float64
	for i := range x.v {
		dx := x.v[i] - x.Mean
		dy := y.v[i] - y.Mean
		sumdxdy += dx * dy
	}
	if x.sample {
		xy.Covariance = sumdxdy / (x.n - 1)
	} else {
		xy.Covariance = sumdxdy / x.n
	}
	xy.Covariance = sumdxdy / x.n
	xy.R = xy.Covariance / (x.StdDev * y.StdDev)
	xy.R2 = xy.R * xy.R
	xy.Gradient = xy.Covariance / x.Variance
	xy.Intercept = y.Mean - xy.Gradient*x.Mean
	return
}
