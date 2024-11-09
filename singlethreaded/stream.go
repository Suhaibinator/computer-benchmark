package singlethreaded

// STREAM Benchmark
func StreamBenchmark() {
	size := 1000000000 // Reduced size to prevent excessive memory usage
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
