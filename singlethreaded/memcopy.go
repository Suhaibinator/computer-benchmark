package singlethreaded

import "crypto/rand"

// Memory Copy and Set Operations Benchmark
// LargeMemoryCopySize specifies the byte count used by MemoryCopySetBenchmark
// when run from the command line.
var LargeMemoryCopySize = 25 * 1_000_000_000

// MemoryCopySetBenchmark exercises memory copy and set operations on a slice of
// the given size. If size <= 0 a default of one million bytes is used.
func MemoryCopySetBenchmark(size int) {
	if size <= 0 {
		size = 1_000_000
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
	for i := 0; i < size; i++ {
		dst[i] = 0
	}
}
