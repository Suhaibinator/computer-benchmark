package singlethreaded

import "crypto/rand"

// Memory Copy and Set Operations Benchmark
const DefaultMemCopySize = 1_000_000

// MemoryCopySetBenchmark performs copy and set operations on a buffer of the
// provided size. A non-positive size falls back to DefaultMemCopySize.
func MemoryCopySetBenchmark(size int) {
	if size <= 0 {
		size = DefaultMemCopySize
	}

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
	for i := 0; i < int(size); i++ {
		dst[i] = 0
	}
}
