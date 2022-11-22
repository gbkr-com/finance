package finance

import (
	"math"
)

// BlackScholes represents the Black-Scholes model.
//
type BlackScholes struct {
	Price        float64
	Strike       float64
	Term         float64 // Fraction of unit time, range (0, 1]
	Yield        float64 // Income over unit time.
	Volatility   float64 // Volatility over unit time.
	RiskFreeRate float64
}

// BlackScholesResult is the result from the model.
//
type BlackScholesResult struct {
	D1        float64 // The d1 term in the Black Scholes model.
	D2        float64 // The d2 term in the Black Scholes model.
	CDF1      float64 // The CDF of d1.
	CDF2      float64 // The CDF of d2.
	DeltaCall float64
	DeltaPut  float64
	Gamma     float64
	Call      float64
	Put       float64
}

// Compute the model.
//
func (b BlackScholes) Compute() (result BlackScholesResult) {
	//
	// d1
	//
	result.D1 = b.RiskFreeRate - b.Yield
	result.D1 += 0.5 * b.Volatility * b.Volatility
	result.D1 *= b.Term
	result.D1 += math.Log(b.Price / b.Strike)
	vt := b.Volatility * math.Sqrt(b.Term)
	result.D1 /= vt
	//
	// d2
	//
	result.D2 = result.D1 - vt
	//
	// CDF's
	//
	result.CDF1 = CDF(result.D1)
	result.CDF2 = CDF(result.D2)
	//
	// Delta's
	//
	result.DeltaCall = result.CDF1
	result.DeltaPut = result.DeltaCall - 1
	//
	// Gamma
	//
	yt := math.Exp(-b.Yield * b.Term)
	result.Gamma = yt * PDF(result.D1)
	result.Gamma /= b.Price * vt
	//
	// Call
	//
	rt := math.Exp(-b.RiskFreeRate * b.Term)
	result.Call = b.Price * yt * result.DeltaCall
	result.Call -= rt * b.Strike * result.CDF2
	//
	// Put
	//
	result.Put = rt * b.Strike * (1 - result.CDF2) // N(-d2) = 1 - N(d2)
	result.Put += b.Price * yt * result.DeltaPut
	return
}
