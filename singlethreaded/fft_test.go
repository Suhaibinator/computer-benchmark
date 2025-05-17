package singlethreaded

import (
	"math"
	"math/cmplx"
	"testing"
)

// naiveDFT computes the discrete Fourier transform using the direct formula.
func naiveDFT(a []complex128) []complex128 {
	n := len(a)
	result := make([]complex128, n)
	for k := 0; k < n; k++ {
		var sum complex128
		for t := 0; t < n; t++ {
			angle := -2 * math.Pi * float64(t*k) / float64(n)
			sum += a[t] * cmplx.Exp(complex(0, angle))
		}
		result[k] = sum
	}
	return result
}

func TestFFTSmallArray(t *testing.T) {
	data := []complex128{1, 2, 3, 4}
	expected := naiveDFT(data)
	// Copy input since fft works in-place
	got := append([]complex128(nil), data...)
	fft(got)
	for i := range expected {
		if cmplx.Abs(got[i]-expected[i]) > 1e-9 {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], got[i])
		}
	}
}

func TestFFTRandomSmallArray(t *testing.T) {
	input := []complex128{0.5, -1.2, 3.3, 4.4, 5.5, -6.6, 7.7, 8.8}
	expected := naiveDFT(input)
	got := append([]complex128(nil), input...)
	fft(got)
	for i := range expected {
		if cmplx.Abs(got[i]-expected[i]) > 1e-9 {
			t.Errorf("index %d: expected %v, got %v", i, expected[i], got[i])
		}
	}
}
