package singlethreaded

// LargeStreamSize defines the default array length used by the
// StreamBenchmark when executed from the command line.
var LargeStreamSize = 1_000_000_000

// StreamBenchmark performs a memory bandwidth test using the provided size.
// If size <= 0 a modest default of one million elements is used.
func StreamBenchmark(size int) {
	if size <= 0 {
		size = 1_000_000
	}
	a := make([]float64, size)
	b := make([]float64, size)
	c := make([]float64, size)
	scalar := 3.0

	// Initialize arrays
	for i := 0; i < size; i++ {
		a[i] = 1.0
		b[i] = 2.0
	}

	// Perform STREAM Copy
	for i := 0; i < size; i++ {
		c[i] = a[i]
	}

	// Perform STREAM Scale
	for i := 0; i < size; i++ {
		b[i] = scalar * c[i]
	}

	// Perform STREAM Add
	for i := 0; i < size; i++ {
		c[i] = a[i] + b[i]
	}

	// Perform STREAM Triad
	for i := 0; i < size; i++ {
		a[i] = b[i] + scalar*c[i]
	}
}
