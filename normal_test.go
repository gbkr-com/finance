package finance

import (
	"testing"
)

func TestPDF(t *testing.T) {
	cases := []struct {
		p        float64
		expected float64
	}{
		{p: 0.1, expected: 0.396953},
		{p: 0.5, expected: 0.352065},
		{p: 0.9, expected: 0.266085},
	}
	for _, c := range cases {
		if x := Round(PDF(c.p), 6); x != c.expected {
			t.Error(x, c.expected)
		}
	}
}

func TestCDF(t *testing.T) {
	cases := []struct {
		p        float64
		expected float64
	}{
		{p: 0.1, expected: 0.539828},
		{p: 0.5, expected: 0.691462},
		{p: 0.9, expected: 0.815940},
	}
	for _, c := range cases {
		if x := Round(CDF(c.p), 6); x != c.expected {
			t.Error(x, c.expected)
		}
	}
}
