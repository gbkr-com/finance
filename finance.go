package finance

// Sharpe returns the Sharpe ratio for an investment using the investment rate
// of return, the risk free rate and the standard deviation of the investment
// return.
//
func Sharpe(rate, riskFreeRate, stdDev float64) float64 {
	return (rate - riskFreeRate) / stdDev
}

// Kelly returns the Kelly fraction for a prospective investment rate of return.
//
func Kelly(rate, riskFreeRate, variance float64) float64 {
	kelly := rate - riskFreeRate
	if kelly <= 0 {
		return 0
	}
	return kelly / variance
}
