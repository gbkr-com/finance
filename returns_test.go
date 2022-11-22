package finance

import (
	"testing"
)

func TestReturns(t *testing.T) {
	hsbc := make([]float64, len(HSBC)-1)
	if LinearReturns(HSBC, hsbc) != 19 {
		t.Error()
	}
	if Round(hsbc[0], 4) != 0.0191 {
		t.Error()
	}
	if CompoundReturns(HSBC, hsbc) != 19 {
		t.Error()
	}
	if Round(hsbc[0], 4) != 0.0189 {
		t.Error()
	}
}

func BenchmarkCompoundReturns(b *testing.B) {
	r := make([]float64, len(HSBC)-1)
	for n := 0; n < b.N; n++ {
		CompoundReturns(HSBC, r)
	}
	dummySeries = r
}
