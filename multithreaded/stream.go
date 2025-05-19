package multithreaded

import (
	"runtime"
	"sync"
)

// LargeStreamArraySize defines the default workload for the multithreaded
// StreamBenchmark when executed from the command line.
var LargeStreamArraySize = 6 * 1_000_000_000

// StreamBenchmark performs a multithreaded memory bandwidth test using the
// provided size. If size <= 0 a default of one million elements is used.
func StreamBenchmark(size int) {
	if size <= 0 {
		size = 1_000_000
	}
	a := make([]float64, size)
	b := make([]float64, size)
	c := make([]float64, size)

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
		if startIndex >= size {
			break
		}
	}
	wg.Wait()
}
