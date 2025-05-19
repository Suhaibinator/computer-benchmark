package singlethreaded

import "math/rand/v2"

// Matrix Multiplication Benchmark
// LargeMatrixMultiplicationSize controls the default workload for the
// MatrixMultiplicationBenchmark when used from the command line.
// Tests should supply a smaller value.
var LargeMatrixMultiplicationSize = 512 * 4

// MatrixMultiplicationBenchmark performs a simple matrix multiplication using
// the provided matrix dimension. If size is <= 0 a small default is used.
func MatrixMultiplicationBenchmark(size int) {
	if size <= 0 {
		size = 256
	}
	A := make([][]float64, size)
	B := make([][]float64, size)
	C := make([][]float64, size)

	// Initialize matrices with random values
	for i := 0; i < size; i++ {
		A[i] = make([]float64, size)
		B[i] = make([]float64, size)
		C[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			A[i][j] = rand.Float64()
			B[i][j] = rand.Float64()
		}
	}

	// Perform matrix multiplication
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum := 0.0
			for k := 0; k < size; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
}
