package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Suhaibinator/computer-benchmark/multithreaded"
	"github.com/Suhaibinator/computer-benchmark/singlethreaded"
)

// BenchmarkResult holds the execution time of a benchmark
type BenchmarkResult struct {
	Name     string
	Duration time.Duration
}

// Benchmark represents a named benchmark function
type Benchmark struct {
	Name string
	Run  func()
}

var benchmarkList = []Benchmark{
	{"Matrix Multiplication", func() { singlethreaded.MatrixMultiplicationBenchmark(singlethreaded.LargeMatrixMultiplicationSize) }},
	{"STREAM Memory Bandwidth", func() { singlethreaded.StreamBenchmark(singlethreaded.LargeStreamSize) }},
	{"Fast Fourier Transform", func() { singlethreaded.FftBenchmark(singlethreaded.LargeFFTSize) }},
	{"Memory Copy and Set", func() { singlethreaded.MemoryCopySetBenchmark(singlethreaded.LargeMemoryCopySize) }},
	{"Sorting Algorithms", func() { singlethreaded.SortingBenchmark(singlethreaded.LargeSortSize) }},
	{"AES Encryption/Decryption", func() { singlethreaded.CryptoBenchmark(singlethreaded.LargePlaintextSize) }},
	{"Multithreaded Matrix Multiplication", func() { multithreaded.MatrixMultiplicationBenchmark(multithreaded.LargeMatrixMultiplicationSize) }},
	{"Multithreaded STREAM Memory Bandwidth", func() { multithreaded.StreamBenchmark(multithreaded.LargeStreamArraySize) }},
	{"Multithreaded Fast Fourier Transform", func() { multithreaded.FftBenchmark(multithreaded.LargeFFTSize) }},
	{"Multithreaded Memory Copy and Set", func() { multithreaded.MemoryCopySetBenchmark(multithreaded.LargeMemcopyArraySize) }},
	{"Multithreaded Sorting Algorithms", func() { multithreaded.SortingBenchmark(multithreaded.LargeSortArraySize) }},
	{"Multithreaded AES Encryption", func() { multithreaded.CryptoBenchmark(multithreaded.LargeDataSize) }},
}

func runBenchmark(name string, fn func()) BenchmarkResult {
	start := time.Now()
	fn()
	duration := time.Since(start)
	return BenchmarkResult{Name: name, Duration: duration}
}

func runAllBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0, len(benchmarkList))
	for _, b := range benchmarkList {
		fmt.Println("Starting " + b.Name)
		results = append(results, runBenchmark(b.Name, b.Run))
	}
	return results
}

func printResults(results []BenchmarkResult) {
	fmt.Println("\nBenchmark Results:")
	for _, result := range results {
		fmt.Printf("%-40s: %v\n", result.Name, result.Duration)
	}
	fmt.Println()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select a benchmark to run:")
		fmt.Println("0) Run all benchmarks")
		for i, b := range benchmarkList {
			fmt.Printf("%d) %s\n", i+1, b.Name)
		}
		fmt.Print("q) Quit\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "q") {
			return
		}

		if input == "0" {
			results := runAllBenchmarks()
			printResults(results)
			continue
		}

		idx, err := strconv.Atoi(input)
		if err != nil || idx < 1 || idx > len(benchmarkList) {
			fmt.Println("Invalid selection")
			continue
		}

		b := benchmarkList[idx-1]
		fmt.Println("Starting " + b.Name)
		result := runBenchmark(b.Name, b.Run)
		printResults([]BenchmarkResult{result})
	}
}
