package multithreaded

import (
	"runtime"
	"sync"
)

const DefaultMemCopySize = 1_000_000

// MemoryCopySetBenchmark performs multithreaded memory copy and set operations
// on a buffer of the provided size. When size is non-positive,
// DefaultMemCopySize is used.
func MemoryCopySetBenchmark(size int) {
	if size <= 0 {
		size = DefaultMemCopySize
	}

	src := make([]byte, size)
	dst := make([]byte, size)

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := size / numWorkers
	remainder := size % numWorkers

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > size {
			endIndex = size
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			// Memory Set
			for i := start; i < end; i++ {
				src[i] = byte(i)
			}
			// Memory Copy
			copy(dst[start:end], src[start:end])
		}(startIndex, endIndex)

		startIndex = endIndex
		if startIndex >= size {
			break
		}
	}
	wg.Wait()
}
