package singlethreaded

const DefaultStreamSize = 1_000_000

// StreamBenchmark performs the classic STREAM memory bandwidth test on an
// array of the provided size. If size is non-positive, DefaultStreamSize is
// used.
func StreamBenchmark(size int) {
	if size <= 0 {
		size = DefaultStreamSize
	}

	a := make([]float64, size)
	b := make([]float64, size)
	c := make([]float64, size)
	scalar := 3.0

	// Initialize arrays
	for i := 0; i < int(size); i++ {
		a[i] = 1.0
		b[i] = 2.0
	}

	// Perform STREAM Copy
	for i := 0; i < int(size); i++ {
		c[i] = a[i]
	}

	// Perform STREAM Scale
	for i := 0; i < int(size); i++ {
		b[i] = scalar * c[i]
	}

	// Perform STREAM Add
	for i := 0; i < int(size); i++ {
		c[i] = a[i] + b[i]
	}

	// Perform STREAM Triad
	for i := 0; i < int(size); i++ {
		a[i] = b[i] + scalar*c[i]
	}
}
