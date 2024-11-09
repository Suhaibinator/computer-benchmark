package multithreaded

import (
	"runtime"
	"sync"
)

const memcopyArraySize = 1_000_000_000 // Fixed array size

// MemoryCopySetBenchmark performs multithreaded memory copy and set operations
func MemoryCopySetBenchmark() {
	src := make([]byte, memcopyArraySize)
	dst := make([]byte, memcopyArraySize)

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := memcopyArraySize / numWorkers
	remainder := memcopyArraySize % numWorkers

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > memcopyArraySize {
			endIndex = memcopyArraySize
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
		if startIndex >= memcopyArraySize {
			break
		}
	}
	wg.Wait()
}
