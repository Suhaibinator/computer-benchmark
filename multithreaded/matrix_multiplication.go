package multithreaded

import (
	"math/rand"
	"runtime"
	"sync"
)

// LargeMatrixMultiplicationSize holds the matrix dimension used by the
// multithreaded matrix multiplication benchmark when run from the command line.
const LargeMatrixMultiplicationSize = 4096 // Fixed matrix size for consistent workload

// MatrixMultiplicationBenchmark performs multithreaded matrix multiplication
// with the provided matrix dimension. If size <= 0 a smaller default is used.
func MatrixMultiplicationBenchmark(size int) {
	if size <= 0 {
		size = 256
	}
	A := generateMatrix(size)
	B := generateMatrix(size)
	C := make([][]float64, size)
	for i := range C {
		C[i] = make([]float64, size)
	}

	numWorkers := runtime.NumCPU() * 3 // Use thrice the number of CPUs
	var wg sync.WaitGroup

	// Calculate the number of rows per worker and handle any remainder
	rowsPerWorker := size / numWorkers
	remainder := size % numWorkers

	startRow := 0
	for w := 0; w < numWorkers; w++ {
		// Distribute extra rows to the first 'remainder' workers
		endRow := startRow + rowsPerWorker
		if w < remainder {
			endRow++
		}
		if endRow > size {
			endRow = size
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				for j := 0; j < size; j++ {
					sum := 0.0
					for k := 0; k < size; k++ {
						sum += A[i][k] * B[k][j]
					}
					C[i][j] = sum
				}
			}
		}(startRow, endRow)

		startRow = endRow
		if startRow >= size {
			break
		}
	}
	wg.Wait()
}

func generateMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		for j := range matrix[i] {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}
