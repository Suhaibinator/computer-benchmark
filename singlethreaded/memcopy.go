package singlethreaded

import "crypto/rand"

// Memory Copy and Set Operations Benchmark
func MemoryCopySetBenchmark() {
	var size int64 = 25 * 1_000_000_000 // Size in bytes
	src := make([]byte, int(size))
	dst := make([]byte, int(size))

	// Initialize source with random data
	_, err := rand.Read(src)
	if err != nil {
		panic(err)
	}

	// Memory Copy
	copy(dst, src)

	// Memory Set
	for i := 0; i < int(size); i++ {
		dst[i] = 0
	}
}
