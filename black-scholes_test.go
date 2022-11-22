package finance

import (
	"testing"
)

func TestCompute(t *testing.T) {
	model := BlackScholes{
		Price:        60,
		Strike:       58,
		Yield:        0.0125,
		Volatility:   0.2,
		Term:         0.5,
		RiskFreeRate: 0.035,
	}
	r := model.Compute()
	if Round(r.D1, 6) != 0.389980 {
		t.Error()
	}
	if Round(r.D2, 6) != 0.248559 {
		t.Error()
	}
	if Round(r.DeltaCall, 6) != 0.651724 {
		t.Error()
	}
	if Round(r.DeltaPut, 6) != -0.348276 {
		t.Error()
	}
	if Round(r.Call, 6) != 4.769029 {
		t.Error()
	}
	if Round(r.Put, 6) != 2.136689 {
		t.Error()
	}
}

func BenchmarkCompute(b *testing.B) {
	model := BlackScholes{
		Price:        60,
		Strike:       58,
		Yield:        0.0125,
		Volatility:   0.2,
		Term:         0.5,
		RiskFreeRate: 0.035,
	}
	var r BlackScholesResult
	for n := 0; n < b.N; n++ {
		r = model.Compute()
	}
	dummy = r.Call
}
