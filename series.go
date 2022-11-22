package finance

import (
	"math"
)

// Series is a wrapper for statistics on a slice of numbers.
//
type Series struct {
	v        []float64
	n        float64
	Sum      float64 // Sum of the series.
	Mean     float64 // Mean of the series.
	MAD      float64 // Mean absolute deviation of the series.
	SumD2    float64 // Sum of the square differences to the mean.
	Variance float64 // Variance of the series.
	StdDev   float64 // Standard deviation of the series.
}

// NewSeries makes a new Series for the given slice.
//
func NewSeries(v []float64) (x Series) {
	if len(v) == 0 {
		panic("zero length dataset")
	}
	x.v = v
	x.n = float64(len(v))
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

// Compare returns the analysis of this series compared to the given series.
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
	xy.Covariance = sumdxdy / x.n
	xy.R = xy.Covariance / (x.StdDev * y.StdDev)
	xy.R2 = xy.R * xy.R
	xy.Gradient = xy.Covariance / x.Variance
	xy.Intercept = y.Mean - xy.Gradient*x.Mean
	return
}
