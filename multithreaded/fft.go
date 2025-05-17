package multithreaded

import (
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"sync"
)

const DefaultFFTSize = 1 << 16

// FftBenchmark performs a multithreaded Fast Fourier Transform on a slice of
// the given size. The size must be a power of two. When size is non-positive,
// DefaultFFTSize is used.
func FftBenchmark(size int) {
	if size <= 0 {
		size = DefaultFFTSize
	}
	data := make([]complex128, size)
	for i := range data {
		data[i] = complex(rand.Float64(), 0)
	}
	parallelFFT(data, runtime.NumCPU()*3)
}

// parallelFFT performs a parallel FFT using a worker limit
func parallelFFT(a []complex128, maxWorkers int) {
	n := len(a)
	if n <= 1 || maxWorkers <= 1 {
		fft(a)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Split even and odd indices
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[i*2]
		odd[i] = a[i*2+1]
	}

	go func() {
		defer wg.Done()
		parallelFFT(even, maxWorkers/2)
	}()
	go func() {
		defer wg.Done()
		parallelFFT(odd, maxWorkers-maxWorkers/2)
	}()
	wg.Wait()

	// Combine
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * odd[k]
		a[k] = even[k] + t
		a[k+n/2] = even[k] - t
	}
}

// fft performs a sequential FFT
func fft(a []complex128) {
	n := len(a)
	if n <= 1 {
		return
	}

	// Split even and odd indices
	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[i*2]
		odd[i] = a[i*2+1]
	}

	fft(even)
	fft(odd)

	// Combine
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * odd[k]
		a[k] = even[k] + t
		a[k+n/2] = even[k] - t
	}
}
