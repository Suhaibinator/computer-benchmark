package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Suhaibinator/computer-benchmark/multithreaded"
	"github.com/Suhaibinator/computer-benchmark/singlethreaded"
)

// BenchmarkResult holds the execution time of a benchmark
type BenchmarkResult struct {
	Name     string
	Duration time.Duration
}

func main() {
	// Ensure maximum CPU utilization
	runtime.GOMAXPROCS(runtime.NumCPU())

	var results []BenchmarkResult

	fmt.Println("Starting Matrix Multiplication")
	// Matrix Multiplication Benchmark
	start := time.Now()
	singlethreaded.MatrixMultiplicationBenchmark()
	duration := time.Since(start)
	results = append(results, BenchmarkResult{"Matrix Multiplication", duration})

	fmt.Println("Starting STREAM Benchmark")
	// STREAM Benchmark
	start = time.Now()
	singlethreaded.StreamBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"STREAM Memory Bandwidth", duration})

	fmt.Println("Starting FFT Benchmark")
	// FFT Benchmark
	start = time.Now()
	singlethreaded.FftBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Fast Fourier Transform", duration})

	fmt.Println("Starting Memcopy Benchmark")
	// Memory Copy and Set Operations Benchmark
	start = time.Now()
	singlethreaded.MemoryCopySetBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Memory Copy and Set", duration})

	fmt.Println("Start Sorting Benchmark")
	// Sorting Benchmark
	start = time.Now()
	singlethreaded.SortingBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Sorting Algorithms", duration})

	fmt.Println("Start Cryptographic Benchmark")
	// Cryptographic Algorithm Benchmark
	start = time.Now()
	singlethreaded.CryptoBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"AES Encryption/Decryption", duration})

	// Print the benchmark results
	fmt.Println("Benchmark Results:")
	for _, result := range results {
		fmt.Printf("%-30s: %v\n", result.Name, result.Duration)
	}

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Using %d CPU cores and %d goroutines per benchmark\n", numCPU, numCPU*3)

	fmt.Println("Starting Multithreaded Matrix Multiplication")
	start = time.Now()
	multithreaded.MatrixMultiplicationBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded Matrix Multiplication", duration})

	fmt.Println("Starting Multithreaded STREAM Benchmark")
	start = time.Now()
	multithreaded.StreamBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded STREAM Memory Bandwidth", duration})

	fmt.Println("Starting Multithreaded FFT Benchmark")
	start = time.Now()
	multithreaded.FftBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded Fast Fourier Transform", duration})

	fmt.Println("Starting Multithreaded Memory Copy and Set Benchmark")
	start = time.Now()
	multithreaded.MemoryCopySetBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded Memory Copy and Set", duration})

	fmt.Println("Starting Multithreaded Sorting Benchmark")
	start = time.Now()
	multithreaded.SortingBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded Sorting Algorithms", duration})

	fmt.Println("Starting Multithreaded Cryptographic Benchmark")
	start = time.Now()
	multithreaded.CryptoBenchmark()
	duration = time.Since(start)
	results = append(results, BenchmarkResult{"Multithreaded AES Encryption", duration})

	// Print the benchmark results
	fmt.Println("\nMultithreaded Benchmark Results:")
	for _, result := range results {
		fmt.Printf("%-40s: %v\n", result.Name, result.Duration)
	}
}
