package singlethreaded

import "math/rand/v2"

// Matrix Multiplication Benchmark
const DefaultMatrixSize = 256

// MatrixMultiplicationBenchmark performs a naive matrix multiplication of an
// NxN matrix. When size is non-positive, DefaultMatrixSize is used instead.
func MatrixMultiplicationBenchmark(size int) {
	if size <= 0 {
		size = DefaultMatrixSize
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
