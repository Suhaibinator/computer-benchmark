package multithreaded

import (
	"math/rand"
	"runtime"
	"sync"
)

const DefaultMatrixSize = 256

// MatrixMultiplicationBenchmark performs multithreaded matrix multiplication on
// an NxN matrix. When size is non-positive, DefaultMatrixSize is used.
func MatrixMultiplicationBenchmark(size int) {
	if size <= 0 {
		size = DefaultMatrixSize
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
