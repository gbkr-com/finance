package finance

// EMA is the exponential moving average.
//
type EMA struct {
	previous float64
	k        float64
}

// Init the moving average with an average over the given number of periods.
//
func (e *EMA) Init(avg float64, periods int) {
	e.previous = avg
	e.k = 2.0 / float64(periods+1)
}

// Next returns the next iteration of the EMA given the price.
//
func (e *EMA) Next(price float64) float64 {
	e.previous = e.k*(price-e.previous) + e.previous
	return e.previous
}
