package finance

import (
	"math"
)

// Closing prices for HSBC (5) in May 2022.
//
var HSBC = []float64{
	49.80,
	50.75,
	50.45,
	49.00,
	48.45,
	48.40,
	46.75,
	47.75,
	47.90,
	48.60,
	48.35,
	47.55,
	48.45,
	48.45,
	49.40,
	50.95,
	51.20,
	51.95,
	52.40,
	52.10,
}

// Closing prices for Sands China (1928) in May 2022.
//
var SANDS = []float64{
	18.12,
	18.0,
	17.44,
	16.5,
	15.36,
	15.7,
	15.38,
	15.76,
	15.68,
	15.92,
	15.92,
	15.08,
	15.3,
	15.14,
	14.0,
	13.98,
	13.86,
	14.2,
	14.82,
	15,
}

// Closing prices for the HSI in May 2022.
//
var HSI = []float64{
	21101.89,
	20869.52,
	20793.40,
	20001.96,
	19633.69,
	19824.57,
	19380.34,
	19898.77,
	19950.21,
	20602.52,
	20644.28,
	20120.68,
	20717.24,
	20470.06,
	20112.10,
	20171.27,
	20116.20,
	20697.36,
	21123.93,
	21415.20,
}

// Variables to prevent destructive compiler optimisation in benchmarks.
//
var (
	dummy       float64
	dummySeries []float64
)

// Round to n decimals.
//
func Round(x float64, n int) float64 {
	f := math.Pow10(n)
	return math.Round(x*f) / f
}
