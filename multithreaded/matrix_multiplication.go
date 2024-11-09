package multithreaded

import (
	"math/rand"
	"runtime"
	"sync"
)

const matrixMultiplicationSize = 4096 // Fixed matrix size for consistent workload

// MatrixMultiplicationBenchmark performs multithreaded matrix multiplication
func MatrixMultiplicationBenchmark() {
	A := generateMatrix(matrixMultiplicationSize)
	B := generateMatrix(matrixMultiplicationSize)
	C := make([][]float64, matrixMultiplicationSize)
	for i := range C {
		C[i] = make([]float64, matrixMultiplicationSize)
	}

	numWorkers := runtime.NumCPU() * 3 // Use thrice the number of CPUs
	var wg sync.WaitGroup

	// Calculate the number of rows per worker and handle any remainder
	rowsPerWorker := matrixMultiplicationSize / numWorkers
	remainder := matrixMultiplicationSize % numWorkers

	startRow := 0
	for w := 0; w < numWorkers; w++ {
		// Distribute extra rows to the first 'remainder' workers
		endRow := startRow + rowsPerWorker
		if w < remainder {
			endRow++
		}
		if endRow > matrixMultiplicationSize {
			endRow = matrixMultiplicationSize
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				for j := 0; j < matrixMultiplicationSize; j++ {
					sum := 0.0
					for k := 0; k < matrixMultiplicationSize; k++ {
						sum += A[i][k] * B[k][j]
					}
					C[i][j] = sum
				}
			}
		}(startRow, endRow)

		startRow = endRow
		if startRow >= matrixMultiplicationSize {
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
