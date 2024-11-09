package singlethreaded

import "crypto/rand"

// Memory Copy and Set Operations Benchmark
func MemoryCopySetBenchmark() {
	size := 25 * 1000_000_000 // Size in bytes
	src := make([]byte, size)
	dst := make([]byte, size)

	// Initialize source with random data
	_, err := rand.Read(src)
	if err != nil {
		panic(err)
	}

	// Memory Copy
	copy(dst, src)

	// Memory Set
	for i := 0; i < size; i++ {
		dst[i] = 0
	}
}
