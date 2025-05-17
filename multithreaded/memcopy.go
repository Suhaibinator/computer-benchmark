package multithreaded

import (
	"runtime"
	"sync"
)

var memcopyArraySize int64 = 1_000_000_000 // Fixed array size

// MemoryCopySetBenchmark performs multithreaded memory copy and set operations
func MemoryCopySetBenchmark() {
	src := make([]byte, int(memcopyArraySize))
	dst := make([]byte, int(memcopyArraySize))

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := int(memcopyArraySize) / numWorkers
	remainder := int(memcopyArraySize) % numWorkers

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > int(memcopyArraySize) {
			endIndex = int(memcopyArraySize)
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
		if startIndex >= int(memcopyArraySize) {
			break
		}
	}
	wg.Wait()
}
