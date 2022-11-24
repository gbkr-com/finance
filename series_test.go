package finance

import (
	"testing"
)

func TestSeries(t *testing.T) {
	x := NewSeries(HSBC, false)
	if Round(x.Mean, 2) != 49.43 {
		t.Error()
	}
	if Round(x.Variance, 2) != 2.63 {
		t.Error()
	}
	if Round(x.MAD, 2) != 1.41 {
		t.Error()
	}
	y := NewSeries(HSI, false)
	r := x.Compare(y)
	if Round(r.Covariance, 2) != 577.08 {
		t.Error()
	}
	if Round(r.R, 2) != 0.67 {
		t.Error()
	}
	if Round(r.R2, 2) != 0.45 {
		t.Error()
	}
	if Round(r.Gradient, 2) != 219.73 {
		t.Error()
	}
	if Round(r.Intercept, 2) != 9520.39 {
		t.Error()
	}
}

func BenchmarkBeta(b *testing.B) {
	hsi := make([]float64, len(HSI)-1)
	sands := make([]float64, len(SANDS)-1)
	LinearReturns(HSI, hsi)
	LinearReturns(SANDS, sands)
	index := NewSeries(hsi, false)
	var r float64
	for n := 0; n < b.N; n++ {
		x := index.Compare(NewSeries(sands, false))
		r = x.Gradient
	}
	dummy = r
}
