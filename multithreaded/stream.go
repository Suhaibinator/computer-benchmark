package multithreaded

import (
	"runtime"
	"sync"
)

var streamArraySize int64 = 6 * 1_000_000_000 // Fixed array size

// StreamBenchmark performs multithreaded memory bandwidth test
func StreamBenchmark() {
	a := make([]float64, int(streamArraySize))
	b := make([]float64, int(streamArraySize))
	c := make([]float64, int(streamArraySize))

	numWorkers := runtime.NumCPU() * 3
	var wg sync.WaitGroup

	chunkSize := int(streamArraySize) / numWorkers
	remainder := int(streamArraySize) % numWorkers

	startIndex := 0
	for w := 0; w < numWorkers; w++ {
		endIndex := startIndex + chunkSize
		if w < remainder {
			endIndex++
		}
		if endIndex > int(streamArraySize) {
			endIndex = int(streamArraySize)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			// Initialize arrays
			for i := start; i < end; i++ {
				a[i] = 1.0
				b[i] = 2.0
			}
			// Perform STREAM operation
			for i := start; i < end; i++ {
				c[i] = a[i] + b[i]
			}
		}(startIndex, endIndex)

		startIndex = endIndex
		if startIndex >= int(streamArraySize) {
			break
		}
	}
	wg.Wait()
}
