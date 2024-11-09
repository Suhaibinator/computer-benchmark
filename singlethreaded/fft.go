package singlethreaded

import (
	"math"
	"math/cmplx"
	"math/rand/v2"
)

// Fast Fourier Transform Benchmark
func FftBenchmark() {
	size := 1 << 20 // FFT size (must be power of 2)
	data := make([]complex128, size)

	// Initialize data with random values
	for i := range data {
		data[i] = complex(rand.Float64(), 0)
	}

	// Perform FFT
	fft(data)
}

// Simple Cooley-Tukey FFT algorithm
func fft(a []complex128) {
	n := len(a)
	if n <= 1 {
		return
	}

	// Divide
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[2*i]
		odd[i] = a[2*i+1]
	}

	// Conquer
	fft(even)
	fft(odd)

	// Combine
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * odd[k]
		a[k] = even[k] + t
		a[k+n/2] = even[k] - t
	}
}
